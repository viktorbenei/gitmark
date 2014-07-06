package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	//
	ConfigFilePath string `json:"-"`
	// Lookup Maps
	lookupRepositoryPaths   map[string]bool       `json:"-"`
	lookupRepositoryByTitle map[string]Repository `json:"-"`
}

var GitmarkConfig Config
var ConfigFileSearchPathes = []string{".gitmarkrc.json", "~/.gitmarkrc.json"}

// ---------------------
// --- Config Functions

func (c *Config) generateLookupMaps() error {
	lookupRepoPaths := make(map[string]bool)
	lookupRepositoryByTitle := make(map[string]Repository)
	for _, aRepo := range c.Repositories {
		if lookupRepoPaths[aRepo.Path] {
			return errors.New(fmt.Sprintf("Repository path already found: %s", aRepo.Path))
		}
		lookupRepoPaths[aRepo.Path] = true
		//
		if _, ok := lookupRepositoryByTitle[aRepo.Title]; ok {
			return errors.New(fmt.Sprintf("Repository title already found: %s", aRepo.Title))
		}
		lookupRepositoryByTitle[aRepo.Title] = aRepo
	}
	c.lookupRepositoryPaths = lookupRepoPaths
	c.lookupRepositoryByTitle = lookupRepositoryByTitle
	return nil
}

func (c *Config) IsRepositoryPathStored(repositoryPath string) bool {
	return c.lookupRepositoryPaths[repositoryPath]
}

func (c *Config) GetRepositoryPaths() []string {
	repoCount := len(c.Repositories)
	repoPathes := make([]string, repoCount, repoCount)
	for idx, aRepo := range c.Repositories {
		repoPathes[idx] = aRepo.Path
	}
	return repoPathes
}

func (c *Config) GetRepositoryByTitle(repoTitle string) (Repository, error) {
	repo, isFound := c.lookupRepositoryByTitle[repoTitle]
	if !isFound {
		return Repository{}, errors.New(fmt.Sprintf("Repository not found: %s", repoTitle))
	}
	return repo, nil
}

func (c *Config) AddRepository(repo Repository) {
	c.Repositories = append(c.Repositories, repo)
	c.generateLookupMaps()
}

func (c *Config) GenerateFormattedJSON() ([]byte, error) {
	jsonContBytes, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return []byte{}, err
	}
	return jsonContBytes, nil
}

// ---------------------
// --- Functions

func readConfigFile(filepath string) (Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	return readConfigFromReader(file)
}

func writeConfigToFile(c Config) error {
	if c.ConfigFilePath == "" {
		return errors.New("No ConfigFilePath found")
	}

	file, err := os.Create(c.ConfigFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonContBytes, err := c.GenerateFormattedJSON()
	if err != nil {
		return err
	}

	_, err = file.Write(jsonContBytes)
	if err != nil {
		return err
	}

	return nil
}

func readConfigFromReader(reader io.Reader) (Config, error) {
	var config Config
	jsonParser := json.NewDecoder(reader)
	if err := jsonParser.Decode(&config); err != nil {
		return Config{}, err
	}
	config.generateLookupMaps()

	return config, nil
}

func WriteGitmarkConfigToFile() error {
	err := writeConfigToFile(GitmarkConfig)
	if err != nil {
		log.Println(" [!] Failed to write the Config into file:", err)
		return err
	} else {
		log.Println(" (i) Written to file [OK]:", GitmarkConfig.ConfigFilePath)
	}
	return nil
}

func ReadGitmarkConfigFromFile() error {
	tryConfigFile := func(filepath string) error {
		config, err := readConfigFile(filepath)
		if err != nil {
			return err
		}
		log.Println(" (i) Using config file:", filepath)
		GitmarkConfig = config
		GitmarkConfig.ConfigFilePath = filepath
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
