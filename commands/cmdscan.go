package commands

import (
// "fmt"
// "github.com/viktorbenei/gitmark/config"
// "os/exec"
)

var cmdScan = &Command{
	Usage: "scan",
	Short: "scan for local repositories which are not yet bookmarked",
	Run:   nil,
	Name:  "scan",
}
