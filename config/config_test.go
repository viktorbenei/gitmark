package config

import (
	"testing"
)

func TestInitialState(t *testing.T) {
	t.Log("Initial state checks")
	if len(GitmarkConfig.Repositories) != 0 {
		t.Error("Initial GitmarkConfig.Repositories length should be 0")
	}

	// check search pathes
	if ConfigFileSearchPathes[0] != ".gitmarkrc.json" || ConfigFileSearchPathes[1] != "~/.gitmarkrc.json" {
		t.Error("ConfigFileSearchPathes check failed")
	}
}

func TestLoadConfigFile(t *testing.T) {
	t.Log("Should be able to load the local config file")

	// prepend the config file search path with a test one
	ConfigFileSearchPathes = append([]string{".gitmarkrc.json-test"}, ConfigFileSearchPathes...)
	if ConfigFileSearchPathes[0] != ".gitmarkrc.json-test" {
		t.Error("Failed to add the test Gitmarkrc to search path")
	}

	// load the config file
	err := ReadConfigFromFile()
	if err != nil {
		t.Error("Failed to load the config file")
	}
}

func TestAfterLoadState(t *testing.T) {
	t.Log("After load state checks")

	// now check it's loaded
	if len(GitmarkConfig.Repositories) != 2 {
		t.Error("GitmarkConfig.Repositories length should be 2")
	}

	secondRepo := GitmarkConfig.Repositories[1]
	if secondRepo.Title != "test/repo2" || secondRepo.Path != "/path/to/test2" {
		t.Error("Repository value checking failed")
	}
}