package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var statusFlags struct {
	timeout int
	json    bool
}

var statusCmd = &cobra.Command{
	Use:   "status <email-id>",
	Short: "Get email status",
	Long: `Get the delivery status and details of a sent email.

Examples:
  # Get email status
  keplars status email_abc123

  # Get status as JSON
  keplars status email_abc123 --json`,
	Args: cobra.ExactArgs(1),
	RunE: getStatus,
}

func init() {
	statusCmd.Flags().IntVar(&statusFlags.timeout, "timeout", 30, "Request timeout in seconds")
	statusCmd.Flags().BoolVar(&statusFlags.json, "json", false, "Output response as JSON")
}

func getStatus(cmd *cobra.Command, args []string) error {
	if err := checkAPIKey(); err != nil {
		return err
	}

	emailID := args[0]

	base := "https://api.keplars.com"
	if baseURL != "" {
		base = baseURL
	}

	url := fmt.Sprintf("%s/api/v1/public/emails/get-email?id=%s", base, emailID)

	httpClient := &http.Client{Timeout: time.Duration(statusFlags.timeout) * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	output, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(output))
	return nil
}
