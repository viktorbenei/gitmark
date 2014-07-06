package commands

import (
	"errors"
	"fmt"
	"github.com/viktorbenei/gitmark/config"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

var cmdScan = &Command{
	Usage: "scan",
	Short: "scan for local repositories which are not yet bookmarked",
	Run:   runScan,
	Name:  "scan",
}

var (
	isStore       bool
	scanRoot      string
	isScanVerbose bool
	scanIgnores   []string
)

func init() {
	cmdScan.Flag.StringVar(&scanRoot, "rootpath", "", "[required] root path of scan")
	cmdScan.Flag.BoolVar(&isStore, "store", false, "store found repositories in bookmarks?")
	cmdScan.Flag.BoolVar(&isScanVerbose, "verbose", false, "verbose?")
}

func generateRepositoryTitleForRepositoryPath(repoPath string) string {
	parentDirName := path.Base(repoPath)
	return parentDirName
}

func runScan(cmd *Command, args []string) error {
	if isScanVerbose {
		fmt.Println("scanning...")
		fmt.Println(" (i) Using ignores:")
		scanIgnores = config.GitmarkConfig.ScanIgnores
		fmt.Println(scanIgnores)
	}

	if scanRoot == "" {
		return errors.New("No scan rootpath provided")
	}

	err := filepath.Walk(scanRoot, func(aPath string, aFileInfo os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err) // can't walk here,
			return nil       // but continue walking elsewhere
		}
		if !aFileInfo.IsDir() {
			return nil // not a dir.  ignore.
		}
		matched, err := filepath.Match("*.git", aFileInfo.Name())
		if err != nil {
			fmt.Println(err) // malformed pattern
			return err       // this is fatal.
		}
		if matched {
			isPathIsMatch := true
			if scanIgnores != nil && len(scanIgnores) > 0 {
				for _, aScanIgnore := range scanIgnores {
					matchedIgnore, err := regexp.MatchString(aScanIgnore, aPath)
					if err != nil {
						fmt.Println(" (!)", err) // malformed pattern
					}
					if matchedIgnore {
						if isScanVerbose {
							fmt.Println(" (i) ignore match:", aPath, "|", aScanIgnore)
						}
						isPathIsMatch = false
						break
					}
				}
			}
			if isPathIsMatch {
				repoPath := path.Dir(aPath)
				fmt.Println(repoPath)
				if config.GitmarkConfig.IsRepositoryPathStored(repoPath) {
					if isScanVerbose {
						fmt.Println(" (i) Path already stored - ignoring:", repoPath)
					}
				} else {
					repo := config.Repository{Title: generateRepositoryTitleForRepositoryPath(repoPath), Path: repoPath}
					config.GitmarkConfig.AddRepository(repo)
					if isScanVerbose {
						fmt.Println(" (+) Path added to bookmarks:", repo)
					}
				}
			}
		}
		return nil
	})

	if isScanVerbose {
		formattedJsonBytes, err := config.GitmarkConfig.GenerateFormattedJSON()
		if err != nil {
			fmt.Println("Failed to generate JSON:", err)
		}
		fmt.Printf("Config: %s\n", formattedJsonBytes)
	}

	return err
}
