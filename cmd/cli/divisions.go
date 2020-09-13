package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spilliams/blaseball/pkg/model"
)

func newDivisionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "divisions",
		Aliases: []string{"d", "division"},
	}

	cmd.AddCommand(newDivisionsListCmd())
	cmd.AddCommand(newDivisionGetCmd())

	return cmd
}

func newDivisionsListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all Divisions",
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := resolveAPI(cmd)
			if err != nil {
				return err
			}
			divs, err := api.GetAllDivisions()
			if err != nil {
				return err
			}
			fmt.Println(divs)
			return nil
		},
	}
}

func newDivisionGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <id_or_name>",
		Short: "Get a single Division",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := resolveAPI(cmd)
			if err != nil {
				return err
			}
			var div *model.Division
			if isGUID(args[0]) {
				div, err = api.GetDivisionByID(args[0])
			} else {
				div, err = api.GetDivisionByName(args[0])
			}
			if err != nil {
				return err
			}
			fmt.Println(div)
			return nil
		},
	}
}
