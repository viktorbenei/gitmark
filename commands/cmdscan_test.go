package commands

import (
	"reflect"
	"testing"
)

func TestCmdScanExistence(t *testing.T) {
	t.Log("cmdScan constant should be accessible")

	if cmtTypeStr := reflect.TypeOf(cmdScan).String(); cmtTypeStr != "*commands.Command" {
		t.Error("cmdScan should be accessible and should be a *commands.Command - but it's a:", cmtTypeStr)
	}
}

func Test_generateRepositoryTitleForRepositoryPath(t *testing.T) {
	t.Log("generateRepositoryTitleForRepositoryPath should generate a title by the path's directory name")

	expectedTitle := "test-repo"
	if generateRepositoryTitleForRepositoryPath("path/to/test-repo") != expectedTitle {
		t.Error("Title should be the repo directory name - even if it doesn't end with /")
	}

	if generateRepositoryTitleForRepositoryPath("path/to/test-repo/") != expectedTitle {
		t.Error("Title should be the repo directory name - even if it ends with /")
	}
}
