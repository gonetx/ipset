package ipset

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

// compiler assert
var _ IPSet = (*set)(nil)

type set struct {
	name    string
	setType SetType
}

// Info holds ipset list contents
type Info struct {
	Name string
	SetType
	Revision     int
	Header       string
	SizeInMemory int
	References   int
	Entries      []string
}

func (s set) List(options ...Option) (*Info, error) {
	c := getCmd(_list, s.name, s.setType)
	defer putCmd(c)
	if err := c.exec(options...); err != nil {
		return nil, err
	}

	info, err := parseInfo(c.out)
	if err != nil {
		return nil, err
	}
	info.Name = s.name
	info.SetType = s.setType
	return info, err
}

func parseInfo(out []byte) (info *Info, err error) {
	info = &Info{}
	s := bufio.NewScanner(bytes.NewReader(out))

	for s.Scan() {
		t := s.Text()
		switch {
		case strings.HasPrefix(t, "Rev"):
			if info.Revision, err = getNumber(t); err != nil {
				return nil, err
			}
		case strings.HasPrefix(t, "H"):
			info.Header = t[8:]
		case strings.HasPrefix(t, "S"):
			if info.SizeInMemory, err = getNumber(t); err != nil {
				return nil, err
			}
		case strings.HasPrefix(t, "Ref"):
			if info.References, err = getNumber(t); err != nil {
				return nil, err
			}
		case strings.HasPrefix(t, "M"):
			goto Entries
		}
	}
Entries:
	for s.Scan() {
		info.Entries = append(info.Entries, s.Text())
	}

	return
}

func getNumber(t string) (n int, err error) {
	if i := strings.LastIndexByte(t, ' '); i != -1 {
		return strconv.Atoi(t[i+1:])
	}
	return
}

func (s set) ListToFile(filename string, options ...Option) error {
	return s.doToFile(_list, filename, options...)
}

func (s set) Name() string {
	return s.name
}

func (s set) Rename(newName string) error {
	return s.do(_rename, newName)
}

func (s set) Add(entry string, options ...Option) error {
	return s.do(_add, entry, options...)
}

func (s set) Del(entry string, options ...Option) error {
	return s.do(_del, entry, options...)
}

var notFlag = []byte("NOT")

func (s set) Test(entry string) (bool, error) {
	out, err := execCommand(ipsetPath, _test, s.name, entry).
		CombinedOutput()

	if err != nil {
		if bytes.Contains(out, notFlag) {
			return false, nil
		}
		return false, fmt.Errorf("ipset: can't test %s %s: %s", s.name, entry, out)
	}

	return true, nil
}

func (s set) Flush() error {
	return flush(s.name)
}

func (s set) Destroy() error {
	return destroy(s.name)
}

func (s set) do(action, entry string, options ...Option) error {
	c := getCmd(action, s.name, s.setType, entry)
	defer putCmd(c)

	if err := c.exec(options...); err != nil {
		return err
	}
	return nil
}

func (s set) Save(options ...Option) (io.Reader, error) {
	c := getCmd(_save, s.name, s.setType)
	defer putCmd(c)
	if err := c.exec(options...); err != nil {
		return nil, err
	}

	return bytes.NewReader(c.out), nil
}

func (s set) SaveToFile(filename string, options ...Option) error {
	return s.doToFile(_save, filename, options...)
}

func (s set) doToFile(action, filename string, options ...Option) error {
	c := getCmd(action, s.name, s.setType)
	defer putCmd(c)
	if err := c.exec(options...); err != nil {
		return err
	}

	return ioutil.WriteFile(filename, c.out, 0600)
}

var maxRestoreSize = 1 << 16

func (s set) Restore(r io.Reader, exist ...bool) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("ipset: can't restore to %s(%s): %s", s.name, s.setType, err)
		}
	}()

	var (
		br = acquireReader(r)
		b  = &bytes.Buffer{}
		bb []byte
	)
	defer releaseReader(br)

	for {
		bb, err = br.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}
		if b.Len()+len(bb) > maxRestoreSize {
			if err = s.restore(b.Bytes(), exist...); err != nil {
				return
			}
			b.Reset()
		}
		if _, err = b.Write(bb); err != nil {
			return
		}
	}
	return s.restore(b.Bytes(), exist...)
}

// restore data to ipset and length of b should
// be less than 64K because of the size limit of
// pipe.
func (s set) restore(b []byte, exist ...bool) (err error) {
	args := []string{_restore}
	if len(exist) > 0 && exist[0] {
		args = append(args, _exist)
	}
	c := execCommand(ipsetPath, args...)

	var pipe io.WriteCloser
	pipe, err = c.StdinPipe()
	if err != nil {
		return err
	}

	if _, err = pipe.Write(b); err != nil {
		return pipe.Close()
	}

	if err = pipe.Close(); err != nil {
		return
	}

	var out []byte
	if out, err = c.CombinedOutput(); err != nil {
		return fmt.Errorf("%s", out)
	}

	return
}

func (s set) RestoreFromFile(filename string, exist ...bool) (err error) {
	var f *os.File
	f, err = os.Open(filepath.Clean(filename))
	if err != nil {
		return
	}
	defer func() {
		if e := f.Close(); e != nil {
			err = e
		}
	}()
	return s.Restore(f, exist...)
}

var readerPool sync.Pool

func acquireReader(r io.Reader) *bufio.Reader {
	v := readerPool.Get()
	if v == nil {
		return bufio.NewReader(r)
	}
	br := v.(*bufio.Reader)
	br.Reset(r)
	return br
}

func releaseReader(br *bufio.Reader) {
	readerPool.Put(br)
}
