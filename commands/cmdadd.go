package commands

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/viktorbenei/gitmark/config"
	"github.com/viktorbenei/gitmark/pathutil"
	"os"
	"os/exec"
	"strings"
)

var cmdAdd = &Command{
	Usage: "add",
	Short: "add bookmarked repositories",
	Run:   runAdd,
	Name:  "add",
}

func generateTitleFromRepositoryUrl(repositoryUrl string) (string, error) {
	trimmedUrl := strings.TrimSpace(repositoryUrl)
	if trimmedUrl == "" {
		return "", errors.New("Invalid repository url: empty url")
	}
	repoUrlCompontents := strings.Split(trimmedUrl, "/")
	if len(repoUrlCompontents) < 1 {
		return "", errors.New("")
	}
	extensionComps := strings.Split(repoUrlCompontents[len(repoUrlCompontents)-1], ".")
	return extensionComps[0], nil
}

func generateLocalPathFromUserInputPathAndRepositoryUrl(userInputPath, repositoryUrl string) (string, error) {
	titleFromUrl, err := generateTitleFromRepositoryUrl(repositoryUrl)
	if err != nil {
		return "", err
	}

	var localPath string
	if userInputPath[len(userInputPath)-1:] == "/" {
		localPath = userInputPath + titleFromUrl
	} else {
		localPath = userInputPath
	}
	return localPath, nil
}

func cloneRepository(repositoryUrl string, r config.Repository) error {
	fmt.Println(" (i) Cloning...")
	c := exec.Command("git", "clone", repositoryUrl, r.Path)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Start(); err != nil {
		return err
	}
	if err := c.Wait(); err != nil {
		return err
	}
	return nil
}

func runAdd(cmd *Command, args []string) error {
	fmt.Println(" (i) Adding new repository")

	fmt.Printf("-> Repository url: ")
	bio := bufio.NewReader(os.Stdin)
	lineBytes, hasMoreInLine, err := bio.ReadLine()
	if err != nil {
		return err
	}
	if hasMoreInLine {
		return errors.New("Input too long!")
	}
	repoUrl := fmt.Sprintf("%s", lineBytes)
	if len(repoUrl) < 1 {
		return errors.New("URL can't be empty!")
	}

	defaultTitle, err := generateTitleFromRepositoryUrl(repoUrl)
	if err != nil {
		fmt.Println("Can't generate a default title for the url")
	}
	fmt.Printf("-> Bookmark Title [%s]: ", defaultTitle)
	bio = bufio.NewReader(os.Stdin)
	lineBytes, hasMoreInLine, err = bio.ReadLine()
	if err != nil {
		return err
	}
	if hasMoreInLine {
		return errors.New("Input too long!")
	}
	repoTitle := fmt.Sprintf("%s", lineBytes)
	if repoTitle == "" {
		// use default
		repoTitle = defaultTitle
	}

	fmt.Printf("-> Local Path: ")
	bio = bufio.NewReader(os.Stdin)
	lineBytes, hasMoreInLine, err = bio.ReadLine()
	if err != nil {
		return err
	}
	if hasMoreInLine {
		return errors.New("Input too long!")
	}
	repoLocalPath := fmt.Sprintf("%s", lineBytes)
	if len(repoLocalPath) < 1 {
		return errors.New("Path can't be empty!")
	}
	repoLocalPath = pathutil.ExpandPath(repoLocalPath)
	if repoLocalPath == "" {
		return errors.New("Failed to expand the provided path.")
	}
	repoLocalPath, err = generateLocalPathFromUserInputPathAndRepositoryUrl(repoLocalPath, repoUrl)
	if err != nil {
		return err
	}

	fmt.Println("--- Adding:")
	fmt.Println("Repository.URL:", repoUrl)
	fmt.Println("Repository.Title:", repoTitle)
	fmt.Println("Repository.LocalPath:", repoLocalPath)
	fmt.Println("-----------")

	// check whether the folder is a local working copy of the repository
	isCloneNeeded := false
	expectedGitDirPath := fmt.Sprintf("%s/.git", repoLocalPath)
	fmt.Println(" (i) Checking for .git at", expectedGitDirPath)
	if _, err := os.Stat(expectedGitDirPath); err != nil {
		if os.IsNotExist(err) {
			isCloneNeeded = true
		}
	}

	r := config.Repository{Title: repoTitle, Path: repoLocalPath}
	if isCloneNeeded {
		if err := cloneRepository(repoUrl, r); err != nil {
			return err
		}
	} else {
		fmt.Println(" (i) Repository found locally - no clone needed")
	}

	fmt.Println(" (i) Storing into Bookmarks...")
	if err := config.GitmarkConfig.AddRepository(r); err != nil {
		return err
	}
	if err := config.WriteGitmarkConfigToFile(); err != nil {
		return err
	}

	fmt.Println(" (i) Added to gitmarks")

	return nil
}
