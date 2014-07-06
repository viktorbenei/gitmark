package commands

import (
	"reflect"
	"testing"
)

func TestCmdOpenExistence(t *testing.T) {
	t.Log("cmdOpen constant should be accessible")

	if cmtTypeStr := reflect.TypeOf(cmdOpen).String(); cmtTypeStr != "*commands.Command" {
		t.Error("cmdOpen should be accessible and should be a *commands.Command - but it's a:", cmtTypeStr)
	}
}

func Test_runOpen(t *testing.T) {
	t.Log("should return an error if the given repository title is not found")

	err := runOpen(nil, []string{"non-existing-repo-title"})
	if err == nil {
		t.Error("Should return an error for a repo title which is not stored/found")
	}
}

func Test_runOpen_shouldAcceptOnlyOneArgument(t *testing.T) {
	t.Log("should accept only one argument, a bookmarked repository's title. Should return an error in any other case.")

	err := runOpen(nil, []string{})
	if err == nil {
		t.Error("Should return an error for 0 arguments")
	}

	err = runOpen(nil, []string{"repo-a", "repo-b"})
	if err == nil {
		t.Error("Should return an error for more than 1 arguments")
	}
}
