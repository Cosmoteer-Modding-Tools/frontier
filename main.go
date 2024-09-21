package main

import (
	"flag"

	"github.com/voidwyrm-2/frontier/cmd"
)

func main() {
	offlineModeFlag := flag.Bool("offline", false, "")
	flag.Parse()
	cmd.Execute(offlineModeFlag)
}
