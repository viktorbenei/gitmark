package commands

import (
	"log"
)

var cmdList = &Command{
	Usage: "list",
	Short: "list bookmarked repositories",
	Run:   runList,
	Name:  "list",
}

func runList(cmd *Command, args []string) {
	log.Println("-> cmd list")
}
