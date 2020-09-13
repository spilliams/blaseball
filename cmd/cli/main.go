package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spilliams/blaseball/pkg/api"
)

type contextLabel string

const (
	apiLabel contextLabel = "api"
)

func main() {
	rootCmd := newRootCmd()

	// TODO: consider putting this in config
	api := api.NewAPI("http://localhost:8080/")
	ctx := context.WithValue(context.Background(), apiLabel, api)

	rootCmd.AddCommand(newDivisionsCmd())

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// newRootCmd returns the blasball base command.
func newRootCmd() *cobra.Command {
	var flags struct {
		forbiddenKnowledge bool
		verbose            bool
	}
	cmd := &cobra.Command{
		Use:   "blase",
		Short: "A tool for getting details about blaseball",
		// Uncomment the following if the root cmd should have its own logic
		// that happens when someone runs it with no subcommand
		// RunE: func(cmd *cobra.Command, args []string) error {
		// 	return nil
		// },
	}

	cmd.PersistentFlags().BoolVarP(&flags.forbiddenKnowledge, "forbidden-knowledge", "f", false, "Display forbidden knowledge")
	cmd.PersistentFlags().BoolVarP(&flags.verbose, "verbose", "v", false, "Display verbose output")

	return cmd
}