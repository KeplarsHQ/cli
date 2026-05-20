package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var domainsCmd = &cobra.Command{
	Use:   "domains",
	Short: "Manage sending domains",
}

var domainsAddCmd = &cobra.Command{
	Use:   "add <domain>",
	Short: "Add a domain",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		result, err := client.Domains.Add(context.Background(), args[0])
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

var domainsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List domains",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		result, err := client.Domains.List(context.Background())
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

var domainsStatusCmd = &cobra.Command{
	Use:   "status <domain-id>",
	Short: "Get domain verification status",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		result, err := client.Domains.GetStatus(context.Background(), args[0])
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

var domainsVerifyCmd = &cobra.Command{
	Use:   "verify <domain-id>",
	Short: "Verify a domain",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		result, err := client.Domains.Verify(context.Background(), args[0])
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

var domainsDeleteCmd = &cobra.Command{
	Use:   "delete <domain-id>",
	Short: "Delete a domain",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		result, err := client.Domains.Delete(context.Background(), args[0])
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

var domainsCreateAPIKeyFlags struct {
	name string
}

var domainsCreateAPIKeyCmd = &cobra.Command{
	Use:   "create-api-key <domain-id>",
	Short: "Create an API key for a domain",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		params := map[string]interface{}{"domain_id": args[0]}
		if domainsCreateAPIKeyFlags.name != "" {
			params["name"] = domainsCreateAPIKeyFlags.name
		}
		result, err := client.Domains.CreateAPIKey(context.Background(), params)
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

func init() {
	domainsCreateAPIKeyCmd.Flags().StringVar(&domainsCreateAPIKeyFlags.name, "name", "", "API key name")

	domainsCmd.AddCommand(domainsAddCmd, domainsListCmd, domainsStatusCmd, domainsVerifyCmd, domainsDeleteCmd, domainsCreateAPIKeyCmd)
}
