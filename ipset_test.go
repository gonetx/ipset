package ipset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func Test_Check(t *testing.T) {
	t.Run("ipset path is ready", func(t *testing.T) {
		ipsetPath = "I'm ready"
		defer func() { ipsetPath = "" }()
		assert.Nil(t, Check())
	})

	t.Run("ipset path is not exist", func(t *testing.T) {
		setupLookPath("error")
		defer teardownLookPath()
		assert.Equal(t, ErrNotFound, Check())
	})

	t.Run("ipset version error", func(t *testing.T) {
		setupLookPath()
		defer teardownLookPath()
		setupCmd(flag)
		defer teardownCmd()

		assert.NotNil(t, Check())
	})

	t.Run("supported version", func(t *testing.T) {
		setupLookPath()
		defer teardownLookPath()
		setupCmd()
		defer teardownCmd()

		assert.Nil(t, Check())
	})

	t.Run("non supported version", func(t *testing.T) {
		setupLookPath("non-supported")
		defer teardownLookPath()
		setupCmd()
		defer teardownCmd()

		assert.Equal(t, ErrVersionNotSupported, Check())
	})
}

func Test_GetMajorVersion(t *testing.T) {
	t.Parallel()

	tt := []struct {
		out []byte
		v   int
	}{
		{[]byte("no version"), 0},
		{[]byte("ipset v5e.rsion"), 0},
		{[]byte("ipset v6.29, protocol version: 6"), 6},
		{[]byte("ipset v7.1, protocol version: 7"), 7},
		{[]byte("ipset v10.0, protocol version: 10"), 10},
		{[]byte("Warning: Kernel support protocol versions 6-6 while userspace supports protocol versions 6-7\nipset v7.1, protocol version: 7"), 7},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.v, getMajorVersion(tc.out))
	}
}

func Test_New(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		setupCmd()
		defer teardownCmd()

		s, err := New("test", HashIp)
		assert.Nil(t, err)
		assert.NotNil(t, s)
	})

	t.Run("error", func(t *testing.T) {
		setupCmd(flag)
		defer teardownCmd()

		_, err := New("test", HashIp)
		require.Error(t, err)
		assert.Equal(t,
			fmt.Sprintf("ipset: can't %s %s %s: fake error", _create, "test", HashIp),
			err.Error())
	})
}

func Test_Flush(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			setupCmd()
			defer teardownCmd()

			assert.Nil(t, Flush())
		})

		t.Run("error", func(t *testing.T) {
			setupCmd(flag)
			defer teardownCmd()

			err := Flush()
			require.Error(t, err)
			assert.Equal(t,
				"ipset: can't flush all set: fake error",
				err.Error())
		})
	})

	t.Run("multi", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			setupCmd()
			defer teardownCmd()

			assert.Nil(t, Flush("a", "b"))
		})

		t.Run("error", func(t *testing.T) {
			setupCmd(flag)
			defer teardownCmd()

			err := Flush("a", "b")
			require.Error(t, err)
			assert.Equal(t,
				"ipset: can't flush set a: fake error",
				err.Error())
		})
	})
}

func Test_Destroy(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			setupCmd()
			defer teardownCmd()

			assert.Nil(t, Destroy())
		})

		t.Run("error", func(t *testing.T) {
			setupCmd(flag)
			defer teardownCmd()

			err := Destroy()
			require.Error(t, err)
			assert.Equal(t,
				"ipset: can't destroy all set: fake error",
				err.Error())
		})
	})

	t.Run("multi", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			setupCmd()
			defer teardownCmd()

			assert.Nil(t, Destroy("a", "b"))
		})

		t.Run("error", func(t *testing.T) {
			setupCmd(flag)
			defer teardownCmd()

			err := Destroy("a", "b")
			require.Error(t, err)
			assert.Equal(t,
				"ipset: can't destroy set a: fake error",
				err.Error())
		})
	})
}

func Test_Swap(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		setupCmd()
		defer teardownCmd()

		assert.Nil(t, Swap("a", "b"))
	})

	t.Run("error", func(t *testing.T) {
		setupCmd(flag)
		defer teardownCmd()

		err := Swap("a", "b")
		require.Error(t, err)
		assert.Equal(t,
			"ipset: can't swap from a to b: fake error",
			err.Error())
	})
}
