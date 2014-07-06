package commands

import (
	"reflect"
	"testing"
)

func TestCmdCheckExistence(t *testing.T) {
	t.Log("cmdCheck constant should be accessible")

	if cmtTypeStr := reflect.TypeOf(cmdCheck).String(); cmtTypeStr != "*commands.Command" {
		t.Error("cmdCheck should be accessible and should be a *commands.Command - but it's a:", cmtTypeStr)
	}
}

func Test_runCheck(t *testing.T) {
	t.Log("should run")

	err := runCheck(nil, nil)
	if err != nil {
		t.Error("should run without errors")
	}
}
