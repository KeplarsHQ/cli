package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var automationsCmd = &cobra.Command{
	Use:   "automations",
	Short: "Manage automations",
}

var automationsListFlags struct {
	page  int
	limit int
}

var automationsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List automations",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		result, err := client.Automations.List(context.Background(), automationsListFlags.page, automationsListFlags.limit)
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

var automationsGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get an automation by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		result, err := client.Automations.Get(context.Background(), args[0])
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

var automationsEnrollFlags struct {
	email string
}

var automationsEnrollCmd = &cobra.Command{
	Use:   "enroll <id>",
	Short: "Enroll a contact in an automation",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		result, err := client.Automations.Enroll(context.Background(), args[0], automationsEnrollFlags.email)
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

var automationsUnenrollFlags struct {
	email string
}

var automationsUnenrollCmd = &cobra.Command{
	Use:   "unenroll <id>",
	Short: "Unenroll a contact from an automation",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		result, err := client.Automations.Unenroll(context.Background(), args[0], automationsUnenrollFlags.email)
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

func init() {
	automationsListCmd.Flags().IntVar(&automationsListFlags.page, "page", 0, "Page number")
	automationsListCmd.Flags().IntVar(&automationsListFlags.limit, "limit", 0, "Results per page")

	automationsEnrollCmd.Flags().StringVar(&automationsEnrollFlags.email, "email", "", "Contact email (required)")
	automationsEnrollCmd.MarkFlagRequired("email")

	automationsUnenrollCmd.Flags().StringVar(&automationsUnenrollFlags.email, "email", "", "Contact email (required)")
	automationsUnenrollCmd.MarkFlagRequired("email")

	automationsCmd.AddCommand(automationsListCmd, automationsGetCmd, automationsEnrollCmd, automationsUnenrollCmd)
}
