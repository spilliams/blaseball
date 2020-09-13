package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spilliams/blaseball/internal"
)

func newDivisionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "divisions",
		Aliases: []string{"d"},
		// RunE: func(cmd *cobra.Command, args []string) error {
		// 	return nil
		// },
	}

	cmd.AddCommand(newDivisionsListCmd())

	return cmd
}

func newDivisionsListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all divisions",
		RunE: func(cmd *cobra.Command, args []string) error {
			api := cmd.Context().Value(apiLabel).(internal.RemoteDataSession)
			divs, err := api.GetAllDivisions()
			if err != nil {
				return err
			}
			fmt.Println(divs)
			return nil
		},
	}
}