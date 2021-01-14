package main

import (
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
	set, _ := ipset.New("test", ipset.HashIp, ipset.Exist(true), ipset.Timeout(time.Hour))
	// output: test
	log.Println(set.Name())

	_ = set.Flush()

	_ = set.Add("1.1.1.1", ipset.Timeout(time.Hour))

	ok, _ := set.Test("1.1.1.1")
	// output: true
	log.Println(ok)

	ok, _ = set.Test("1.1.1.2")
	// output: false
	log.Println(ok)

	info, _ := set.List()
	// output: &{test hash:ip 4 family inet hashsize 1024 maxelem 65536 timeout 3600 216 0 [1.1.1.1 timeout 3599]}
	log.Println(info)

	_ = set.Del("1.1.1.1")

	_ = set.Destroy()
}
