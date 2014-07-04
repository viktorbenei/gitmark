package commands

import (
	"fmt"
	"github.com/viktorbenei/gitmark/config"
)

var cmdList = &Command{
	Usage: "list",
	Short: "list bookmarked repositories",
	Run:   runList,
	Name:  "list",
}

func runList(cmd *Command, args []string) error {
	fmt.Println("Repositories:")
	for idx, repo := range config.GitmarkConfig.Repositories {
		fmt.Printf(" [%d] %s (%s)\n", idx, repo.Title, repo.Path)
	}
	return nil
}
