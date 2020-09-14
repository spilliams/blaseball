package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spilliams/blaseball/internal"
	"github.com/spilliams/blaseball/internal/serverdata"
)

type contextLabel string

const (
	apiLabel contextLabel = "api"

	baseAPIURLFlag = "api"
	verboseFlag    = "verbose"
)

var (
	baseAPIURL         string
	forbiddenKnowledge bool
	verbose            bool
)

func main() {
	rootCmd := newRootCmd()

	cobra.OnInitialize(initLogger)

	rootCmd.AddCommand(newDivisionsCmd())
	rootCmd.AddCommand(newTeamsCmd())
	rootCmd.AddCommand(newPlayersCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// newRootCmd returns the blasball base command.
func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "blase",
		Short: "A tool for getting details about blaseball",
	}

	cmd.PersistentFlags().StringVar(&baseAPIURL, baseAPIURLFlag, "http://localhost:8080/", "The base URL of the API to use")
	cmd.PersistentFlags().BoolVarP(&forbiddenKnowledge, "forbidden-knowledge", "f", false, "Display forbidden knowledge")
	cmd.PersistentFlags().BoolVarP(&verbose, verboseFlag, "v", false, "Display verbose output")

	return cmd
}

func initLogger() {
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func resolveAPI(cmd *cobra.Command) (internal.ServerDataSession, error) {
	if len(baseAPIURL) == 0 {
		return nil, fmt.Errorf("no API URL specified. Please use --%s", baseAPIURLFlag)
	}

	apiService := serverdata.NewAPI(baseAPIURL, logrus.GetLevel())
	return apiService, nil
}

func isGUID(s string) bool {
	guidRE := regexp.MustCompile("[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}")
	return guidRE.MatchString(s)
}
