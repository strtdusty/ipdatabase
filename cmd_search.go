package main

import (
	"context"
	"flag"
	"net"
	"strings"

	"github.com/google/subcommands"
)

type cmdSearch struct {
}

func newCmdSearch() *cmdSearch {
	cmd := &cmdSearch{}
	return cmd
}

func (*cmdSearch) Name() string {
	return "search"
}

func (*cmdSearch) Synopsis() string { return "Search all entries" }
func (*cmdSearch) Usage() string {
	return `search 'searchstring':
  Search entries and print results to stdout.
`
}

func (p *cmdSearch) SetFlags(f *flag.FlagSet) {
}

//Search Search the database entry
func (d DatabaseEntry) Search(searchString string) {
	ip, net, e := net.ParseCIDR(searchString)
	if e == nil && net.IP.Equal(d.IPNet().IP) {
		d.Print()
	} else {
		if ip != nil && d.IPNet().Contains(ip) {
			d.Print()
		} else {
			for j := 0; j < len(d.Labels); j++ {
				if strings.Contains(d.Labels[j], searchString) {
					d.Print()
					break
				}
			}
		}
	}

}
func (p *cmdSearch) Execute(_ context.Context, f *flag.FlagSet,
	_ ...interface{}) subcommands.ExitStatus {
	database := Load()
	searchString := f.Arg(0)
	if strings.Count(searchString, ".") == 3 && !strings.Contains(searchString, "/32") {
		searchString = searchString + "/32"
	}

	for i := 0; i < len(database); i++ {
		database[i].Search(searchString)
	}
	return 0
}
