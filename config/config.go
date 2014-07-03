package config

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
}

func (c *Config) tryToReadConfigFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func (c *Config) ReadConfigsFromFile() error {
	tryConfigFile := func(filepath string) error {
		err := c.tryToReadConfigFile(filepath)
		if err != nil {
			return err
		}
		fmt.Println(" (i) Using config file:", filepath)
		return nil
	}
	configFilePathes := []string{".gitmarkrc.json", "~/.gitmarkrc.json"}
	configRead := false

	for _, aConfFile := range configFilePathes {
		if !configRead {
			err := tryConfigFile(aConfFile)
			if err == nil {
				configRead = true
				return nil
			}
		}
	}
	return errors.New("Could not find config file")
}
