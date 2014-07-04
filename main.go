package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/viktorbenei/gitmark/commands"
	"github.com/viktorbenei/gitmark/config"
)

var (
	// filePath   = flag.String("filepath", "", "[REQUIRED] input file's path")
	// prefixThis = flag.String("prefix", "", "prefix with this")
	isVerbose = flag.Bool("verbose", false, "is verbose?")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s command [FLAGS]", os.Args[0])
	flag.PrintDefaults()
	fmt.Println("Available commands:")
	for _, cmd := range commands.AvailableCommands {
		fmt.Println(" *", cmd.Name)
		fmt.Println("    ", cmd.UsageString())
	}
}

func main() {
	err := config.ReadConfigFromFile()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("conf: ", config.GitmarkConfig)

	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	// fmt.Println(args)
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	for _, cmd := range commands.AvailableCommands {
		if cmd.Name == args[0] {
			cmd.Flag.Usage = func() { cmd.UsageExit() }
			cmd.Flag.Parse(args[1:])
			err := cmd.Run(cmd, cmd.Flag.Args())
			if err != nil {
				log.Fatalln(err)
			}
			return
		}
	}
}
