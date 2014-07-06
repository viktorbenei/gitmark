package commands

import (
	"flag"
	"fmt"
	"os"
)

type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(cmd *Command, args []string) error

	// Usage is the one-line usage message.
	// The first word in the line is taken to be the command name.
	Usage string

	// Short is the short description shown in the 'godep help' output.
	Short string

	Name string

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet
}

func (c *Command) UsageString() string {
	return fmt.Sprintf("Usage: gitmark %s", c.Usage)
}

func (c *Command) UsageExit() {
	fmt.Fprintf(os.Stderr, "Usage: gitmark %s\n\n", c.Usage)
	// fmt.Fprintf(os.Stderr, "Run 'godep help %s' for help.\n", c.Name)
	os.Exit(2)
}

var AvailableCommands = []*Command{
	cmdList,
	cmdCheck,
	cmdScan,
	cmdGoto,
}
