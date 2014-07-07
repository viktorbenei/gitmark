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

func Test_runAdd(t *testing.T) {
	t.Log("should run")

	err := runAdd(nil, nil)
	if err != nil {
		t.Error("should run without errors")
	}
}
