package commands

import (
	"reflect"
	"testing"
)

func TestCmdGotoExistence(t *testing.T) {
	t.Log("cmdGoto constant should be accessible")

	if cmtTypeStr := reflect.TypeOf(cmdGoto).String(); cmtTypeStr != "*commands.Command" {
		t.Error("cmdGoto should be accessible and should be a *commands.Command - but it's a:", cmtTypeStr)
	}
}

func Test_runGoto(t *testing.T) {
	t.Log("should return an error if the given repository title is not found")

	err := runGoto(nil, []string{"non-existing-repo-title"})
	if err == nil {
		t.Error("Should return an error for a repo title which is not stored/found")
	}
}

func Test_runGoto_shouldAcceptOnlyOneArgument(t *testing.T) {
	t.Log("should accept only one argument, a bookmarked repository's title. Should return an error in any other case.")

	err := runGoto(nil, []string{})
	if err == nil {
		t.Error("Should return an error for 0 arguments")
	}

	err = runGoto(nil, []string{"repo-a", "repo-b"})
	if err == nil {
		t.Error("Should return an error for more than 1 arguments")
	}
}
