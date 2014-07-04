package commands

import (
	"testing"
)

func TestUsageString(t *testing.T) {
	t.Log("should return 'Usage: gitmark ' + the command's .Usage, as a string")

	var aTestCmd = &Command{
		Usage: "a-test",
	}

	if usgString := aTestCmd.UsageString(); usgString != "Usage: gitmark a-test" {
		t.Error("returned:", usgString)
	}
}
