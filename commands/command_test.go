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

func TestAvailableCommands(t *testing.T) {
	t.Log("should contain all the available commands")

	var checkIsCommandAvailable = func(commandName string) {
		for _, anAvailableCmd := range AvailableCommands {
			if anAvailableCmd.Name == commandName {
				return
			}
		}
		t.Error(commandName, "command missing")
	}

	checkIsCommandAvailable("list")
	checkIsCommandAvailable("check")
	checkIsCommandAvailable("scan")
	checkIsCommandAvailable("open")
	checkIsCommandAvailable("add")
}
