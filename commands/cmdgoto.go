package commands

import (
	"errors"
	"fmt"
	"github.com/viktorbenei/gitmark/config"
)

var cmdGoto = &Command{
	Usage: "goto",
	Short: "goto bookmarked repository (by title)",
	Run:   runGoto,
	Name:  "goto",
}

func runGoto(cmd *Command, args []string) error {
	if len(args) != 1 {
		return errors.New("Only 1 argument is accepted, and it should be the title of a bookmarked repository")
	}
	repoTitle := args[0]

	repo, err := config.GitmarkConfig.GetRepositoryByTitle(repoTitle)
	if err != nil {
		return err
	}

	fmt.Println(repo)

	return nil
}
