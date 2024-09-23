package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	fr "github.com/Cosmoteer-Modding-Tools/frontier/common"
	"github.com/spf13/cobra"
)

var cmod_init = &cobra.Command{
	Use:   "init",
	Short: "Generates the boilerplate for a Cosmoteer mod",
	Run: func(cmd *cobra.Command, args []string) {
		pprefix := ""
		if len(args) > 0 {
			ar := strings.TrimSpace(args[0])
			if ar != "" {
				if !fr.DoesItemExist(ar) {
					_, oe, err := fr.RunCommand("mkdir "+ar, []string{})
					if oe != "" {
						fmt.Println(oe)
						os.Exit(1)
					} else if err != nil {
						fmt.Println(err.Error())
						os.Exit(1)
					}
				}
				pprefix = ar
			}
		}

		if !fr.DoesItemExist(path.Join(pprefix, "mod.rules")) {
			modID := fr.ReadNonemptyLine("What is the ID of your mod?\n", "The mod ID cannot be empty")
			modName := fr.ReadNonemptyLine("What is the name of your mod?\n", "The mod name cannot be empty")
			modStringsFolder := fr.ReadNonemptyLine("What is the strings folder of your mod?\n", "The mod strings folder path cannot be empty")
			modAuthor := fr.ReadNonemptyLine("What is the author name of your mod?\n", "The mod author name cannot be empty")

			final := []string{}
			final = append(final, fmt.Sprintf("ID = \"%s\"", modID))
			final = append(final, fmt.Sprintf("Name = \"%s\"", modName))
			final = append(final, "Version = 1.0.0")
			final = append(final, "CompatibleGameVersions = [\"0.27.1\"]")
			final = append(final, "ModifiesMultiplayer = true")
			final = append(final, fmt.Sprintf("StringFolder = \"%s\"", modStringsFolder))
			final = append(final, fmt.Sprintf("Author = \"%s\"", modAuthor))
			final = append(final, "Description = \"\"")
			final = append(final, `
Actions
[
    {
        Action = AddMany
        AddTo = "<ships/terran/terran.rules>/Terran/Parts"
        ManyToAdd =
        [
            
        ]
    }
]
`)
			err := fr.WriteFile(path.Join(pprefix, "mod.rules"), strings.Join(final, "\n"))
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			if !fr.DoesItemExist(path.Join(pprefix, modStringsFolder)) {
				_, oe, err := fr.RunCommand("mkdir "+path.Join(pprefix, modStringsFolder), []string{})
				if oe != "" {
					fmt.Println(oe)
					os.Exit(1)
				} else if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
			}

			err = fr.WriteFile(path.Join(pprefix, modStringsFolder, "en.rules"), "Parts\n{\n}")
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	},
}

func init() {
	cmod.AddCommand(cmod_init)
}
