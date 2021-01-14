package ipset

import (
	"sync"
	"time"
)

// Option is for ipset commands
type Option func(opt *options)

type options struct {
	exist           bool
	resolve         bool
	timeout         time.Duration
	counters        bool
	countersPackets uint
	countersBytes   uint
	comment         bool
	commentContent  string
	skbinfo         bool
	skbmark         string
	skbprio         string
	skbqueue        uint
	hashSize        uint
	maxElem         uint
	family          NetFamily
	nomatch         bool
	forceadd        bool
	ipRange         string
	portRange       string
	netmask         byte
	markmask        uint32
	listSize        uint
}

func (o *options) apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

var optionsPool = sync.Pool{
	New: func() interface{} {
		return &options{}
	},
}

func acquireOptions() *options {
	return optionsPool.Get().(*options)
}

func releaseOptions(o *options) {
	o.timeout = 0
	o.exist = false
	o.resolve = false
	o.counters = false
	o.countersPackets = 0
	o.countersBytes = 0
	o.comment = false
	o.commentContent = ""
	o.skbinfo = false
	o.skbmark = ""
	o.skbprio = ""
	o.skbqueue = 0
	o.hashSize = 0
	o.maxElem = 0
	o.family = ""
	o.nomatch = false
	o.forceadd = false
	o.netmask = 0
	o.markmask = 0
	o.listSize = 0
	o.ipRange = ""
	o.portRange = ""
	optionsPool.Put(o)
}

// Timeout option is used for create and add command.
// All set types supports the optional timeout parameter
// when creating a set and adding entries. The value of
// the timeout parameter for the create command means
// the default timeout value (in seconds) for new entries.
// If a set is created with timeout support, then the
// same timeout option can be used to specify non-default
// timeout values when adding entries. The largest possible
// timeout value is 2147483 seconds.
//
// Zero timeout value means the entry is added permanent
// to the set. The timeout value of already added
// elements can be changed by re-adding the element using
// the -exist option. Example:
//
//      ipset create test hash:ip timeout 300
//
//      ipset add test 192.168.0.1 timeout 60
//
//      ipset -exist add test 192.168.0.1 timeout 600
//
// When listing the set, the number of entries printed
// in the header might be larger than the listed number
// of entries for sets with the timeout extensions: the
// number of entries in the set is updated when elements
// added/deleted to the set and periodically when the
// garbage collector evicts the timed out entries.
func Timeout(timeout time.Duration) Option {
	return func(opt *options) {
		opt.timeout = timeout
	}
}

// Exist option ignores errors when exactly the same set is to
// be created or already added entry is added or missing
// entry is deleted.
func Exist(exist bool) Option {
	return func(opt *options) {
		opt.exist = exist
	}
}

// Resolve option is for listing sets, enforce action lookup. The
// program will try to display the IP entries resolved to host
// names which requires slow DNS lookups.
func Resolve(resolve bool) Option {
	return func(opt *options) {
		opt.resolve = resolve
	}
}

// Comment option is used for create and add command
// All set types support the optional comment extension.
// Enabling this extension on an ipset enables you to
// annotate an ipset entry with an arbitrary string. This
// string is completely ignored by both the kernel and
// ipset itself and is purely for providing a convenient
// means to document the reason for an entry's existence.
// Comments must not contain any quotation marks and the
// usual escape character (\) has no meaning.
//
// For example, the following shell command is illegal:
//
//      ipset add foo 1.1.1.1 comment "this comment is \"bad\""
//
// In the above, your shell will of course escape the
// quotation marks and ipset will see the quote marks in
// the argument for the comment, which will result in a
// parse error. If you are writing your own system, you
// should avoid creating comments containing a quotation
// mark if you do not want to break "ipset save" and
// "ipset restore", nonetheless, the kernel will not stop
// you from doing so. The following is perfectly acceptable:
//
//      ipset create foo hash:ip comment
//
//      ipset add foo 192.168.1.1/24 comment "allow access to SMB share on \\\\fileserv\\"
//
// the above would appear as: "allow access to SMB share on \\fileserv\"
func Comment(comment bool) Option {
	return func(opt *options) {
		opt.comment = comment
	}
}

// CommentContent is used for add command. And the set
// must be created with comment option.
func CommentContent(commentContent string) Option {
	return func(opt *options) {
		opt.commentContent = commentContent
	}
}

// Counters is used for create command.
// All set types support the optional counters option when
// creating a set. If the option is specified then the set
// is created with packet and byte counters per element
// support.
// The packet and byte counters are initialized to zero when
// the elements are (re-)added to the set, unless the packet
// and byte counter values are explicitly specified by the
// packets and bytes options. An example when an element is
// added to a set with non-zero counter values:
//
//      ipset create foo hash:ip counters
//
//      ipset add foo 192.168.1.1 packets 42 bytes 1024
func Counters(counters bool) Option {
	return func(opt *options) {
		opt.counters = counters
	}
}

// Packets option is used with Counters option
func Packets(packets uint) Option {
	return func(opt *options) {
		opt.countersPackets = packets
	}
}

// Bytes option is used with Counters option
func Bytes(b uint) Option {
	return func(opt *options) {
		opt.countersBytes = b
	}
}

