package ipset

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Set_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		setupCmd()
		defer teardownCmd()
		s := getSet()

		info, err := s.List()
		require.Nil(t, err)
		require.NotNil(t, info)
		assert.Equal(t, s.name, info.Name)
		assert.Equal(t, s.setType, info.SetType)
		assert.Equal(t, 4, info.Revision)
		assert.Equal(t, "family inet hashsize 1024 maxelem 65536", info.Header)
		assert.Equal(t, 0, info.References)
		assert.Equal(t, "1.1.1.1", info.Entries[0])
	})

	t.Run("error", func(t *testing.T) {
		setupCmd(flag)
		defer teardownCmd()
		s := getSet()

		_, err := s.List()
		require.Error(t, err)

		assert.Equal(t,
			fmt.Sprintf("ipset: can't %s %s: fake error", _list, s.name),
			err.Error())

	})
}

func Test_Set_ListToFile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		setupCmd()
		defer teardownCmd()
		s := getSet()

		filename := "list.test"
		defer removeFile(t, filename)

		err := s.ListToFile(filename, Resolve(true))
		require.Nil(t, err)

		b, err := ioutil.ReadFile(filename)
		require.Nil(t, err)
		assert.Contains(t, string(b), "one.one.one.one")
	})

	t.Run("error", func(t *testing.T) {
		setupCmd(flag)
		defer teardownCmd()
		s := getSet()

		err := s.ListToFile("")
		require.Error(t, err)

		assert.Equal(t,
			fmt.Sprintf("ipset: can't %s %s: fake error", _list, s.name),
			err.Error())

	})
}

func Test_Set_Name(t *testing.T) {
	s := getSet()
	assert.Equal(t, s.name, s.Name())
}

func Test_Set_Rename(t *testing.T) {
	newName := "newName"

	t.Run("success", func(t *testing.T) {
		setupCmd()
		defer teardownCmd()
		s := getSet()

		err := s.Rename(newName)
		require.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		setupCmd(flag)
		defer teardownCmd()
		s := getSet()

		err := s.Rename(newName)
		require.Error(t, err)

		assert.Equal(t,
			fmt.Sprintf("ipset: can't %s %s %s: fake error", _rename, s.name, newName),
			err.Error())

	})
}

func Test_Set_Add(t *testing.T) {
	ip := "1.1.1.1"
	t.Run("success", func(t *testing.T) {
		setupCmd()
		defer teardownCmd()
		s := getSet()

		err := s.Add(ip)
		require.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		setupCmd(flag)
		defer teardownCmd()
		s := getSet()

		err := s.Add(ip)
		require.Error(t, err)

		assert.Equal(t,
			fmt.Sprintf("ipset: can't %s %s %s: fake error", _add, s.name, ip),
			err.Error())

	})
}

func Test_Set_Del(t *testing.T) {
	ip := "1.1.1.1"
	t.Run("success", func(t *testing.T) {
		setupCmd()
		defer teardownCmd()
		s := getSet()

		err := s.Del(ip)
		require.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		setupCmd(flag)
		defer teardownCmd()
		s := getSet()

		err := s.Del(ip)
		require.Error(t, err)

		assert.Equal(t,
			fmt.Sprintf("ipset: can't %s %s %s: fake error", _del, s.name, ip),
			err.Error())

	})
}

func Test_Set_Test(t *testing.T) {
	ip := "1.1.1.1"
	t.Run("success", func(t *testing.T) {
		setupCmd()
		defer teardownCmd()
		s := getSet()

		ok, err := s.Test(ip)
		require.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("not exist", func(t *testing.T) {
		setupCmd()
		defer teardownCmd()
		s := getSet()

		ok, err := s.Test(testNotExistIp)
		assert.Nil(t, err)
		assert.False(t, ok)
	})

	t.Run("error", func(t *testing.T) {
		setupCmd(flag)
		defer teardownCmd()
		s := getSet()

		ok, err := s.Test(ip)
		require.Error(t, err)
		assert.False(t, ok)

		assert.Equal(t,
			fmt.Sprintf("ipset: can't %s %s %s: fake error", _test, s.name, ip),
			err.Error())
	})
}

func Test_Set_Flush(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		setupCmd()
		defer teardownCmd()
		s := getSet()

		assert.Nil(t, s.Flush())
	})

	t.Run("error", func(t *testing.T) {
		setupCmd(flag)
		defer teardownCmd()
		s := getSet()

		err := s.Flush()
		require.Error(t, err)
		assert.Equal(t,
			fmt.Sprintf("ipset: can't flush set %s: fake error", s.name),
			err.Error())
	})
}

