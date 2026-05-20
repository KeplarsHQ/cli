package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var audiencesCmd = &cobra.Command{
	Use:   "audiences",
	Short: "Manage audiences",
}

var audiencesCreateFlags struct {
	name        string
	description string
}

var audiencesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an audience",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		result, err := client.Audiences.Create(context.Background(), audiencesCreateFlags.name, audiencesCreateFlags.description)
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

var audiencesListFlags struct {
	page  int
	limit int
}

var audiencesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List audiences",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		result, err := client.Audiences.List(context.Background(), audiencesListFlags.page, audiencesListFlags.limit)
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

var audiencesGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get an audience by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		result, err := client.Audiences.Get(context.Background(), args[0])
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

var audiencesDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete an audience",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAPIKey(); err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		result, err := client.Audiences.Delete(context.Background(), args[0])
		if err != nil {
			return err
		}
		return printResult(result)
	},
}

func init() {
	audiencesCreateCmd.Flags().StringVar(&audiencesCreateFlags.name, "name", "", "Audience name (required)")
	audiencesCreateCmd.Flags().StringVar(&audiencesCreateFlags.description, "description", "", "Audience description")
	audiencesCreateCmd.MarkFlagRequired("name")

	audiencesListCmd.Flags().IntVar(&audiencesListFlags.page, "page", 0, "Page number")
	audiencesListCmd.Flags().IntVar(&audiencesListFlags.limit, "limit", 0, "Results per page")

	audiencesCmd.AddCommand(audiencesCreateCmd, audiencesListCmd, audiencesGetCmd, audiencesDeleteCmd)
}
