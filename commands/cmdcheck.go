package commands

import (
	"fmt"
	"github.com/viktorbenei/gitmark/config"
	"os/exec"
)

var cmdCheck = &Command{
	Usage: "check",
	Short: "check bookmarked repositories",
	Run:   runCheck,
	Name:  "check",
}

var (
	isVerbose = false
)

func init() {
	cmdCheck.Flag.BoolVar(&isVerbose, "verbose", false, "verbose?")
}

func runCheck(cmd *Command, args []string) error {
	for _, repo := range config.GitmarkConfig.Repositories {
		c := exec.Command("git", "status", "--porcelain")
		c.Dir = repo.Path
		if outp, err := c.Output(); err != nil {
			return err
		} else {
			if len(outp) > 0 {
				fmt.Println("->", repo.Title)
				fmt.Printf("%s", outp)
				fmt.Println()
			} else if isVerbose {
				fmt.Println("->", repo.Title)
				fmt.Println(" (i) Nothing to commit")
				fmt.Println()
			}
		}
	}
	return nil
}
