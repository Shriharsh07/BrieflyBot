package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type EmailData struct {
	Subject string
	Body    string
}

func main() {
	_ = godotenv.Load()

	ctx := context.Background()

	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatal(err)
	}

	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatal(err)
	}

	client := getClient(config)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := srv.Users.Messages.List("me").
		Q("is:unread").
		MaxResults(10).
		Do()

	if err != nil || len(resp.Messages) == 0 {
		fmt.Println("No unread emails")
		return
	}

	emails := []EmailData{}

	for _, msg := range resp.Messages {
		m, err := srv.Users.Messages.Get("me", msg.Id).Do()
		if err != nil {
			continue
		}

		var subject string
		for _, h := range m.Payload.Headers {
			if h.Name == "Subject" {
				subject = h.Value
			}
		}

		body := getEmailBody(m.Payload)
		if strings.TrimSpace(body) == "" {
			body = subject
		}

		emails = append(emails, EmailData{
			Subject: subject,
			Body:    body,
		})
	}

	summaries, err := SummarizeEmailsBatch(emails)
	if err != nil {
		log.Println("Gemini error:", err)
		return
	}

	for i, s := range summaries {
		SendToTelegram(emails[i].Subject, s)
	}
}

func getClient(config *oauth2.Config) *http.Client {
	tok, err := tokenFromFile("token.json")
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken("token.json", tok)
	}
	return config.Client(context.Background(), tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Println("Open this URL:", url)

	var code string
	fmt.Scan(&code)

	tok, err := config.Exchange(context.Background(), code)
	if err != nil {
		log.Fatal(err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var tok oauth2.Token
	err = json.NewDecoder(f).Decode(&tok)
	return &tok, err
}

func saveToken(path string, token *oauth2.Token) {
	f, _ := os.Create(path)
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
