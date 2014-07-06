package commands

import (
	"errors"
	"github.com/viktorbenei/gitmark/config"
	"os/exec"
	"strings"
)

var cmdOpen = &Command{
	Usage: "open",
	Short: "open bookmarked repository (by title)",
	Run:   runOpen,
	Name:  "open",
}

func runOpen(cmd *Command, args []string) error {
	if len(args) != 1 {
		return errors.New("Only 1 argument is accepted, and it should be the title of a bookmarked repository")
	}
	repoTitle := args[0]

	repo, err := config.GitmarkConfig.GetRepositoryByTitle(repoTitle)
	if err != nil {
		return err
	}

	openCommandString := config.GitmarkConfig.Preferences.OpenCommand
	if openCommandString == "" {
		return errors.New("No open-command found in preferences - required for open")
	}

	openCompontents := strings.Split(openCommandString, " ")

	var openArgs []string
	if len(openCompontents) > 1 {
		openArgs = append(openCompontents[1:], repo.Path)
	} else {
		openArgs = []string{repo.Path}
	}

	c := exec.Command(openCompontents[0], openArgs...)
	if err := c.Run(); err != nil {
		return err
	}

	return nil
}
