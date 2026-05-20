package cmd

import (
	"fmt"
	"os"

	"github.com/Swing-Technologies/keplars-email-cli/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Keplars CLI configuration",
}

var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a configuration value",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}
		switch args[0] {
		case "api-key":
			cfg.APIKey = args[1]
		case "base-url":
			cfg.BaseURL = args[1]
		default:
			return fmt.Errorf("unknown key %q (valid: api-key, base-url)", args[0])
		}
		if err := config.Save(cfg); err != nil {
			return err
		}
		fmt.Printf("Saved %s to %s\n", args[0], config.Path())
		return nil
	},
}

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Print resolved configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		resolvedKey := apiKey
		keySource := "flag/env"
		if resolvedKey == "" && cfg.APIKey != "" {
			resolvedKey = cfg.APIKey
			keySource = config.Path()
		}

		resolvedURL := baseURL
		urlSource := "flag/env"
		if resolvedURL == "" && cfg.BaseURL != "" {
			resolvedURL = cfg.BaseURL
			urlSource = config.Path()
		}

		if resolvedKey != "" {
			fmt.Printf("api-key: %s  (source: %s)\n", resolvedKey, keySource)
		} else {
			fmt.Println("api-key: <not set>")
		}

		if resolvedURL != "" {
			fmt.Printf("base-url: %s  (source: %s)\n", resolvedURL, urlSource)
		} else {
			fmt.Printf("base-url: https://api.keplars.com  (default)\n")
		}

		fmt.Printf("\nConfig file: %s", config.Path())
		if _, err := os.Stat(config.Path()); os.IsNotExist(err) {
			fmt.Print("  (not found)")
		}
		fmt.Println()
		return nil
	},
}

var configDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove the config file",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.Delete(); err != nil {
			return err
		}
		fmt.Printf("Deleted %s\n", config.Path())
		return nil
	},
}

func init() {
	configCmd.AddCommand(configSetCmd, configGetCmd, configDeleteCmd)
}
