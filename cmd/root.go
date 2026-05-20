package cmd

import (
	"fmt"
	"os"

	"github.com/KeplarsHQ/cli/internal/config"
	"github.com/spf13/cobra"
)

var (
	apiKey  string
	baseURL string
)

var rootCmd = &cobra.Command{
	Use:   "keplars",
	Short: "Keplars CLI - Manage your Keplars account from the command line",
	Long: `Keplars CLI is a command-line tool for sending emails and managing contacts,
audiences, automations, and domains via the Keplars API.

Set your API key using the KEPLARS_API_KEY environment variable, the --api-key flag,
or run: keplars config set api-key <key>`,
	Version: "0.1.0",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if apiKey != "" {
			return nil
		}
		cfg, err := config.Load()
		if err != nil {
			return err
		}
		if cfg.APIKey != "" {
			apiKey = cfg.APIKey
		}
		if baseURL == "" && cfg.BaseURL != "" {
			baseURL = cfg.BaseURL
		}
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&apiKey, "api-key", os.Getenv("KEPLARS_API_KEY"), "Keplars API key (or set KEPLARS_API_KEY env var)")
	rootCmd.PersistentFlags().StringVar(&baseURL, "base-url", os.Getenv("KEPLARS_BASE_URL"), "API base URL (optional, defaults to production)")

	rootCmd.AddCommand(sendCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(contactsCmd)
	rootCmd.AddCommand(audiencesCmd)
	rootCmd.AddCommand(automationsCmd)
	rootCmd.AddCommand(domainsCmd)
}

func checkAPIKey() error {
	if apiKey == "" {
		return fmt.Errorf("API key is required. Set KEPLARS_API_KEY environment variable, use --api-key flag, or run: keplars config set api-key <key>")
	}
	return nil
}
