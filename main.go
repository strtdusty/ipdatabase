package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
)

func main() {
	fmt.Println(os.Args)
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	subcommands.Register(newCmdList(), "")
	subcommands.Register(newCmdAdd(), "")
	subcommands.Register(newCmdSearch(), "")
	subcommands.Register(newCmdDelete(), "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
