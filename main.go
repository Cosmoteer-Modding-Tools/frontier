package main

import (
	"flag"

	"github.com/Cosmoteer-Modding-Tools/frontier/cmd"
)

func main() {
	offlineModeFlag := flag.Bool("offline", false, "")
	flag.Parse()
	cmd.Execute(offlineModeFlag)
}