func Test_Set_Destroy(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		setupCmd()
		defer teardownCmd()
		s := getSet()

		err := s.Destroy()
		require.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		setupCmd(flag)
		defer teardownCmd()
		s := getSet()

		err := s.Destroy()
		require.Error(t, err)

		assert.Equal(t,
			fmt.Sprintf("ipset: can't %s set %s: fake error", _destroy, s.name),
			err.Error())

	})
}

func Test_Set_Save(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		setupCmd()
		defer teardownCmd()
		s := getSet()

		r, err := s.Save()
		require.Nil(t, err)
		b := bytes.Buffer{}
		_, err = b.ReadFrom(r)
		require.Nil(t, err)
		assert.Equal(t, saveInfo, b.String())
	})

	t.Run("error", func(t *testing.T) {
		setupCmd(flag)
		defer teardownCmd()
		s := getSet()

		_, err := s.List()
		require.Error(t, err)

		assert.Equal(t,
			fmt.Sprintf("ipset: can't %s %s: fake error", _list, s.name),
			err.Error())

	})
}

func Test_Set_SaveToFile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		setupCmd()
		defer teardownCmd()
		s := getSet()

		filename := "save.test"
		defer removeFile(t, filename)

		err := s.SaveToFile(filename, Resolve(true))
		require.Nil(t, err)

		b, err := ioutil.ReadFile(filename)
		require.Nil(t, err)
		assert.Contains(t, string(b), "one.one.one.one")
	})

	t.Run("error", func(t *testing.T) {
		setupCmd(flag)
		defer teardownCmd()
		s := getSet()

		err := s.SaveToFile("")
		require.Error(t, err)

		assert.Equal(t,
			fmt.Sprintf("ipset: can't %s %s: fake error", _save, s.name),
			err.Error())

	})
}

func Test_Set_Restore(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		maxRestoreSize = 10
		setupCmd()
		defer teardownCmd()
		s := getSet()

		err := s.Restore(bytes.NewReader([]byte("1.1.1.1\n2.2.2.2\n")))
		require.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		setupCmd(flag)
		defer teardownCmd()
		s := getSet()

		err := s.Restore(bytes.NewReader([]byte("1.1.1.1\n")))
		require.Error(t, err)

		assert.Equal(t,
			fmt.Sprintf("ipset: can't restore to %s(%s): fake error", s.name, s.setType),
			err.Error())

	})
}

func Test_Set_RestoreFromFile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		maxRestoreSize = 10
		setupCmd()
		defer teardownCmd()
		s := getSet()

		filename := "restore.test"
		require.NoError(t, ioutil.WriteFile(filename, []byte("1.1.1.1\n"), 0600))
		defer removeFile(t, filename)

		err := s.RestoreFromFile(filename, true)
		require.Nil(t, err)
	})

	t.Run("no file error", func(t *testing.T) {
		setupCmd(flag)
		defer teardownCmd()
		s := getSet()

		filename := "restore.test"
		err := s.RestoreFromFile(filename)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "open")

	})

	t.Run("cmd error", func(t *testing.T) {
		setupCmd(flag)
		defer teardownCmd()
		s := getSet()

		filename := "restore.test"
		require.NoError(t, ioutil.WriteFile(filename, []byte("1.1.1.1\n"), 0600))
		defer removeFile(t, filename)

		err := s.RestoreFromFile(filename)
		require.Error(t, err)

		assert.Equal(t,
			fmt.Sprintf("ipset: can't restore to %s(%s): fake error", s.name, s.setType),
			err.Error())

	})
}

func getSet(setType ...SetType) set {
	s := set{"test", HashIp}
	if len(setType) > 0 {
		s.setType = setType[0]
	}
	return s
}

func removeFile(t assert.TestingT, filename string) {
	assert.Nil(t, os.Remove(filename))
}
