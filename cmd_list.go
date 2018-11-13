package main

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type cmdList struct {
}

func newCmdList() *cmdList {
	cmd := &cmdList{}
	return cmd
}

func (*cmdList) Name() string {
	return "list"
}
func (*cmdList) Synopsis() string { return "List all entries" }
func (*cmdList) Usage() string {
	return `list:
  Print entries to stdout.
`
}

func (p *cmdList) SetFlags(f *flag.FlagSet) {
}

func (p *cmdList) Execute(_ context.Context, f *flag.FlagSet,
	_ ...interface{}) subcommands.ExitStatus {
	database := Load()

	for i := 0; i < len(database); i++ {
		database[i].Print()
	}
	return 0
}
