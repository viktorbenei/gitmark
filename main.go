package main

import (
	"./commands"
	"./config"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	filePath   = flag.String("filepath", "", "[REQUIRED] input file's path")
	prefixThis = flag.String("prefix", "", "prefix with this")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s command [FLAGS]", os.Args[0])
	flag.PrintDefaults()
	fmt.Println("Available commands:")
	for _, cmd := range commands.Commands {
		fmt.Println(" *", cmd.Name, ":", cmd.UsageString())
	}
}

func main() {
	err := config.ReadConfigFromFile()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("conf: ", config.GitmarkConfig)

	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	fmt.Println(args)
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	for _, cmd := range commands.Commands {
		if cmd.Name == args[0] {
			cmd.Flag.Usage = func() { cmd.UsageExit() }
			cmd.Flag.Parse(args[1:])
			cmd.Run(cmd, cmd.Flag.Args())
			return
		}
	}
}
