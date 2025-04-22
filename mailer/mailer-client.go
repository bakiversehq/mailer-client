// Package mailer provides a lightweight client to interact with the internal Bakiverse Mailer API.
//
// It enables any Go application to easily send transactional or custom emails by calling a simple HTTP endpoint.
//
// # Usage
//
//	import "github.com/yourusername/mailer-client/mailer"
//
//	client := mailer.NewClient("https://mailer.bakiverse.com")
//	err := client.Send(mailer.EmailReq{
//	    Creds: mailer.Creds{
//	        Email: "noreply@bakiverse.com",
//	        Pwd:   "your-password",
//	    },
//	    FromName: "Bakiverse Bot",
//	    ToList:   []string{"user@example.com"},
//	    Subject:  "Welcome to Bakiverse!",
//	    Body:     "<h1>Hello world</h1><p>This is a test.</p>",
//	    Html:     true,
//	})
//
//	if err != nil {
//	    log.Fatal("Email failed:", err)
//	}
package mailer

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// Client holds the configuration to interact with the remote Mailer service.
// You can optionally provide your own *http.Client, otherwise a default one with a 10s timeout is used.
type Client struct {
	BaseURL string       // BaseURL of the Mailer API (e.g. https://mailer.bakiverse.com)
	Client  *http.Client // Optional custom HTTP client
}

// EmailReq represents the full request body for sending an email.
type EmailReq struct {
	Creds    Creds    `json:"creds"`      // Authentication credentials for the mailer (email + password)
	ToList   []string `json:"to_list"`    // List of recipient email addresses
	Subject  string   `json:"subject"`    // Subject of the email
	Body     string   `json:"body"`       // Content of the email (HTML or plain text)
	Html     bool     `json:"html"`       // Indicates whether Body is HTML (true) or plain text (false)
	FromName string   `json:"from_name"`  // Display name of the sender
}

// EmailRes describes the JSON structure returned by the mailer API after sending an email.
type EmailRes struct {
	Success bool   `json:"success"` // Indicates whether the email was successfully sent
	Message string `json:"message"` // Additional info or error details
}

// Creds contains the authentication data required by the mailer backend.
type Creds struct {
	Email string `json:"email"` // The email address used to authenticate
	Pwd   string `json:"pwd"`   // The corresponding password
}

// NewClient initializes and returns a Mailer client.
// You must pass the base URL where your Mailer backend is hosted.
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		Client:  &http.Client{Timeout: 10 * time.Second},
	}
}

// Send sends an email request to the configured Mailer API.
//
// It returns an error if the HTTP request fails, or if the backend indicates a failed delivery.
func (c *Client) Send(req EmailReq) error {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	res, err := c.Client.Post(c.BaseURL+"/send", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var resp EmailRes
	if err = json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return errors.New("email sent but failed to decode response")
	}
	if !resp.Success {
		return errors.New("failed to send email: " + resp.Message)
	}

	return nil
}
