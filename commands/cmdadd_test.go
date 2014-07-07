package commands

import (
	"reflect"
	"testing"
)

func TestCmdAddExistence(t *testing.T) {
	t.Log("cmdAdd constant should be accessible")

	if cmtTypeStr := reflect.TypeOf(cmdAdd).String(); cmtTypeStr != "*commands.Command" {
		t.Error("cmdAdd should be accessible and should be a *commands.Command - but it's a:", cmtTypeStr)
	}
}

func Test_generateTitleFromRepositoryUrl(t *testing.T) {
	t.Log("should return the last component of the url, without any extension (.git, .ext, ...)")

	repoTitle, err := generateTitleFromRepositoryUrl("")
	if err == nil {
		t.Error("Should return an error for empty url")
	}
	repoTitle, err = generateTitleFromRepositoryUrl(" ")
	if err == nil {
		t.Error("Should return an error for a whitespace url")
	}

	repoTitle, err = generateTitleFromRepositoryUrl("a")
	if err != nil {
		t.Error("Should NOT return an error for a non empty url")
	}
	if repoTitle != "a" {
		t.Error("Not the expected title. Got:", repoTitle)
	}

	repoTitle, err = generateTitleFromRepositoryUrl("a/b.git")
	if err != nil {
		t.Error("Should NOT return an error for a non empty url")
	}
	if repoTitle != "b" {
		t.Error("Not the expected title. Got:", repoTitle)
	}

	repoTitle, err = generateTitleFromRepositoryUrl("https://github.com/viktorbenei/gitmark.git")
	if err != nil {
		t.Error("Should NOT return an error for a non empty url")
	}
	if repoTitle != "gitmark" {
		t.Error("Not the expected title. Got:", repoTitle)
	}

	repoTitle, err = generateTitleFromRepositoryUrl("git@github.com:viktorbenei/gitmark.git")
	if err != nil {
		t.Error("Should NOT return an error for a non empty url")
	}
	if repoTitle != "gitmark" {
		t.Error("Not the expected title. Got:", repoTitle)
	}
}

func Test_generateLocalPathFromUserInputPathAndRepositoryUrl(t *testing.T) {
	t.Log("should return the same path if it doesn't end with a slash, or append the repo url's last component if it ends with /")

	testRepoUrl := "git@github.com:viktorbenei/gitmark.git"

	genLocalPath, err := generateLocalPathFromUserInputPathAndRepositoryUrl("/some/path", testRepoUrl)
	if err != nil {
		t.Error("Should not return an error for a valid path and URL")
	}
	if genLocalPath != "/some/path" {
		t.Error("Path without a trailing / should generate the same. Got:", genLocalPath)
	}

	genLocalPath, err = generateLocalPathFromUserInputPathAndRepositoryUrl("/some/path/", testRepoUrl)
	if err != nil {
		t.Error("Should not return an error for a valid path and URL")
	}
	if genLocalPath != "/some/path/gitmark" {
		t.Error("Path with a trailing / should append the repo URL's last component. Got:", genLocalPath)
	}
}

// func Test_runAdd(t *testing.T) {
// 	t.Log("should run")

// 	err := runAdd(nil, nil)
// 	if err != nil {
// 		t.Error("should run without errors")
// 	}
// }
