# 📬 Mailer Client for Go

**Mailer Client** is a lightweight Go package that allows your applications to send emails via the internal [Bakiverse](https://bakiverse.com) Mailer API.
It provides a simple and extensible abstraction over HTTP, so you can integrate email sending logic in any backend or CLI written in Go.

---

## ✨ Features

- ✅ Simple and developer-friendly API
- ✅ HTML or plain text support
- ✅ Multiple recipients
- ✅ Lightweight, no external dependencies beyond `net/http`
- ✅ Easily pluggable into any Go backend

---

## 🚀 Installation

```bash
go get github.com/bakiversehq/mailer-client/mailer
```

---

## 🛠 Usage

```go
package main

import (
    "log"
    "github.com/bakiversehq/mailer-client/mailer"
)

func main() {
    client := mailer.NewClient("https://mailer.bakiverse.com")

    err := client.Send(mailer.EmailReq{
        Creds: mailer.Creds{
            Email: "noreply@bakiverse.com",
            Pwd:   "your-password",
        },
        FromName: "Bakiverse Bot",
        ToList:   []string{"user@example.com"},
        Subject:  "Welcome to Bakiverse!",
        Body:     "<h1>Hello world</h1><p>This is a test.</p>",
        Html:     true,
    })

    if err != nil {
        log.Fatal("Email failed:", err)
    }
}
```

---

## 💡 API Reference

### `func NewClient(baseURL string) *Client`
Initializes a new mailer client.
- `baseURL`: The full base URL to the mailer backend (e.g. `https://mailer.bakiverse.com`)

### `func (c *Client) Send(req EmailReq) error`
Sends an email request.
- Returns an error if sending fails or the mailer backend returns an error.

---

## 📄 Structures

### `EmailReq`
| Field     | Type     | Required | Description                            |
|-----------|----------|----------|----------------------------------------|
| `Creds`   | `Creds`  | Yes      | Email + password credentials           |
| `ToList`  | `[]string` | Yes    | List of recipients                     |
| `Subject` | `string` | Yes      | Subject of the email                   |
| `Body`    | `string` | Yes      | Email content (HTML or plain text)     |
| `Html`    | `bool`   | Yes      | Set to true for HTML emails            |
| `FromName`| `string` | No       | Optional display name for sender       |

### `Creds`
| Field   | Type   | Description            |
|---------|--------|------------------------|
| `Email` | string | Email used to login    |
| `Pwd`   | string | Corresponding password |

---

## 🌐 License
MIT — feel free to use in personal and commercial projects.

---

## 🚀 Roadmap Ideas
- Optional CC and BCC fields
- Attachments support
- Retry on failure

---

Made with ❤️ by the [Bakiverse](https://bakiverse.com) team.

