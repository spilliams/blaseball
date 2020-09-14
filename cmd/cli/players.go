package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spilliams/blaseball/pkg/model"
)

func newPlayersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "players",
		Aliases: []string{"p", "player"},
	}

	cmd.AddCommand(newPlayersListCmd())
	cmd.AddCommand(newPlayerGetCmd())

	return cmd
}

func newPlayersListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all Players",
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := resolveAPI(cmd)
			if err != nil {
				return err
			}
			players, err := api.GetAllPlayers(forbiddenKnowledge)
			if err != nil {
				return err
			}
			fmt.Println(players)
			return nil
		},
	}
}

func newPlayerGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <id_or_name>",
		Short: "Get a single Player",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := resolveAPI(cmd)
			if err != nil {
				return err
			}
			var player *model.Player
			if isGUID(args[0]) {
				var players *model.PlayerList
				players, err = api.GetPlayersByID([]string{args[0]}, forbiddenKnowledge)
				if players != nil && len(players.List) > 0 {
					player = players.List[0]
				}
			} else {
				player, err = api.GetPlayerByName(args[0], forbiddenKnowledge)
			}
			if err != nil {
				return err
			}
			fmt.Println(player)
			return nil
		},
	}
}
