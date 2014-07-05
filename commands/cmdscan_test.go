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
