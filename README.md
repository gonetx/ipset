# gonetx/ipset
<p align="center">
  <a href="https://pkg.go.dev/github.com/gonetx/ipset?tab=doc">
    <img src="https://img.shields.io/badge/%F0%9F%93%9A%20godoc-pkg-00ACD7.svg?color=00ACD7&style=flat">
  </a>
  <a href="https://goreportcard.com/report/github.com/gonetx">
    <img src="https://img.shields.io/badge/%F0%9F%93%9D%20goreport-A%2B-75C46B">
  </a>
  <a href="https://gocover.io/github.com/gonetx/ipset">
    <img src="https://img.shields.io/badge/%F0%9F%94%8E%20gocover-97.8%25-75C46B.svg?style=flat">
  </a>
  <a href="https://github.com/gonetx/actions?query=workflow%3ASecurity">
    <img src="https://img.shields.io/github/workflow/status/gofiber/fiber/Security?label=%F0%9F%94%91%20gosec&style=flat&color=75C46B">
  </a>
  <a href="https://github.com/gonetx/actions?query=workflow%3ATest">
    <img src="https://img.shields.io/github/workflow/status/gofiber/fiber/Test?label=%F0%9F%A7%AA%20tests&style=flat&color=75C46B">
  </a>
  <a href="https://github.com/gonetx/tree/master/README_zh-CN.md">
      <img height="20px" src="https://img.shields.io/badge/CN-flag.svg?color=555555&style=flat&logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAxMjAwIDgwMCIgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiPg0KPHBhdGggZmlsbD0iI2RlMjkxMCIgZD0ibTAsMGgxMjAwdjgwMGgtMTIwMHoiLz4NCjxwYXRoIGZpbGw9IiNmZmRlMDAiIGQ9Im0tMTYuNTc5Niw5OS42MDA3bDIuMzY4Ni04LjEwMzItNi45NTMtNC43ODgzIDguNDM4Ni0uMjUxNCAyLjQwNTMtOC4wOTI0IDIuODQ2Nyw3Ljk0NzkgOC40Mzk2LS4yMTMxLTYuNjc5Miw1LjE2MzQgMi44MTA2LDcuOTYwNy02Ljk3NDctNC43NTY3LTYuNzAyNSw1LjEzMzF6IiB0cmFuc2Zvcm09Im1hdHJpeCg5LjkzMzUyIC4yNzc0NyAtLjI3NzQ3IDkuOTMzNTIgMzI0LjI5MjUgLTY5NS4yNDE1KSIvPg0KPHBhdGggZmlsbD0iI2ZmZGUwMCIgaWQ9InN0YXIiIGQ9Im0zNjUuODU1MiwzMzIuNjg5NWwyOC4zMDY4LDExLjM3NTcgMTkuNjcyMi0yMy4zMTcxLTIuMDcxNiwzMC40MzY3IDI4LjI1NDksMTEuNTA0LTI5LjU4NzIsNy40MzUyLTIuMjA5NywzMC40MjY5LTE2LjIxNDItMjUuODQxNS0yOS42MjA2LDcuMzAwOSAxOS41NjYyLTIzLjQwNjEtMTYuMDk2OC0yNS45MTQ4eiIvPg0KPGcgZmlsbD0iI2ZmZGUwMCI+DQo8cGF0aCBkPSJtNTE5LjA3NzksMTc5LjMxMjlsLTMwLjA1MzQtNS4yNDE4LTE0LjM5NDUsMjYuODk3Ni00LjMwMTctMzAuMjAyMy0zMC4wMjkzLTUuMzc4MSAyNy4zOTQ4LTEzLjQyNDItNC4xNjQ3LTMwLjIyMTUgMjEuMjMyNiwyMS45MDU3IDI3LjQ1NTQtMTMuMjk5OC0xNC4yNzIzLDI2Ljk2MjcgMjEuMTMzMSwyMi4wMDE3eiIvPg0KPHBhdGggZD0ibTQ1NS4yNTkyLDMxNS45Nzk1bDkuMzczNC0yOS4wMzE0LTI0LjYzMjUtMTcuOTk3OCAzMC41MDctLjA1NjYgOS41MDUtMjguOTg4NiA5LjQ4MSwyOC45OTY0IDMwLjUwNywuMDgxOC0yNC42NDc0LDE3Ljk3NzQgOS4zNDkzLDI5LjAzOTItMjQuNzE0LTE3Ljg4NTgtMjQuNzI4OCwxNy44NjUzeiIvPg0KPC9nPg0KPHVzZSB4bGluazpocmVmPSIjc3RhciIgdHJhbnNmb3JtPSJtYXRyaXgoLjk5ODYzIC4wNTIzNCAtLjA1MjM0IC45OTg2MyAxOS40MDAwNSAtMzAwLjUzNjgxKSIvPg0KPC9zdmc+DQo=">
    </a>
