# 📬 BrieflyBot

**BrieflyBot** is a personal Telegram bot that reads emails from your Gmail inbox, simplifies them into clear, easy-to-understand summaries using **free AI (Google Gemini)**, and sends the summaries directly to you on Telegram.

It is designed for **personal productivity**, privacy, and minimal noise — so you understand emails without reading long content.

---

## ✨ Why BrieflyBot?

- Avoid reading long or formal emails
- Get quick explanations instead of raw email content
- Personal & private — runs only for you
- Uses **FREE AI** (no paid OpenAI required)
- Secure Gmail access using OAuth (no passwords stored)

---

## 🧠 Features

- 📧 Reads unread emails from Gmail
- 🧠 Simplifies emails into short bullet-point summaries
- 🤖 Uses **Google Gemini (free tier)**
- 📬 Sends summaries to Telegram
- 🔐 Secure OAuth-based Gmail access
- 🚫 No email passwords stored

---

## 🏗️ Project Structure
```
├── 📁 .github
│   └── 📁 appmod
│       └── 📁 appcat
├── 📁 service
│   ├── 🐹 gemini_service.go
│   └── 🐹 telegram_service.go
├── ⚙️ .gitignore
├── 📝 README.md
├── 📄 go.mod
├── 📄 go.sum
└── 🐹 main.go
```

---

## 🔑 Prerequisites

- Go **1.21+**
- Gmail account
- Telegram account
- Google account (for Gmail API & Gemini)

---

# 🔐 API KEY & SERVICE SETUP (STEP-BY-STEP)

---

## 1️⃣ Gmail API Setup (Read Emails)

### Step 1: Create Google Cloud Project
1. Open: https://console.cloud.google.com
2. Create a **New Project**
3. Name it: `BrieflyBot`

---

### Step 2: Enable Gmail API
1. Go to **APIs & Services → Library**
2. Search for **Gmail API**
3. Click **Enable**

---

### Step 3: Configure OAuth Consent Screen
1. Go to **APIs & Services → OAuth consent screen**
2. Select **External**
3. App name: `BrieflyBot`
4. Add your email as **Test User**
5. Save and continue

---

### Step 4: Create OAuth Credentials
1. Go to **APIs & Services → Credentials**
2. Click **Create Credentials → OAuth Client ID**
3. Application type: **Desktop App**
4. Name: `BrieflyBot Desktop`
5. Download `credentials.json`
6. Place it in the project root

📌 Gmail uses OAuth — **no email password is stored**.

---

## 2️⃣ Gemini API Key (FREE AI)

BrieflyBot uses **Google Gemini** for summarization.

### Step 1: Open Google AI Studio
👉 https://aistudio.google.com

---

### Step 2: Create API Key
1. Click **Get API Key**
2. Create a new key
3. Copy the key

---

### Step 3: (Optional) Verify Available Models
Open in browser:
👉 https://generativelanguage.googleapis.com/v1/models?key=YOUR_GEMINI_API_KEY


BrieflyBot uses:  `models/gemini-2.5-flash`

---

## 3️⃣ Telegram Bot Setup

### Step 1: Create Telegram Bot
1. Open Telegram
2. Search for **@BotFather**
3. Run `/newbot`
4. Bot name: `BrieflyBot`
5. Username: `brieflybot` (or any available name)
6. Copy the **Bot Token**

---

### Step 2: Get Your Chat ID
1. Open your bot and send:  `hi`

2. Open this URL (replace token): https://api.telegram.org/bot<BOT_TOKEN>getUpdates

3. Copy: `message.chat.id`


---

## ⚙️ Environment Configuration

Create a `.env` file in the project root:

```env
TELEGRAM_BOT_TOKEN=123456:ABCDEF...
TELEGRAM_CHAT_ID=123456789
GEMINI_API_KEY=AIzaSy...
