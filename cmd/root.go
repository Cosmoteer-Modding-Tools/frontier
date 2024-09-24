package cmd

import (
	"fmt"
	"os"

	fr "github.com/Cosmoteer-Modding-Tools/frontier/common"
	"github.com/spf13/cobra"
)

var version fr.Version

var RootCmd = &cobra.Command{
	Use:   "Frontier",
	Short: "Frontier is a application to (hopefully) make Cosmoteer modding slightly easier",
}

func Execute(offlineMode *bool) {
	if !*offlineMode {
		needsUpdate, versions, err := fr.CheckForFrontierUpdate()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if needsUpdate {
			fmt.Println("A new Frontier version is available!(" + versions[0].Fmt() + " -> " + versions[1].Fmt() + ")\nTo install it, run `frontier upgrade`")
			os.Exit(1)
		}
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().Bool("offline", false, "If given, Frontier will not check for a new version, which needs an internet connection")
}
