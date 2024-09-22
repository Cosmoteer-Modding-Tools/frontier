package cmd

import (
	//"fmt"
	//"os"
	//"path"
	//"strings"

	"github.com/spf13/cobra"
	// fr "github.com/voidwyrm-2/frontier/common"
)

var cmod = &cobra.Command{
	Use:   "cmod",
	Short: "Commands related to mod structure and management",
}

func init() {
	RootCmd.AddCommand(cmod)
}
