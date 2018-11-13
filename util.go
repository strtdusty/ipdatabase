package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"path/filepath"
	"sort"

	"github.com/kardianos/osext"
)

//Load the database from file
func Load() []DatabaseEntry {
	fmt.Println("Loading from ", path())
	database := make([]DatabaseEntry, 0)

	jsonFile, _ := ioutil.ReadFile(path())
	json.Unmarshal(jsonFile, &database)
	return database
}

//Save save the database to file
func Save(db []DatabaseEntry) {
	sort.Slice(db, func(i, j int) bool {
		a := db[i].IPNet()
		b := db[j].IPNet()
		return a.IP[0] < b.IP[0] ||
			a.IP[1] < b.IP[1] ||
			a.IP[2] < b.IP[2] ||
			a.IP[3] < b.IP[3] ||
			a.Mask[0] < b.Mask[0] ||
			a.Mask[1] < b.Mask[1] ||
			a.Mask[2] < b.Mask[2] ||
			a.Mask[3] < b.Mask[3]
	})

	json, _ := json.MarshalIndent(db, "", "    ")
	ioutil.WriteFile(path(), json, 0777)
}
func path() string {
	path, _ := osext.ExecutableFolder()
	path = filepath.Join(path, "ipdata.json")
	return path
}

//Print Print the database entry
func (d DatabaseEntry) Print() {
	fmt.Printf("%s ", d.IP)
	for i := 0; i < len(d.Labels); i++ {
		fmt.Printf("%s ", d.Labels[i])
	}
	fmt.Println("")
}

//DatabaseEntry IP Address database entry
type DatabaseEntry struct {
	IP     string
	Labels []string
}

//IPNet Convert the IP of the DatabaseEntry into an IPNet
func (d DatabaseEntry) IPNet() *net.IPNet {
	_, ipNet, _ := net.ParseCIDR(d.IP)
	return ipNet
}
