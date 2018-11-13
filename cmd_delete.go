package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"sort"
	"strings"

	"github.com/google/subcommands"
)

type cmdDelete struct {
	address string
}

func (*cmdDelete) Name() string {
	return "delete"
}
func (*cmdDelete) Synopsis() string { return "Delete an entry" }
func (*cmdDelete) Usage() string {
	return `delete addresstodelete:
  Deletes an entry from the list.
`
}

func (p *cmdDelete) SetFlags(f *flag.FlagSet) {
}

func newCmdDelete() *cmdDelete {
	cmd := &cmdDelete{}
	return cmd
}

func (p *cmdDelete) Execute(_ context.Context, f *flag.FlagSet,
	_ ...interface{}) subcommands.ExitStatus {
	database := Load()

	if !strings.Contains(p.address, "/32") {
		p.address = p.address + "/32"
	}
	_, _, e := net.ParseCIDR(p.address)
	if e != nil {
		fmt.Printf("%s", e.Error())
		return 1
	}
	idx := sort.Search(len(database), func(i int) bool {
		return p.address == database[i].IP
	})
	if idx > -1 && idx < len(database) {
		database = append(database[:idx], database[idx+1:]...)
		Save(database)
		fmt.Printf("Deleted and updated")
	} else {
		fmt.Print("Address not found")
	}
	return 0
}
