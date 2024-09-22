package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	fr "github.com/voidwyrm-2/frontier/common"
)

var cmod_init = &cobra.Command{
	Use:   "init",
	Short: "Generates the boilerplate for a Cosmoteer mod",
	Run: func(cmd *cobra.Command, args []string) {
		pprefix := ""
		if len(args) > 0 {
			ar := strings.TrimSpace(args[0])
			if ar != "" {
				_, oe, err := fr.RunCommand("mkdir "+ar, []string{})
				if oe != "" {
					fmt.Println(oe)
					os.Exit(1)
				} else if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
				pprefix = ar
			}
		}

		if !fr.DoesItemExist(path.Join(pprefix, "mod.rules")) {
			// TODO: add text prompts
		}
	},
}

func init() {
	cmod.AddCommand(cmod_init)
}
