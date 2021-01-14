package ipset

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	testActions = []string{
		_create,
		_add,
		_del,
		_test,
		_destroy,
		_list,
		_save,
		_restore,
		_flush,
		_rename,
		_swap,
	}

	testSetTypes = []SetType{
		BitmapIp,
		BitmapIpMac,
		BitmapPort,
		HashIp,
		HashMac,
		HashIpMac,
		HashNet,
		HashNetNet,
		HashIpPort,
		HashNetPort,
		HashIpPortIp,
		HashIpPortNet,
		HashIpMark,
		HashNetPortNet,
		HashNetIface,
		ListSet,
	}
)

func Test_Options_Exist(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without exist", func(t *testing.T) {
			args := c.appendArgs(nil, Exist(false))
			assert.Len(t, args, 0)
		})

		if c.needExist() {
			t.Run(action+" need exist", func(t *testing.T) {
				args := c.appendArgs(nil, Exist(true))
				assert.Equal(t, _exist, args[0])
			})
		} else {
			t.Run(action+" ignore exist", func(t *testing.T) {
				args := c.appendArgs(nil, Exist(true))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Resolve(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without resolve", func(t *testing.T) {
			args := c.appendArgs(nil, Resolve(false))
			assert.Len(t, args, 0)
		})

		if c.needResolve() {
			t.Run(action+" need resolve", func(t *testing.T) {
				args := c.appendArgs(nil, Resolve(true))
				assert.Equal(t, _resolve, args[0])
			})
		} else {
			t.Run(action+" ignore resolve", func(t *testing.T) {
				args := c.appendArgs(nil, Resolve(true))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Timeout(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without timeout", func(t *testing.T) {
			args := c.appendArgs(nil, Timeout(0))
			assert.Len(t, args, 0)
		})

		if c.needTimeout() {
			t.Run(action+" need timeout", func(t *testing.T) {
				args := c.appendArgs(nil, Timeout(time.Second))
				assert.Equal(t, _timeout, args[0])
				assert.Equal(t, "1", args[1])
			})
		} else {
			t.Run(action+" ignore timeout", func(t *testing.T) {
				args := c.appendArgs(nil, Timeout(time.Second))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Counters(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without counters", func(t *testing.T) {
			args := c.appendArgs(nil, Counters(false))
			assert.Len(t, args, 0)
		})

		if c.needCounters() {
			t.Run(action+" need counters", func(t *testing.T) {
				args := c.appendArgs(nil, Counters(true))
				assert.Equal(t, _counters, args[0])
			})
		} else {
			t.Run(action+" ignore counters", func(t *testing.T) {
				args := c.appendArgs(nil, Counters(true))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Packets(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without packets", func(t *testing.T) {
			args := c.appendArgs(nil, Packets(0))
			assert.Len(t, args, 0)
		})

		if c.onlyAdd() {
			t.Run(action+" need packets", func(t *testing.T) {
				args := c.appendArgs(nil, Packets(1))
				assert.Equal(t, _packets, args[0])
				assert.Equal(t, "1", args[1])
			})
		} else {
			t.Run(action+" ignore packets", func(t *testing.T) {
				args := c.appendArgs(nil, Packets(1))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Bytes(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without bytes", func(t *testing.T) {
			args := c.appendArgs(nil, Bytes(0))
			assert.Len(t, args, 0)
		})

		if c.onlyAdd() {
			t.Run(action+" need bytes", func(t *testing.T) {
				args := c.appendArgs(nil, Bytes(1))
				assert.Equal(t, _bytes, args[0])
				assert.Equal(t, "1", args[1])
			})
		} else {
			t.Run(action+" ignore bytes", func(t *testing.T) {
				args := c.appendArgs(nil, Bytes(1))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Comment(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without comment", func(t *testing.T) {
			args := c.appendArgs(nil, Comment(false))
			assert.Len(t, args, 0)
		})

		if c.onlyCreate() {
			t.Run(action+" need comment", func(t *testing.T) {
				args := c.appendArgs(nil, Comment(true))
				assert.Equal(t, _comment, args[0])
			})
		} else {
			t.Run(action+" ignore comment", func(t *testing.T) {
				args := c.appendArgs(nil, Comment(true))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_CommentContent(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without comment content", func(t *testing.T) {
			args := c.appendArgs(nil, CommentContent(""))
			assert.Len(t, args, 0)
		})

		if c.onlyAdd() {
			t.Run(action+" need comment content", func(t *testing.T) {
				args := c.appendArgs(nil, CommentContent("comment"))
				assert.Equal(t, _comment, args[0])
				assert.Equal(t, "comment", args[1])
			})
		} else {
			t.Run(action+" ignore comment content", func(t *testing.T) {
				args := c.appendArgs(nil, CommentContent("comment"))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Skbinfo(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without skbinfo", func(t *testing.T) {
			args := c.appendArgs(nil, Skbinfo(false))
			assert.Len(t, args, 0)
		})

		if c.onlyCreate() {
			t.Run(action+" need skbinfo", func(t *testing.T) {
				args := c.appendArgs(nil, Skbinfo(true))
				assert.Equal(t, _skbinfo, args[0])
			})
		} else {
			t.Run(action+" ignore skbinfo", func(t *testing.T) {
				args := c.appendArgs(nil, Skbinfo(true))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Skbmark(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without skbmark", func(t *testing.T) {
			args := c.appendArgs(nil, Skbmark(""))
			assert.Len(t, args, 0)
		})

		if c.onlyAdd() {
			t.Run(action+" need skbmark", func(t *testing.T) {
				args := c.appendArgs(nil, Skbmark("skbmark"))
				assert.Equal(t, _skbmark, args[0])
				assert.Equal(t, "skbmark", args[1])
			})
		} else {
			t.Run(action+" ignore skbmark", func(t *testing.T) {
				args := c.appendArgs(nil, Skbmark("skbmark"))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Skbprio(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without skbprio", func(t *testing.T) {
			args := c.appendArgs(nil, Skbprio(""))
			assert.Len(t, args, 0)
		})

		if c.onlyAdd() {
			t.Run(action+" need skbprio", func(t *testing.T) {
				args := c.appendArgs(nil, Skbprio("skbprio"))
				assert.Equal(t, _skbprio, args[0])
				assert.Equal(t, "skbprio", args[1])
			})
		} else {
			t.Run(action+" ignore skbprio", func(t *testing.T) {
				args := c.appendArgs(nil, Skbprio("skbprio"))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Skbqueue(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without skbqueue", func(t *testing.T) {
			args := c.appendArgs(nil, Skbqueue(0))
			assert.Len(t, args, 0)
		})

		if c.onlyAdd() {
			t.Run(action+" need skbqueue", func(t *testing.T) {
				args := c.appendArgs(nil, Skbqueue(1))
				assert.Equal(t, _skbqueue, args[0])
				assert.Equal(t, "1", args[1])
			})
		} else {
			t.Run(action+" ignore skbqueue", func(t *testing.T) {
				args := c.appendArgs(nil, Skbqueue(1))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Nomatch(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		for _, setType := range testSetTypes {
			c := getFakeCmd(action, setType)
			t.Run(action+" "+string(setType)+" without nomatch", func(t *testing.T) {
				args := c.appendArgs(nil, Nomatch(false))
				assert.Len(t, args, 0)
			})

			if c.needNomatch() {
				t.Run(action+" "+string(setType)+" need nomatch", func(t *testing.T) {
					args := c.appendArgs(nil, Nomatch(true))
					assert.Equal(t, _nomatch, args[0])
				})
			} else {
				t.Run(action+" "+string(setType)+" ignore nomatch", func(t *testing.T) {
					args := c.appendArgs(nil, Nomatch(true))
					assert.Len(t, args, 0)
				})
			}
		}
	}
}

func Test_Options_Forceadd(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without forceadd", func(t *testing.T) {
			args := c.appendArgs(nil, Forceadd(false))
			assert.Len(t, args, 0)
		})

		if c.onlyCreate() {
			t.Run(action+" need forceadd", func(t *testing.T) {
				args := c.appendArgs(nil, Forceadd(true))
				assert.Equal(t, _forceadd, args[0])
			})
		} else {
			t.Run(action+" ignore forceadd", func(t *testing.T) {
				args := c.appendArgs(nil, Forceadd(true))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Family(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		for _, setType := range testSetTypes {
			c := getFakeCmd(action, setType)
			t.Run(action+" "+string(setType)+" without family", func(t *testing.T) {
				args := c.appendArgs(nil, Family(""))
				assert.Len(t, args, 0)
			})

			if c.needFamily() {
				t.Run(action+" "+string(setType)+" need family", func(t *testing.T) {
					args := c.appendArgs(nil, Family("inet"))
					assert.Equal(t, _family, args[0])
					assert.Equal(t, "inet", args[1])
				})
			} else {
				t.Run(action+" "+string(setType)+" ignore family", func(t *testing.T) {
					args := c.appendArgs(nil, Family("inet"))
					assert.Len(t, args, 0)
				})
			}
		}
	}
}

func Test_Options_HashSize(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		for _, setType := range testSetTypes {
			c := getFakeCmd(action, setType)
			t.Run(action+" "+string(setType)+" without hashsize", func(t *testing.T) {
				args := c.appendArgs(nil, HashSize(0))
				assert.Len(t, args, 0)
			})

			if c.needHash() {
				t.Run(action+" "+string(setType)+" need hashsize", func(t *testing.T) {
					args := c.appendArgs(nil, HashSize(1))
					assert.Equal(t, _hashsize, args[0])
					assert.Equal(t, "1", args[1])
				})
			} else {
				t.Run(action+" "+string(setType)+" ignore hashsize", func(t *testing.T) {
					args := c.appendArgs(nil, HashSize(1))
					assert.Len(t, args, 0)
				})
			}
		}
	}
}

func Test_Options_MaxElem(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		for _, setType := range testSetTypes {
			c := getFakeCmd(action, setType)
			t.Run(action+" "+string(setType)+" without maxelem", func(t *testing.T) {
				args := c.appendArgs(nil, MaxElem(0))
				assert.Len(t, args, 0)
			})

			if c.needHash() {
				t.Run(action+" "+string(setType)+" need maxelem", func(t *testing.T) {
					args := c.appendArgs(nil, MaxElem(1))
					assert.Equal(t, _maxelem, args[0])
					assert.Equal(t, "1", args[1])
				})
			} else {
				t.Run(action+" "+string(setType)+" ignore maxelem", func(t *testing.T) {
					args := c.appendArgs(nil, MaxElem(1))
					assert.Len(t, args, 0)
				})
			}
		}
	}
}

func Test_Options_Netmask(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		for _, setType := range testSetTypes {
			c := getFakeCmd(action, setType)
			t.Run(action+" "+string(setType)+" without netmask", func(t *testing.T) {
				args := c.appendArgs(nil, Netmask(0))
				assert.Len(t, args, 0)
			})

			if c.needNetmask() {
				t.Run(action+" "+string(setType)+" need netmask", func(t *testing.T) {
					args := c.appendArgs(nil, Netmask(1))
					assert.Equal(t, _netmask, args[0])
					assert.Equal(t, "1", args[1])
				})
			} else {
				t.Run(action+" "+string(setType)+" ignore netmask", func(t *testing.T) {
					args := c.appendArgs(nil, Netmask(1))
					assert.Len(t, args, 0)
				})
			}
		}
	}
}

func Test_Options_Markmask(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		for _, setType := range testSetTypes {
			c := getFakeCmd(action, setType)
			t.Run(action+" "+string(setType)+" without markmask", func(t *testing.T) {
				args := c.appendArgs(nil, Markmask(0))
				assert.Len(t, args, 0)
			})

			if c.needMarkmask() {
				t.Run(action+" "+string(setType)+" need markmask", func(t *testing.T) {
					args := c.appendArgs(nil, Markmask(1))
					assert.Equal(t, _markmask, args[0])
					assert.Equal(t, "1", args[1])
				})
			} else {
				t.Run(action+" "+string(setType)+" ignore markmask", func(t *testing.T) {
					args := c.appendArgs(nil, Markmask(1))
					assert.Len(t, args, 0)
				})
			}
		}
	}
}

func Test_Options_IpRange(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		for _, setType := range testSetTypes {
			c := getFakeCmd(action, setType)
			t.Run(action+" "+string(setType)+" without ip range", func(t *testing.T) {
				args := c.appendArgs(nil, IpRange(""))
				assert.Len(t, args, 0)
			})

			if c.needIpRange() {
				t.Run(action+" "+string(setType)+" need ip range", func(t *testing.T) {
					args := c.appendArgs(nil, IpRange("1.1.1.1/24"))
					assert.Equal(t, _range, args[0])
					assert.Equal(t, "1.1.1.1/24", args[1])
				})
			} else {
				t.Run(action+" "+string(setType)+" ignore ip range", func(t *testing.T) {
					args := c.appendArgs(nil, IpRange("1.1.1.1/24"))
					assert.Len(t, args, 0)
				})
			}
		}
	}
}

func Test_Options_PortRange(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		for _, setType := range testSetTypes {
			c := getFakeCmd(action, setType)
			t.Run(action+" "+string(setType)+" without port range", func(t *testing.T) {
				args := c.appendArgs(nil, PortRange(""))
				assert.Len(t, args, 0)
			})

			if c.needPortRange() {
				t.Run(action+" "+string(setType)+" need port range", func(t *testing.T) {
					args := c.appendArgs(nil, PortRange("1000-2000"))
					assert.Equal(t, _range, args[0])
					assert.Equal(t, "1000-2000", args[1])
				})
			} else {
				t.Run(action+" "+string(setType)+" ignore port range", func(t *testing.T) {
					args := c.appendArgs(nil, PortRange("1000-2000"))
					assert.Len(t, args, 0)
				})
			}
		}
	}
}

func Test_Options_ListSize(t *testing.T) {
	t.Parallel()

	for _, action := range testActions {
		for _, setType := range testSetTypes {
			c := getFakeCmd(action, setType)
			t.Run(action+" "+string(setType)+" without list size", func(t *testing.T) {
				args := c.appendArgs(nil, ListSize(0))
				assert.Len(t, args, 0)
			})

			if c.needListSize() {
				t.Run(action+" "+string(setType)+" need list size", func(t *testing.T) {
					args := c.appendArgs(nil, ListSize(1))
					assert.Equal(t, _size, args[0])
					assert.Equal(t, "1", args[1])
				})
			} else {
				t.Run(action+" "+string(setType)+" ignore list size", func(t *testing.T) {
					args := c.appendArgs(nil, ListSize(1))
					assert.Len(t, args, 0)
				})
			}
		}
	}
}

func getFakeCmd(action string, setType ...SetType) *cmd {
	st := HashIp
	if len(setType) > 0 {
		st = setType[0]
	}
	return &cmd{
		action:  action,
		name:    "test",
		setType: st,
	}
}
