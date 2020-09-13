package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/api"
)

type contextLabel string

const (
	apiLabel contextLabel = "api"

	customAPIFlag = "custom-api"
	localAPIFlag  = "local-api"
	remoteAPIFlag = "remote-api"
	verboseFlag   = "verbose"
)

var (
	customAPIURL       string
	forbiddenKnowledge bool
	localAPI           bool
	remoteAPI          bool
	verbose            bool
)

func main() {
	rootCmd := newRootCmd()

	cobra.OnInitialize(initLogger)

	rootCmd.AddCommand(newDivisionsCmd())

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

	cmd.PersistentFlags().StringVar(&customAPIURL, customAPIFlag, "", "Use a custom API")
	cmd.PersistentFlags().BoolVarP(&forbiddenKnowledge, "forbidden-knowledge", "f", false, "Display forbidden knowledge")
	cmd.PersistentFlags().BoolVarP(&localAPI, localAPIFlag, "l", false, "Use the API at http://localhost:8080/")
	cmd.PersistentFlags().BoolVarP(&remoteAPI, remoteAPIFlag, "r", true, "Use the API at https://www.blaseball.com/database/")
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

func resolveAPI(cmd *cobra.Command) (pkg.RemoteDataSession, error) {
	// resolve API url
	apiURL := ""
	if len(customAPIURL) != 0 {
		logrus.Debugf("using custom api url %s", customAPIURL)
		apiURL = customAPIURL
	} else if localAPI {
		logrus.Debug("using local api url (http://localhost:8080/)")
		apiURL = "http://localhost:8080/"
	} else if remoteAPI {
		logrus.Debug("using remote api url (https://www.blaseball.com/database/)")
		apiURL = "https://www.blaseball.com/database/"
	} else {
		return nil, fmt.Errorf("no API URL specified. Please use one of --%s, --%s or --%s", customAPIFlag, localAPIFlag, remoteAPIFlag)
	}

	apiService := api.NewAPI(apiURL, logrus.GetLevel())
	return apiService, nil
}

func isGUID(s string) bool {
	guidRE := regexp.MustCompile("[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}")
	return guidRE.MatchString(s)
}
