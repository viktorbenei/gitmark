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
}

var GitmarkConfig Config

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

func ReadConfigFromFile() error {
	tryConfigFile := func(filepath string) error {
		if err := tryToReadConfigFile(filepath); err != nil {
			return err
		}
		log.Println(" (i) Using config file:", filepath)
		return nil
	}
	configFilePathes := []string{".gitmarkrc.json", "~/.gitmarkrc.json"}
	configRead := false

	for _, aConfFile := range configFilePathes {
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
