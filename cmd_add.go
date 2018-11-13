package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/google/subcommands"
)

type cmdAdd struct {
	label   string
	address string
}

//NewCmdAdd shortcut to create new command
func newCmdAdd() *cmdAdd {
	cmd := &cmdAdd{}
	return cmd
}

func (*cmdAdd) Name() string {
	return "add"
}

func (*cmdAdd) Synopsis() string { return "Add a new entry" }
func (*cmdAdd) Usage() string {
	return `add -address someaddress -label somelabel:
  Add a new entry
`
}

func (p *cmdAdd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.label, "label", "", "label to add to the address")
	f.StringVar(&p.address, "address", "", "address to add")
}
func (p *cmdAdd) Execute(_ context.Context, f *flag.FlagSet,
	_ ...interface{}) subcommands.ExitStatus {
	database := Load()
	address := p.address
	if !strings.Contains(address, "/") {
		address = address + "/32"
	}
	_, _, e := net.ParseCIDR(address)
	if e != nil {
		fmt.Printf("%s", e.Error())
		os.Exit(1)
	}
	d := DatabaseEntry{
		IP: address,
	}
	d.Labels = make([]string, 0)
	d.Labels = append(d.Labels, p.label)
	database = append(database, d)
	Save(database)
	fmt.Printf("Added and updated")
	return 0
}