</p>

This package is a almost whole `Golang` wrapper to the `ipset` userspace utility. It allows `Golang` programs to easily manipulate `ipset`.

Visit [http://ipset.netfilter.org/ipset.man.html](http://ipset.netfilter.org/ipset.man.html) for more `ipset` command details. And `ipset` requires version `v6.0+`.

## Installation
Install `ipset` using the `go get` command:
```bash
go get -u github.com/gonetx/ipset
```

## Quickstart

```go
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
```

## Check
Before using this library, you should remember always to call `ipset.Check` first and handle the error. This method will check if `ipset` is in `OS PATH` and if its version is valid.

```go
func init() {
	// err will be ipset.ErrNotFound
	// or ipset.ErrVersionNotSupported
	// if check failed.
	if err := ipset.Check(); err != nil {
		panic(err)
	}
}
```

## New
Use `ipset.New` to create a set identified with setname and specified set type. If the `ipset.Exist` option is specified, `ipset` will ignore the error when the same set (setname and create parameters are identical) already exists.

```go
set, _ := ipset.New("test", ipset.HashIp, ipset.Exist(true), ipset.Netmask(24))
```

Every set type may has different create options, visit [SetType](https://pkg.go.dev/github.com/gonetx/ipset?tab=doc#SetType) and [Option](https://pkg.go.dev/github.com/gonetx/ipset?tab=doc#Option) for more details.

Once you created a set, you can use these methods:
```go
// IPSet is abstract of ipset
type IPSet interface {
	// List dumps header data and the entries for the set to an
	// *Info instance. The Resolve option can be used to force
	// action lookups(which may be slow).
	List(options ...Option) (*Info, error)

	// List dumps header data and the entries for the set to the
	// specific file. The Resolve option can be used to force
	// action lookups(which may be slow).
	ListToFile(filename string, options ...Option) error

	// Name returns the set's name
	Name() string

	// Rename the set's action and the new action must not exist.
	Rename(newName string) error

	// Add adds a given entry to the set. If the Exist option is
	// specified, ipset ignores the error if the entry already
	// added to the set.
	Add(entry string, options ...Option) error

	// Del deletes an entry from a set. If the Exist option is
	// specified and the entry is not in the set (maybe already
	// expired), then the command ignores the error.
	Del(entry string, options ...Option) error

	// Test tests whether an entry is in a set or not.
	Test(entry string) (bool, error)

	// Flush flushed all entries from the the set.
	Flush() error

	// Destroy removes the set from kernel.
	Destroy() error

	// Save dumps the set data to a io.Reader in a format that restore
	// can read.
	Save(options ...Option) (io.Reader, error)

	// SaveToFile dumps the set data to s specific file in a format
	// that restore can read.
	SaveToFile(filename string, options ...Option) error

	// Restore restores a saved session from io.Reader generated by
	// save. Set exist to true to ignore exist error.
	Restore(r io.Reader, exist ...bool) error

	// RestoreFromFile restores a saved session from a specific file
	// generated by save. Set exist to true to ignore exist error.
	RestoreFromFile(filename string, exist ...bool) error
}
```

## Swap
Use `ipset.Swap` to swap the content of two sets, or in another words, exchange the action of two sets. The referred sets must exist and compatible type of sets can be swapped only. 

```go
ipset.Swap("foo", "bar")
```
## Flush
Use `ipset.Flush` to flush all entries from the specified set or flush all sets if none is given.

```go
// Flush foo and bar set
ipset.Flush("foo", "bar")

// Flush all
ipset.Flush()
```

## Destroy
Use `ipset.Destroy` to removes the specified set or all the sets if none is given. If the set has got reference(s), nothing is done and no set destroyed.

```go
// Destroy foo and bar set
ipset.Destroy("foo", "bar")

// Destroy all
ipset.Destroy()
```
