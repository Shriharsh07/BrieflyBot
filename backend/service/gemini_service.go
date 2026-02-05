package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
	"google.golang.org/api/gmail/v1"
)

type EmailData struct {
	Subject string
	Body    string
}

func SummarizeEmailsBatch(emails []EmailData) ([]string, error) {

	_ = godotenv.Load()

	apiKey := os.Getenv("GEMINI_API_KEY")

	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1/models/gemini-2.5-flash:generateContent?key=%s",
		apiKey,
	)

	var input strings.Builder
	for i, e := range emails {
		input.WriteString(fmt.Sprintf(
			"\nEMAIL %d\nSubject: %s\nBody: %s\n",
			i+1, e.Subject, e.Body,
		))
	}

	prompt := fmt.Sprintf(
		`You are my personal assistant.

		Explain each email below in SIMPLE language.

		Rules:
		- Use short bullet points
		- Max 4 bullets per email
		- Explain what the email is about
		- Explain if any action is needed
		- Do NOT repeat the subject
		- Do NOT rewrite the email

		Return summaries in the SAME ORDER, separated by "---".
		%s`,
		input.String(),
	)

	payload := map[string]interface{}{
		"contents": []map[string]any{
			{
				"parts": []map[string]string{
					{"text": prompt},
				},
			},
		},
	}

	data, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	var result struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}

	json.Unmarshal(bodyBytes, &result)

	if len(result.Candidates) == 0 {
		log.Println("Gemini returned zero candidates:", string(bodyBytes))
		return fallbackSummaries(emails), nil
	}

	if len(result.Candidates[0].Content.Parts) == 0 {
		log.Println("Gemini returned empty content parts:", string(bodyBytes))
		return fallbackSummaries(emails), nil
	}

	raw := result.Candidates[0].Content.Parts[0].Text
	parts := strings.Split(raw, "---")

	out := []string{}
	for _, p := range parts {
		out = append(out, CleanSummary(p))
	}

	return out, nil
}

func CleanSummary(text string) string {
	lines := strings.Split(text, "\n")
	out := []string{}

	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		if strings.HasPrefix(l, "*") || strings.HasPrefix(l, "-") {
			l = "• " + strings.TrimSpace(l[1:])
		} else if !strings.HasPrefix(l, "•") {
			l = "• " + l
		}
		out = append(out, l)
	}
	return strings.Join(out, "\n")
}

func GetEmailBody(part *gmail.MessagePart) string {
	if part.Body != nil && part.Body.Data != "" {
		return decodeBody(part.Body.Data)
	}

	for _, p := range part.Parts {
		if p.MimeType == "text/plain" && p.Body != nil && p.Body.Data != "" {
			return decodeBody(p.Body.Data)
		}
	}

	for _, p := range part.Parts {
		if p.MimeType == "text/html" && p.Body != nil && p.Body.Data != "" {
			return stripHTML(decodeBody(p.Body.Data))
		}
	}

	for _, p := range part.Parts {
		if body := GetEmailBody(p); body != "" {
			return body
		}
	}

	return ""
}

func decodeBody(data string) string {
	decoded, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return ""
	}
	return string(decoded)
}

func stripHTML(html string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	return strings.TrimSpace(re.ReplaceAllString(html, ""))
}

func fallbackSummaries(emails []EmailData) []string {
	out := []string{}
	for _, e := range emails {
		out = append(out,
			"• New email received\n"+
				"• Subject: "+e.Subject+"\n"+
				"• Please review the email",
		)
	}
	return out
}

func StripBulletsForUI(s string) string {
	lines := strings.Split(s, "\n")
	out := []string{}

	for _, line := range lines {
		l := strings.TrimSpace(line)

		// Remove common bullet styles
		l = strings.TrimPrefix(l, "•")
		l = strings.TrimPrefix(l, "-")
		l = strings.TrimPrefix(l, "*")

		// Remove numbered bullets: 1. 2)
		l = regexp.MustCompile(`^\d+[\.\)]\s*`).ReplaceAllString(l, "")

		l = strings.TrimSpace(l)

		if l != "" {
			out = append(out, l)
		}
	}

	return strings.Join(out, "\n")
}
