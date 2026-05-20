package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Swing-Technologies/keplars-sdk/go/keplars"
	"github.com/spf13/cobra"
)

var sendFlags struct {
	to      []string
	from    string
	subject string
	html    string
	text    string
	cc      []string
	bcc     []string
	replyTo string
	timeout int
	json    bool
}

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send an email",
	Long: `Send a transactional email using the Keplars Email API.

Examples:
  # Send a simple text email
  keplars send --to user@example.com --from hello@yourdomain.com --subject "Hello" --text "Hello, World!"

  # Send an HTML email
  keplars send --to user@example.com --from hello@yourdomain.com --subject "Welcome" --html "<h1>Welcome!</h1>"

  # Send to multiple recipients with CC
  keplars send --to user1@example.com --to user2@example.com --cc boss@example.com --from hello@yourdomain.com --subject "Team Update" --text "Hello team!"`,
	RunE: sendEmail,
}

func init() {
	sendCmd.Flags().StringSliceVar(&sendFlags.to, "to", nil, "Recipient email address(es) (can specify multiple)")
	sendCmd.Flags().StringVar(&sendFlags.from, "from", "", "Sender email address")
	sendCmd.Flags().StringVar(&sendFlags.subject, "subject", "", "Email subject")
	sendCmd.Flags().StringVar(&sendFlags.html, "html", "", "HTML content")
	sendCmd.Flags().StringVar(&sendFlags.text, "text", "", "Plain text content")
	sendCmd.Flags().StringSliceVar(&sendFlags.cc, "cc", nil, "CC recipient(s)")
	sendCmd.Flags().StringSliceVar(&sendFlags.bcc, "bcc", nil, "BCC recipient(s)")
	sendCmd.Flags().StringVar(&sendFlags.replyTo, "reply-to", "", "Reply-To address")
	sendCmd.Flags().IntVar(&sendFlags.timeout, "timeout", 30, "Request timeout in seconds")
	sendCmd.Flags().BoolVar(&sendFlags.json, "json", false, "Output response as JSON")

	sendCmd.MarkFlagRequired("to")
	sendCmd.MarkFlagRequired("from")
	sendCmd.MarkFlagRequired("subject")
}

func sendEmail(cmd *cobra.Command, args []string) error {
	if err := checkAPIKey(); err != nil {
		return err
	}

	if sendFlags.html == "" && sendFlags.text == "" {
		return fmt.Errorf("either --html or --text must be provided")
	}

	client, err := keplars.NewClient(apiKey, func(c *keplars.Config) {
		if baseURL != "" {
			c.BaseURL = baseURL
		}
		c.Timeout = time.Duration(sendFlags.timeout) * time.Second
	})
	if err != nil {
		return err
	}

	req := &keplars.SendEmailRequest{
		To:      sendFlags.to,
		From:    sendFlags.from,
		Subject: sendFlags.subject,
		HTML:    sendFlags.html,
		Text:    sendFlags.text,
	}

	for _, addr := range sendFlags.cc {
		req.CC = append(req.CC, keplars.EmailRecipient{Email: addr})
	}
	for _, addr := range sendFlags.bcc {
		req.BCC = append(req.BCC, keplars.EmailRecipient{Email: addr})
	}

	if sendFlags.replyTo != "" {
		req.ReplyTo = sendFlags.replyTo
	}

	ctx := context.Background()
	resp, err := client.Emails.Send(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	if sendFlags.json {
		output, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(output))
	} else {
		fmt.Println("Email sent successfully!")
		fmt.Printf("\nJob ID:   %s\n", resp.Data.JobID)
		fmt.Printf("Priority: %s\n", resp.Data.Priority)
		fmt.Printf("Message:  %s\n", resp.Message)
		fmt.Printf("Sent to:  %s\n", strings.Join(sendFlags.to, ", "))
		if len(sendFlags.cc) > 0 {
			fmt.Printf("CC:       %s\n", strings.Join(sendFlags.cc, ", "))
		}
	}

	return nil
}
