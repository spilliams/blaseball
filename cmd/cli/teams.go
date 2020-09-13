package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spilliams/blaseball/pkg/model"
)

func newTeamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "teams",
		Aliases: []string{"t", "team"},
	}

	cmd.AddCommand(newTeamsListCmd())
	cmd.AddCommand(newTeamGetCmd())

	return cmd
}

func newTeamsListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all teams",
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := resolveAPI(cmd)
			if err != nil {
				return err
			}
			teams, err := api.GetAllTeams()
			if err != nil {
				return err
			}
			fmt.Println(teams)
			return nil
		},
	}
}

func newTeamGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <id_or_name>",
		Short: "Get a single Team",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := resolveAPI(cmd)
			if err != nil {
				return err
			}
			var team *model.Team
			if isGUID(args[0]) {
				team, err = api.GetTeamByID(args[0])
			} else {
				team, err = api.GetTeamByFullName(args[0])
				if err != nil {
					team, err = api.GetTeamByNickname(args[0])
				}
			}
			if err != nil {
				return err
			}
			fmt.Println(team)
			return nil
		},
	}
}
