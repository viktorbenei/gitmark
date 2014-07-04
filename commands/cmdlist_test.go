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
