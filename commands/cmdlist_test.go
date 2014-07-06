package commands

import (
	"reflect"
	"testing"
)

func TestExistence(t *testing.T) {
	t.Log("cmdList constant should be accessible")

	if cmtListTypeStr := reflect.TypeOf(cmdList).String(); cmtListTypeStr != "*commands.Command" {
		t.Error("cmdList should be accessible and should be a *commands.Command - but it's a:", cmtListTypeStr)
	}
}

func Test_runList(t *testing.T) {
	t.Log("should run")

	err := runList(nil, nil)
	if err != nil {
		t.Error("should run without errors")
	}
}
