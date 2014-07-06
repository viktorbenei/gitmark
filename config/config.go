package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

// ---------------------
// --- Models

type Repository struct {
	Title string `json:"title"`
	Path  string `json:"path"`
}

type Config struct {
	Repositories []Repository `json:"repositories"`
	ScanIgnores  []string     `json:"scanignores"`
}

var GitmarkConfig Config
var ConfigFileSearchPathes = []string{".gitmarkrc.json", "~/.gitmarkrc.json"}

// ---------------------
// --- Functions

func tryToReadConfigFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	var config Config
	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&config); err != nil {
		return err
	}
	GitmarkConfig = config

	return nil
}

func (c *Config) GetRepositoryPaths() []string {
	repoCount := len(c.Repositories)
	repoPathes := make([]string, repoCount, repoCount)
	for idx, aRepo := range c.Repositories {
		repoPathes[idx] = aRepo.Path
	}
	return repoPathes
}

func ReadConfigFromFile() error {
	tryConfigFile := func(filepath string) error {
		if err := tryToReadConfigFile(filepath); err != nil {
			return err
		}
		log.Println(" (i) Using config file:", filepath)
		return nil
	}
	configRead := false

	for _, aConfFile := range ConfigFileSearchPathes {
		if !configRead {
			if err := tryConfigFile(aConfFile); err == nil {
				configRead = true
				return nil
			} else {
				log.Println(err)
			}
		}
	}
	return errors.New("Could not find config file")
}
