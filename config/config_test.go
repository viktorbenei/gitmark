package config

import (
	"fmt"
	"strings"
	"testing"
)

var (
	testConfigFilePath = ".gitmarkrc.json-test"
	repoTitle1         = "test/repo1"
	repoPath1          = "/path/to/test1"
	repoTitle2         = "test/repo2"
	repoPath2          = "/path/to/test2"
	newTestRepository  = Repository{Title: "new-test-repo", Path: "/path/to/new/test/repo"}
)

func TestInitialState(t *testing.T) {
	t.Log("Initial state checks")
	if len(GitmarkConfig.Repositories) != 0 {
		t.Error("Initial Repositories length should be 0")
	}

	if len(GitmarkConfig.ScanIgnores) != 0 {
		t.Error("Initial ScanIgnores length should be 0")
	}

	// check search pathes
	if ConfigFileSearchPathes[0] != ".gitmarkrc.json" || ConfigFileSearchPathes[1] != "~/.gitmarkrc.json" {
		t.Error("ConfigFileSearchPathes check failed")
	}
}

func TestLoadConfigFile(t *testing.T) {
	t.Log("Should be able to load the local config file")

	// prepend the config file search path with a test one
	ConfigFileSearchPathes = append([]string{testConfigFilePath}, ConfigFileSearchPathes...)
	if ConfigFileSearchPathes[0] != testConfigFilePath {
		t.Error("Failed to add the test Gitmarkrc to search path")
	}

	// load the config file
	err := ReadGitmarkConfigFromFile()
	if err != nil {
		t.Error("Failed to load the config file")
	}
}

func TestAfterLoadState(t *testing.T) {
	t.Log("After load state checks")

	if GitmarkConfig.ConfigFilePath != testConfigFilePath {
		t.Error("The test config file should be the loaded one")
	}

	// now check it's loaded
	if len(GitmarkConfig.Repositories) != 2 {
		t.Error("Repositories length should be 2")
	}

	if len(GitmarkConfig.ScanIgnores) != 2 {
		t.Error("ScanIgnores length should be 2")
	}

	secondRepo := GitmarkConfig.Repositories[1]
	if secondRepo.Title != "test/repo2" || secondRepo.Path != "/path/to/test2" {
		t.Error("Repository value checking failed")
	}
}

// --- the functions below will test a brand new loaded/mocked Config, not the GitmarkConig

func CreateTestConfig() Config {
	testConfigJsonContent := `
		{
		"repositories": [
			{
				"title": "test/repo1",
				"path": "/path/to/test1"
			},
			{
				"title": "test/repo2",
				"path": "/path/to/test2"
			}
		],
		"scanignores": [
			"/path/to/*/ignore1",
			"/path/to/*/ignore2"
		]
	}`
	testConfig, err := readConfigFromReader(strings.NewReader(testConfigJsonContent))
	if err != nil {
		panic("Could not create the test config")
	}
	return testConfig
}

func TestShouldNotFailForEmptyJSONConfigFile(t *testing.T) {
	t.Log("Should not fail for 'empty' config json")

	conf, err := readConfigFromReader(strings.NewReader("{}"))
	if err != nil {
		t.Error("Failed to create the test Config from {}")
	}
	if len(conf.Repositories) != 0 {
		t.Error("Repositories count should be 0")
	}

	// BUT a truly empty content is not a valid JSON!
	conf, err = readConfigFromReader(strings.NewReader(""))
	if err == nil {
		t.Error("Should return error if the content is not a valid json - an empty content is not a valid JSON!")
	}
}

func TestGetRepositoryPaths(t *testing.T) {
	t.Log("GetRepositoryPaths should return a slice of the repo pathes")
	testConfig := CreateTestConfig()

	repoPathes := testConfig.GetRepositoryPaths()
	if repoPathes[0] != repoPath1 || repoPathes[1] != repoPath2 {
		t.Error("Repository pathes check failed")
	}
}

func TestIsRepositoryPathStored(t *testing.T) {
	t.Log("IsRepositoryPathStored should return true for stored pathes and false for not stored ones")
	testConfig := CreateTestConfig()

	if !testConfig.IsRepositoryPathStored(repoPath1) || !testConfig.IsRepositoryPathStored(repoPath2) {
		t.Error("Repo path not found - should be")
	}

	if testConfig.IsRepositoryPathStored("/this/path/should/not/be/stored") {
		t.Error("Repo path found - should NOT be")
	}
}

func TestGetRepositoryByTitle(t *testing.T) {
	t.Log("GetRepositoryByTitle should return the related Repository by Title, or an error if not found")
	testConfig := CreateTestConfig()

	repo, err := testConfig.GetRepositoryByTitle(repoTitle1)
	if err != nil || repo.Title != repoTitle1 {
		t.Error("Repo should be found")
	}

	repo, err = testConfig.GetRepositoryByTitle("this repo doesn't exist")
	if err == nil {
		t.Error("Repo should NOT be found - error should be returned")
	}
}

func TestAddRepository(t *testing.T) {
	t.Log("AddRepository should add the new repository")
	testConfig := CreateTestConfig()

	if len(testConfig.Repositories) != 2 {
		t.Error("Repo count should be 2")
	}

	testConfig.AddRepository(newTestRepository)

	if len(testConfig.Repositories) != 3 {
		t.Error("Repo count should be 3")
	}

	if !testConfig.IsRepositoryPathStored(newTestRepository.Path) {
		t.Error("Repo path not found!")
	}
}

func TestWriteConfigToFile(t *testing.T) {
	t.Log("If config file path is not defined it should fail to write the config to file")
	testConfig := CreateTestConfig()

	err := writeConfigToFile(testConfig)
	if err == nil {
		t.Error("An error should be returned")
	}
}

func TestGenerateFormattedJSON(t *testing.T) {
	t.Log("GenerateFormattedJSON")
	testConfigJsonString := `{
	"repositories": [
		{
			"title": "test/repo1",
			"path": "/path/to/test1"
		},
		{
			"title": "test/repo2",
			"path": "/path/to/test2"
		}
	],
	"scanignores": [
		"/path/to/*/ignore1",
		"/path/to/*/ignore2"
	],
	"preferences": {
		"open-command": ""
	}
}`
	testConfig, err := readConfigFromReader(strings.NewReader(testConfigJsonString))
	if err != nil {
		t.Error("Failed to create the test Config")
	}

	jsonBytes, err := testConfig.GenerateFormattedJSON()
	if err != nil {
		t.Error("Failed to generate JSON:", err)
	}
	jsonString := fmt.Sprintf("%s", jsonBytes)
	if jsonString != testConfigJsonString {
		t.Log("Expected: ", testConfigJsonString)
		t.Log("Given: ", jsonString)
		t.Error("Generated JSON doesn't match")
	}
}
