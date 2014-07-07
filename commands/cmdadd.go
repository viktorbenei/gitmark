package commands

import (
// "fmt"
// "github.com/viktorbenei/gitmark/config"
)

var cmdAdd = &Command{
	Usage: "add",
	Short: "add bookmarked repositories",
	Run:   runAdd,
	Name:  "add",
}

func runAdd(cmd *Command, args []string) error {
	return nil
}
