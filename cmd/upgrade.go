package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	fr "github.com/Cosmoteer-Modding-Tools/frontier/common"
	"github.com/spf13/cobra"
)

var upgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Updates Frontier",
	Run: func(cmd *cobra.Command, args []string) {
		_, oe, err := fr.RunCommand("go install \"github.com/Cosmoteer-Modding-Tools/frontier@latest\"", []string{})
		if oe != "" {
			fmt.Println("error while updating Frontier:")
			fmt.Println(oe)
			os.Exit(1)
		} else if err != nil {
			fmt.Println("error while updating Frontier:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		remoteVersionText, err := getFrontierVersion()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		remoteVersion, err := fr.NewVersionFromVersionString(strings.TrimSpace(remoteVersionText))
		if err != nil {
			fmt.Println("error while fetching new Frontier version, version fetched is invalid, please contact one of the maintainers\nactual error: " + err.Error())
			os.Exit(1)
		}

		homedir, err := fr.GetHomeDir()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		vpath := path.Join(path.Clean(homedir), ".frontierversion")
		if err != nil {
			if fr.ErrorIsFile404(err) {
				fr.WriteFile(vpath, remoteVersion.Fmt())
			} else {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(upgrade)
}