// Skbinfo option is used for create command.
// All set types support the optional skbinfo extension.
// This extension allows you to store the metainfo
// (firewall mark, tc class and hardware queue) with every
// entry and map it to packets by usage of SET netfilter
// target with --map-set option. See skbmark, skbprio,
// skbqueue options for more detail. Example:
//
//      ipset create foo hash:ip skbinfo
//
//      ipset add foo skbmark 0x1111/0xff00ffff skbprio 1:10 skbqueue 10
func Skbinfo(skbinfo bool) Option {
	return func(opt *options) {
		opt.skbinfo = skbinfo
	}
}

// Skbmark option should be used with Skbinfo extension.
// Its format is:
//      MARK or MARK/MASK
// where MARK and MASK are 32bit hex numbers with 0x prefix.
// If only mark is specified mask 0xffffffff are used.
func Skbmark(skbmark string) Option {
	return func(opt *options) {
		opt.skbmark = skbmark
	}
}

// Skbprio option should be used with Skbinfo extension.
// It has tc class format:
//      MAJOR:MINOR
// where major and minor numbers are hex without 0x prefix.
func Skbprio(skbprio string) Option {
	return func(opt *options) {
		opt.skbprio = skbprio
	}
}

// Skbqueue option should be used with Skbinfo extension.
// It is just decimal number.
func Skbqueue(skbqueue uint) Option {
	return func(opt *options) {
		opt.skbqueue = skbqueue
	}
}

// HashSize option is valid for the create command of all
// hash type sets. It defines the initial hash size for the
// set, default is 1024. The hash size must be a power of
// two, the kernel automatically rounds up non power of
// two hash sizes to the first correct value.  Example:
//
//      ipset create test hash:ip hashsize 1536
func HashSize(hashSize uint) Option {
	return func(opt *options) {
		opt.hashSize = hashSize
	}
}

// MaxElem option is valid for the create command of all
// hash type sets. It does define the maximal number of
// elements which can be stored in the set, default 65536.
// Example:
//
//      ipset create test hash:ip maxelem 2048.
func MaxElem(maxElem uint) Option {
	return func(opt *options) {
		opt.maxElem = maxElem
	}
}

// NetFamily defines the protocol family of the IP addresses
type NetFamily string

const (
	// Inet indicates IPv4
	Inet NetFamily = "inet"
	// Inet6 indicates IPv6
	Inet6 NetFamily = "inet6"
)

// Family option is valid for the create command of all hash
// type sets except for hash:mac. It defines the protocol
// family of the IP addresses to be stored in the set. The
// default is inet, i.e IPv4. Another is inet6, i.e IPv6.
// For the inet family one can add or delete multiple entries
// by specifying a range or a network of IPv4 addresses in
// the IP address part of the entry:
//
//      ipaddr := { ip | fromaddr-toaddr | ip/cidr }
//
//      netaddr := { fromaddr-toaddr | ip/cidr }
//
// Example:
//
//      ipset create test hash:ip family inet6
func Family(family NetFamily) Option {
	return func(opt *options) {
		opt.family = family
	}
}

// Nomatch option is supported for the hash set types which
// can store net type of data (i.e. hash:*net*) when adding
// entries. When matching elements in the set, entries marked
// as nomatch are skipped as if those were not added to the
// set, which makes possible to build up sets with exceptions.
// See the example at hash type hash:net.
//
// When elements are tested by ipset, the nomatch flags are
// taken into account. If one wants to test the existence of
// an element marked with nomatch in a set, then the flag
// must be specified too.
func Nomatch(nomatch bool) Option {
	return func(opt *options) {
		opt.nomatch = nomatch
	}
}

// Forceadd option is supported for all hash set types
// when creating a set. When sets created with this option
// become full the next addition to the set may succeed and
// evict a random entry from the set.
//
//      ipset create foo hash:ip forceadd
func Forceadd(forceadd bool) Option {
	return func(opt *options) {
		opt.forceadd = forceadd
	}
}

// Netmask option is for ip datatype. Network addresses will be
// stored in the set instead of IP host addresses. The cidr
// prefix value must be between 1-32 for IPv4 and between 1-128
// for IPv6.
func Netmask(netmask byte) Option {
	return func(opt *options) {
		opt.netmask = netmask
	}
}

// Markmask option is for HashIpMark set type. It allows you to
// set bits you are interested in the packet mark. This values is
// then used to perform bitwise AND operation for every mark
// added. markmask can be any value between 1 and 4294967295, by
// default all 32 bits are set.
func Markmask(markmask uint32) Option {
	return func(opt *options) {
		opt.markmask = markmask
	}
}

// ListSize option is for ListSet set type. It's size of the list,
// the default is 8.
func ListSize(listSize uint) Option {
	return func(opt *options) {
		opt.listSize = listSize
	}
}

// IpRange option should be used with BitmapIp and BitmapIpMac set
// type. Creating the set from the specified inclusive address
// range expressed in an IPv4 address range or network. The size
// of the range cannot exceed the limit of maximum 65536 entries.
func IpRange(ipRange string) Option {
	return func(opt *options) {
		opt.ipRange = ipRange
	}
}

// PortRange option should be used with BitmapPort set type.
// Creating the set from the specified inclusive port range.
func PortRange(portRange string) Option {
	return func(opt *options) {
		opt.portRange = portRange
	}
}
