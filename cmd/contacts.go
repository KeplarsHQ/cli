package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/KeplarsHQ/go-sdk/keplars"
	"github.com/spf13/cobra"
)

var contactsCmd = &cobra.Command{
	Use:   "contacts",
	Short: "Manage contacts",
}

var contactsAddFlags struct {
	email      string
	name       string
	audienceID string
}

var contactsAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a contact",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		params := map[string]interface{}{"email": contactsAddFlags.email}
		if contactsAddFlags.name != "" {
			params["name"] = contactsAddFlags.name
		}
		if contactsAddFlags.audienceID != "" {
			params["audience_id"] = contactsAddFlags.audienceID
		}
		result, err := client.Contacts.Add(context.Background(), params)
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

var contactsGetCmd = &cobra.Command{
	Use:   "get <email>",
	Short: "Get a contact by email",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		result, err := client.Contacts.Get(context.Background(), args[0])
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

var contactsListFlags struct {
	audienceID string
	page       int
	limit      int
}

var contactsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List contacts",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		result, err := client.Contacts.List(context.Background(), contactsListFlags.audienceID, contactsListFlags.page, contactsListFlags.limit)
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

var contactsUpdateFlags struct {
	name       string
	audienceID string
}

var contactsUpdateCmd = &cobra.Command{
	Use:   "update <email>",
	Short: "Update a contact",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		params := map[string]interface{}{}
		if contactsUpdateFlags.name != "" {
			params["name"] = contactsUpdateFlags.name
		}
		if contactsUpdateFlags.audienceID != "" {
			params["audience_id"] = contactsUpdateFlags.audienceID
		}
		result, err := client.Contacts.Update(context.Background(), args[0], params)
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

var contactsDeleteCmd = &cobra.Command{
	Use:   "delete <email>",
	Short: "Delete a contact",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		result, err := client.Contacts.Delete(context.Background(), args[0])
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

func init() {
	contactsAddCmd.Flags().StringVar(&contactsAddFlags.email, "email", "", "Contact email (required)")
	contactsAddCmd.Flags().StringVar(&contactsAddFlags.name, "name", "", "Contact name")
	contactsAddCmd.Flags().StringVar(&contactsAddFlags.audienceID, "audience-id", "", "Audience ID")
	contactsAddCmd.MarkFlagRequired("email")

	contactsListCmd.Flags().StringVar(&contactsListFlags.audienceID, "audience-id", "", "Filter by audience ID")
	contactsListCmd.Flags().IntVar(&contactsListFlags.page, "page", 0, "Page number")
	contactsListCmd.Flags().IntVar(&contactsListFlags.limit, "limit", 0, "Results per page")

	contactsUpdateCmd.Flags().StringVar(&contactsUpdateFlags.name, "name", "", "New name")
	contactsUpdateCmd.Flags().StringVar(&contactsUpdateFlags.audienceID, "audience-id", "", "New audience ID")

	contactsCmd.AddCommand(contactsAddCmd, contactsGetCmd, contactsListCmd, contactsUpdateCmd, contactsDeleteCmd)
}

func newClient() (*keplars.Client, error) {
	var opts []func(*keplars.Config)
	if baseURL != "" {
		opts = append(opts, func(c *keplars.Config) { c.BaseURL = baseURL })
	}
	return keplars.NewClient(apiKey, opts...)
}

func printResult(result map[string]interface{}) error {
	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
	return nil
}
