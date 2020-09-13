package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spilliams/blaseball/internal"
	"github.com/spilliams/blaseball/pkg/api"
)

type contextLabel string

const (
	apiLabel contextLabel = "api"
)

func main() {
	rootCmd := newRootCmd()

	rootCmd.AddCommand(newDivisionsCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

const (
	verboseFlag   = "verbose"
	customAPIFlag = "custom-api"
	remoteAPIFlag = "remote-api"
	localAPIFlag  = "local-api"
)

// newRootCmd returns the blasball base command.
func newRootCmd() *cobra.Command {
	var flags struct {
		forbiddenKnowledge bool
		verbose            bool
		customAPI          string
		localAPI           bool
		remoteAPI          bool
	}
	cmd := &cobra.Command{
		Use:   "blase",
		Short: "A tool for getting details about blaseball",
	}

	cmd.PersistentFlags().BoolVarP(&flags.forbiddenKnowledge, "forbidden-knowledge", "f", false, "Display forbidden knowledge")
	cmd.PersistentFlags().BoolVarP(&flags.verbose, verboseFlag, "v", false, "Display verbose output")
	cmd.PersistentFlags().StringVar(&flags.customAPI, customAPIFlag, "", "Use a custom API")
	cmd.PersistentFlags().BoolVarP(&flags.localAPI, localAPIFlag, "l", false, "Use the API at http://localhost:8080/")
	cmd.PersistentFlags().BoolVarP(&flags.remoteAPI, remoteAPIFlag, "r", true, "Use the API at https://www.blaseball.com/database/")

	return cmd
}

func resolveAPI(cmd *cobra.Command) (internal.RemoteDataSession, error) {
	customAPIURL, err := cmd.Flags().GetString(customAPIFlag)
	if err != nil {
		return nil, err
	}
	localAPI, err := cmd.Flags().GetBool(localAPIFlag)
	if err != nil {
		return nil, err
	}
	remoteAPI, err := cmd.Flags().GetBool(remoteAPIFlag)
	if err != nil {
		return nil, err
	}
	// resolve API url
	apiURL := ""
	if len(customAPIURL) != 0 {
		vlog(cmd, "using custom api url %s", customAPIURL)
		apiURL = customAPIURL
	} else if localAPI {
		vlog(cmd, "using local api url (http://localhost:8080/)")
		apiURL = "http://localhost:8080/"
	} else if remoteAPI {
		vlog(cmd, "using remote api url (https://www.blaseball.com/database/)")
		apiURL = "https://www.blaseball.com/database/"
	} else {
		return nil, fmt.Errorf("no API URL specified. Please use one of --%s, --%s or --%s", customAPIFlag, localAPIFlag, remoteAPIFlag)
	}

	verbose, err := cmd.Flags().GetBool(verboseFlag)
	if err != nil {
		return nil, err
	}
	level := logrus.InfoLevel
	if verbose {
		level = logrus.DebugLevel
	}

	apiService := api.NewAPI(apiURL, level)
	return apiService, nil
}

// TODO: use logrus levels for the CLI too
func vlog(cmd *cobra.Command, format string, parts ...interface{}) {
	verbose, err := cmd.Flags().GetBool(verboseFlag)
	if err != nil {
		panic(err)
	}
	if !verbose {
		return
	}

	msg := fmt.Sprintf(format, parts...)
	if msg[len(msg)-1] != '\n' {
		msg += "\n"
	}
	fmt.Print(msg)
}

func isGUID(s string) bool {
	// TODO
	return false
}
