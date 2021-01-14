package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"time"

	"github.com/gonetx/ipset"
)

func init() {
	if err := ipset.Check(); err != nil {
		panic(err)
	}
}

func main() {
	// create test set even it's exist
	set, _ := ipset.New("test", ipset.HashIp,
		ipset.Exist(true), ipset.Timeout(time.Hour*3),
		ipset.Family(ipset.Inet), ipset.HashSize(1024),
		ipset.MaxElem(100000),
	)

	// Saved content:
	_ = ioutil.WriteFile("saved",
		[]byte("add test 1.1.1.1 timeout 3600\nadd test 1.1.1.2 timeout 3600\n"),
		0600)
	_ = set.RestoreFromFile("saved", true)

	data := &bytes.Buffer{}
	data.WriteString("add test 1.1.1.3 timeout 3600 -exist\n")
	data.WriteString("add test 1.1.1.4 timeout 3600 -exist\n")

	_ = set.Restore(data)

	info, _ := set.List()
	// output: &{test hash:ip 4 family inet hashsize 1024 maxelem 100000 timeout 10800 504 0 [1.1.1.3 timeout 3599 1.1.1.2 timeout 3599 1.1.1.1 timeout 3599 1.1.1.4 timeout 3599]}
	log.Println(info)

	_ = set.SaveToFile("saved")
	// cat saved:
	//create test hash:ip family inet hashsize 1024 maxelem 100000 timeout 10800
	//add test 1.1.1.2 timeout 3599
	//add test 1.1.1.1 timeout 3599
	//add test 1.1.1.3 timeout 3599
	//add test 1.1.1.4 timeout 3599
}
