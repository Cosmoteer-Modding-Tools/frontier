package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	fr "github.com/voidwyrm-2/frontier/common"
)

func getFrontierVersion() (string, error) {
	res, err := http.Get("https://raw.githubusercontent.com/Cosmoteer-Modding-Tools/frontier/version.txt")
	if err != nil {
		return "", err
	}

	version, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	} else if string(version) == "404: Not Found" {
		return "", err
	}

	return string(version), nil
}

var (
	version     fr.Version
	offlineMode bool = false
)

var RootCmd = &cobra.Command{
	Use:   "Frontier",
	Short: "Frontier is a application to (hopefully) make Cosmoteer modding slightly easier",
}

func Execute() {
	dir, err := fr.GetHomeDir()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var remoteVersion fr.Version

	if !offlineMode {
		remoteVersionText, err := getFrontierVersion()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		remoteVersion, err = fr.NewVersionFromVersionString(remoteVersionText)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	vpath := path.Join(path.Clean(dir), ".frontierversion")
	content, err := fr.ReadFile(vpath)
	if err != nil {
		if fr.ErrorIsFile404(err) && !offlineMode {
			fr.WriteFile(vpath, remoteVersion.Fmt())
		} else {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	localVersion, err := fr.NewVersionFromVersionString(strings.TrimSpace(content))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if offlineMode {
		remoteVersion = localVersion
	}

	if localVersion.Compare(remoteVersion) == -1 {
		fmt.Println("A new Frontier version is available!\nTo install it, run `go install github.com/Cosmoteer-Modding-Tools/frontier@latest`")
		os.Exit(1)
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
}
