package ipset

// SetType indicates a set type comprises of the storage method
// by which the data is stored and the data type(s) which are
// stored in the set. Therefore the TYPENAME parameter of the
// create command follows the syntax
//      TYPENAME := method:datatype[,datatype[,datatype]]
// where the current list of the methods are
//      bitmap, hash, and list
// and the possible data types are
//      ip, net, mac, port and iface.
// The dimension of a set is equal to the number of data types
// in its type action.
//
// When adding, deleting or testing entries in a set, the same
// comma separated data syntax must be used for the entry
// parameter of the commands, i.e
//      ipset add foo ipaddr,portnum,ipaddr
// If host names or service names with dash in the action are
// used instead of IP addresses or service numbers, then the
// host action or service action must be enclosed in square brackets.
// Example:
//      ipset add foo [test-hostname],[ftp-data]
// In the case of host names the DNS resolver is called internally
// by ipset but if it returns multiple IP addresses, only the
// first one is used.
//
// The bitmap and list types use a fixed sized storage. The hash
// types use a hash to store the elements. In order to avoid
// clashes in the hash, a limited number of chaining, and if that
// is exhausted, the doubling of the hash size is performed when
// adding entries by the ipset command. When entries added by
// the SET target of iptables/ip6tables, then the hash size is
// fixed and the set won't be duplicated, even if the new entry
// cannot be added to the set.
type SetType string

// BitmapIp set type uses a memory range to store either IPv4
// host (default) or IPv4 network addresses. A bitmap:ip type
// of set can store up to 65536 entries.
//
//      CREATE-OPTIONS := range fromip-toip|ip/cidr
//              [ netmask cidr ] [ timeout value ]
//              [ counters ] [ comment ] [ skbinfo ]
//
//      ADD-ENTRY := { ip | fromip-toip | ip/cidr }
//
//      ADD-OPTIONS := [ timeout value ] [ packets value ]
//                     [ bytes value ] [ comment string ]
//                     [ skbmark value ] [ skbprio value ]
//                     [ skbqueue value ]
//
//      DEL-ENTRY := { ip | fromip-toip | ip/cidr }
//
//      TEST-ENTRY := ip
//
// Mandatory create options:
//
//      range fromip-toip|ip/cidr
//
// Create the set from the specified inclusive address range
// expressed in an IPv4 address range or network. The size of
// the range (in entries) cannot exceed the limit of maximum
// 65536 elements.
//
// Optional create options:
//
//      netmask cidr
//
// When the optional netmask parameter specified, network
// addresses will be stored in the set instead of IP host
// addresses. The cidr prefix value must be between 1-32.
// An IP address will be in the set if the network address,
// which is resulted by masking the address with the specified
// netmask, can be found in the set.
//
// The BitmapIp type supports adding or deleting multiple
// entries in one command.
//
// Examples:
//
//      ipset create foo bitmap:ip range 192.168.0.0/16
//
//      ipset add foo 192.168.1/24
//
//      ipset test foo 192.168.1.1
//
const BitmapIp SetType = "bitmap:ip"

// BitmapIpMac set type uses a memory range to store IPv4 and
// a MAC address pairs. A BitmapIpMac type of set can store
// up to 65536 entries.
//
//      CREATE-OPTIONS := range fromip-toip|ip/cidr
//          [ timeout value ] [ counters ] [ comment ]
//          [ skbinfo ]
//
//      ADD-ENTRY := ip[,macaddr]
//
//      ADD-OPTIONS := [ timeout value ] [ packets value ]
//                     [ bytes value ] [ comment string ]
//                     [ skbmark value ] [ skbprio value ]
//                     [ skbqueue value ]
//
//      DEL-ENTRY := ip[,macaddr]
//
//      TEST-ENTRY := ip[,macaddr]
//
// Mandatory options to use when creating a BitmapIpMac type
// of set:
//
//       range fromip-toip|ip/cidr
//
// Create the set from the specified inclusive address range
// expressed in an IPv4 address range or network. The size
// of the range cannot exceed  the  limit  of  maximum 65536
// entries.
//
// The BitmapIpMac type is exceptional in the sense that the
// MAC part can be left out when adding/deleting/testing
// entries in the set. If we add an entry without the MAC
// address specified, then when the first time the entry is
// matched by the kernel, it will automatically fill out the
// missing MAC address with the source MAC address from the
// packet. If the entry was specified with a timeout value,
// the timer starts off when the IP and MAC address pair is
// complete.
//
// The BitmapIpMac type of sets require two src/dst
// parameters of the set match and SET target netfilter
// kernel modules and the second one must be src to match,
// add or delete entries, because the set match and SET
// target have access to the source MAC address only.
//
// Examples:
//
//      ipset create foo bitmap:ip,mac range 192.168.0.0/16
//
//      ipset add foo 192.168.1.1,12:34:56:78:9A:BC
//
//      ipset test foo 192.168.1.1
const BitmapIpMac SetType = "bitmap:ip,mac"

// BitmapPort set type uses a memory range to store port numbers and such a set can store up to 65536 ports.
//
//      CREATE-OPTIONS := range fromport-toport
//      [ timeout value ] [ counters ] [ comment ] [ skbinfo ]
//
//      ADD-ENTRY := { [proto:]port |
//                          [proto:]fromport-toport }
//
//      ADD-OPTIONS := [ timeout value ] [ packets value ]
//                     [ bytes value ] [ comment string ]
//                     [ skbmark value ] [ skbprio value ]
//                     [ skbqueue value ]
//
//      DEL-ENTRY := { [proto:]port |
//                          [proto:]fromport-toport }
//
//      TEST-ENTRY := [proto:]port
//
// Mandatory options to use when creating a BitmapPort type
// of set:
//
//      range [proto:]fromport-toport
//
// Create the set from the specified inclusive port range.
//
// The set match and SET target netfilter kernel modules
// interpret the stored numbers as TCP or UDP port numbers.
//
// proto only needs to be specified if a service action is
// used, and that action does not exist as a TCP service.
//
// Examples:
//
//      ipset create foo bitmap:port range 0-1024
//
//      ipset add foo 80
//
//      ipset test foo 80
//
//      ipset del foo udp:[macon-udp]-[tn-tl-w2]
const BitmapPort SetType = "bitmap:port"

// HashIp set type uses a hash to store IP host addresses
// (default) or network addresses. Zero valued IP address
// cannot be stored in a HashIp type of set.
//
//      CREATE-OPTIONS := [ family { inet | inet6 } ] |
//                        [ hashsize value ] [ maxelem value ]
//                        [ netmask cidr ] [ timeout value ]
//                        [ counters ] [ comment ] [ skbinfo ]
//
//      ADD-ENTRY := ipaddr
//
//      ADD-OPTIONS := [ timeout value ] [ packets value ]
//                     [ bytes value ] [ comment string ]
//                     [ skbmark value ] [ skbprio value ]
//                     [ skbqueue value ]
//
//      DEL-ENTRY := ipaddr
//
//      TEST-ENTRY := ipaddr
//
// Optional create options:
//
//       netmask cidr
//
// When the optional netmask parameter specified, network
// addresses will be stored in the set instead of IP host
// addresses. The cidr prefix value must be between 1-32
// for IPv4 and between 1-128 for IPv6. An IP address will
// be in the set if the network address, which is resulted
// by masking the address with the netmask, can be found in
// the set.
// Examples:
//
//      ipset create foo hash:ip netmask 30
//
//      ipset add foo 192.168.1.0/24
//
//      ipset test foo 192.168.1.2
const HashIp SetType = "hash:ip"

// HashMac set type uses a hash to store MAC addresses. Zero
// valued MAC addresses cannot be stored in a HashMac type of
// set.
//
//      CREATE-OPTIONS := [ hashsize value ]
//                        [ maxelem value ] [ timeout value ]
//                        [ counters ] [ comment ] [ skbinfo ]
//
//      ADD-ENTRY := macaddr
//
//      ADD-OPTIONS := [ timeout value ] [ packets value ]
//                     [ bytes value ] [ comment string ]
//                     [ skbmark value ] [ skbprio value ]
//                     [ skbqueue value ]
//
//      DEL-ENTRY := macaddr
//
//      TEST-ENTRY := macaddr
//
// Examples:
//
//      ipset create foo hash:mac
//
//      ipset add foo 01:02:03:04:05:06
//
//      ipset test foo 01:02:03:04:05:06
const HashMac SetType = "hash:mac"

// HashIpMac set type uses a hash to store IP and a MAC
// address pairs. Zero valued MAC addresses cannot be stored
// in a Hash:IpMac type of set. For matches on destination
// MAC addresses, see COMMENTS below.
//
//      CREATE-OPTIONS  := [ family { inet | inet6 } ] |
//                         [ hashsize value ] [ maxelem value ]
//                         [ timeout value ] [ counters ]
//                         [ comment ] [ skbinfo ]
//
//      ADD-ENTRY := ipaddr,macaddr
//
//      ADD-OPTIONS := [ timeout value ] [ packets value ]
//                     [ bytes value ] [ comment string ]
//                     [ skbmark value ] [ skbprio value ]
//                     [ skbqueue value ]
//
//      DEL-ENTRY := ipaddr,macaddr
//
//      TEST-ENTRY := ipaddr,macaddr
//
// Examples:
//
//      ipset create foo hash:ip,mac
//
//      ipset add foo 1.1.1.1,01:02:03:04:05:06
//
//      ipset test foo 1.1.1.1,01:02:03:04:05:06
const HashIpMac SetType = "hash:ip,mac"

// HashNet set type uses a hash to store different sized IP
// network addresses. Network address with zero prefix size
// cannot be stored in this type of sets.
//
//      CREATE-OPTIONS := [ family { inet | inet6 } ] |
//                        [ hashsize value ]
//                        [ maxelem value ] [ timeout value ]
//                        [ counters ] [ comment ]
//                        [ skbinfo ]
//
//      ADD-ENTRY := netaddr
//
//      ADD-OPTIONS := [ timeout value ] [ nomatch ]
//                     [ packets value ] [ bytes value ]
//                     [ comment string ] [ skbmark value ]
//                     [ skbprio value ] [ skbqueue value ]
//
//      DEL-ENTRY := netaddr
//
//      TEST-ENTRY := netaddr
//
//      where netaddr := ip[/cidr]
//
// When adding/deleting/testing entries, if the cidr prefix
// parameter is not specified, then the host prefix value is
// assumed. When adding/deleting entries, the exact element is
// added/deleted and overlapping elements are not checked by
// the kernel. When testing entries, if a host address is
// tested, then the kernel tries to match the host address in
// the networks added to the set and reports the result
// accordingly.
//
// From the set netfilter match point of view the searching
// for a match always  starts  from  the smallest  size  of
// netblock (most specific prefix) to the largest  one (least
// specific prefix) added to the set. When adding/deleting IP
// addresses to the set by the SET netfilter target, it will
// be added/deleted by the most specific prefix which can be
// found in the set, or by the host prefix value if the set
// is empty.
//
// The lookup time grows linearly with the number of the
// different prefix values added to the set.
//
// Example:
//
//      ipset create foo hash:net
//
//      ipset add foo 192.168.0.0/24
//
//      ipset add foo 10.1.0.0/16
//
//      ipset add foo 192.168.0/24
//
//      ipset add foo 192.168.0/30 nomatch
//
// When matching the elements in the set above, all IP
// addresses will match from the networks 192.168.0.0/24,
// 10.1.0.0/16 and 192.168.0/24 except the ones from
// 192.168.0/30.
const HashNet SetType = "hash:net"

// HashNetNet set type uses a hash to store pairs of
// different sized IP network addresses. Bear in mind that
// the first parameter has precedence over the second, so a
// nomatch entry could be potentially be ineffective if a
// more specific first parameter existed with a suitable
// second parameter. Network address with zero prefix size
// cannot be stored in this type of set.
//
//      CREATE-OPTIONS := [ family { inet | inet6 } ] |
//                        [ hashsize value ] [ maxelem value ]
//                        [ timeout value ] [ counters ]
//                        [ comment ] [ skbinfo ]
//
//      ADD-ENTRY := netaddr,netaddr
//
//      ADD-OPTIONS := [ timeout value ] [ nomatch ]
//                     [ packets value ] [ bytes value ]
//                     [ comment string ] [ skbmark value ]
//                     [ skbprio value ] [ skbqueue value ]
//
//      DEL-ENTRY := netaddr,netaddr
//
//      TEST-ENTRY := netaddr,netaddr
//
//      where netaddr := ip[/cidr]
//
// When adding/deleting/testing entries, if the cidr prefix
// parameter is not specified, then the host prefix value is
// assumed. When adding/deleting entries, the exact element is
// added/deleted and overlapping elements are not checked by
// the kernel. When testing entries, if a host address is
// tested, then the kernel tries to match the host address in
// the networks added to the set and reports the result
// accordingly.
//
// From the set netfilter match point of view the searching
// for a match always starts from the smallest size of
// netblock (most specific prefix) to the largest one (least
// specific prefix) with the first param having precedence.
// When adding/deleting IP addresses to the set by the SET
// netfilter target, it will be added/deleted by the most
// specific prefix which can be found in the set, or by the
// host prefix value if the set is empty.
//
// The lookup time grows linearly with the number of the
// different prefix values added to the first parameter of
// the set. The number of secondary prefixes further
// increases this as the list of secondary prefixes is
// traversed per primary prefix.
//
// Example:
//
//      ipset create foo hash:net,net
//
//      ipset add foo 192.168.0.0/24,10.0.1.0/24
//
//      ipset add foo 10.1.0.0/16,10.255.0.0/24
//
//      ipset add foo 192.168.0/24,192.168.54.0-192.168.54.255
//
//      ipset add foo 192.168.0/30,192.168.64/30 nomatch
//
// When matching the elements in the set above, all IP
// addresses will match from the networks
//      192.168.0.0/24<->10.0.1.0/24,
//      10.1.0.0/16<->10.255.0.0/24,
//      192.168.0/24<->192.168.54.0/24,
// except the ones from
//      192.168.0/30<->192.168.64/30.
const HashNetNet SetType = "hash:net,net"

// HashIpPort set type uses a hash to store IP address and
// port number pairs.  The port number is interpreted
// together with a protocol (default TCP) and  zero  protocol
// number cannot be used.
//
//      CREATE-OPTIONS := [ family { inet | inet6 } ] |
//                        [ hashsize value ] [ maxelem value ]
//                        [ timeout value ] [ counters ]
//                        [ comment ] [ skbinfo ]
//
//      ADD-ENTRY := ipaddr,[proto:]port
//
//      ADD-OPTIONS := [ timeout value ] [ packets value ]
//                     [ bytes value ] [ comment string ]
//                     [ skbmark value ] [ skbprio value ]
//                     [ skbqueue value ]
//
//      DEL-ENTRY := ipaddr,[proto:]port
//
//      TEST-ENTRY := ipaddr,[proto:]port
//
// The [proto:]port part of the elements may be expressed in
// the following forms, where the range variations are valid
// when adding or deleting entries:
//
//      portname[-portname]
// TCP port or range of ports expressed in TCP portname
// identifiers from /etc/services
//
//       portnumber[-portnumber]
// TCP port or range of ports expressed in TCP port numbers
//
//       tcp|sctp|udp|udplite:portname|
//          portnumber[-portname|portnumber]
// TCP, SCTP, UDP or UDPLITE port or port range expressed in
// port action(s) or port number(s)
//
//       icmp:codename|type/code
// ICMP codename or type/code. The supported ICMP codename
// identifiers can always be listed by the help command.
//
//       icmpv6:codename|type/code
// ICMPv6 codename or type/code. The supported ICMPv6
// codename identifiers can always be listed by the help
// command.
//
//       proto:0
// All other protocols, as an identifier from /etc/protocols
// or number. The pseudo port number must be zero.
//
// The HashIpPort type of sets require two src/dst parameters
// of the set match and SET target kernel modules.
//
// Examples:
//
//      ipset create foo hash:ip,port
//
//      ipset add foo 192.168.1.0/24,80-82
//
//      ipset add foo 192.168.1.1,udp:53
//
//      ipset add foo 192.168.1.1,vrrp:0
//
//      ipset test foo 192.168.1.1,80
const HashIpPort SetType = "hash:ip,port"

// HashNetPort set type uses a hash to store different sized
// IP network address and port pairs. The port number is
// interpreted together with a protocol (default TCP) and zero
// protocol number cannot be used. Network address with zero
// prefix size is not accepted either.
//
//      CREATE-OPTIONS := [ family { inet | inet6 } ] |
//                        [ hashsize value ] [ maxelem value ]
//                        [ timeout value ] [ counters ]
//                        [ comment ] [ skbinfo ]
//
//      ADD-ENTRY := netaddr,[proto:]port
//
//      ADD-OPTIONS := [ timeout value ]  [ nomatch ]
//                     [ packets value ] [ bytes value ]
//                     [ comment string ] [ skbmark value ]
//                     [ skbprio value ] [ skbqueue value ]
//
//      DEL-ENTRY := netaddr,[proto:]port
//
//      TEST-ENTRY := netaddr,[proto:]port
//
//      where netaddr := ip[/cidr]
//
// For the netaddr part of the elements see the description at
// the HashNet set type. For the [proto:]port part of the
// elements see the description at the hash:ip,port set type.
//
// When adding/deleting/testing entries, if the cidr prefix
// parameter is not specified, then the host prefix value is
// assumed. When adding/deleting entries, the exact element is
// added/deleted and overlapping elements are not checked by
// the kernel. When testing entries, if a host address is
// tested, then the kernel tries to match the host address in
// the networks added to the set and reports the result
// accordingly.
//
// From the set netfilter match point of view the searching
// for a match always starts from the smallest size of
// netblock (most specific prefix) to the largest one (least
// specific prefix) added to the set. When adding/deleting IP
// addresses to the set by the SET netfilter target, it will
// be added/deleted by the most specific prefix which can be
// found in the set, or by the host prefix value if the set is
// empty.
//
// The lookup time grows linearly with the number of the
// different prefix values added to the set.
//
// Examples:
//
//      ipset create foo hash:net,port
//
//      ipset add foo 192.168.0/24,25
//
//      ipset add foo 10.1.0.0/16,80
//
//      ipset test foo 192.168.0/24,25
const HashNetPort SetType = "hash:net,port"

// HashIpPortIp set type uses a hash to store IP address, port
// number and a second IP address triples. The port number is
// interpreted together with a protocol (default TCP) and zero
// protocol number cannot be used.
//
//      CREATE-OPTIONS := [ family { inet | inet6 } ] |
//                        [ hashsize value ] [ maxelem value ]
//                        [ timeout value ] [ counters ]
//                        [ comment ] [ skbinfo ]
//
//      ADD-ENTRY := ipaddr,[proto:]port,ip
//
//      ADD-OPTIONS := [ timeout value ] [ packets value ]
//                     [ bytes value ] [ comment string ]
//                     [ skbmark value ] [ skbprio value ]
//                     [ skbqueue value ]
//
//      DEL-ENTRY := ipaddr,[proto:]port,ip
//
//      TEST-ENTRY := ipaddr,[proto:]port,ip
//
// For the first ipaddr and [proto:]port parts of the elements
// see the descriptions at the hash:ip,port set type.
//
// The hash:ip,port,ip type of sets require three src/dst
// parameters of the set match and SET target kernel modules.
//
// Examples:
//
//      ipset create foo hash:ip,port,ip
//
//      ipset add foo 192.168.1.1,80,10.0.0.1
//
//      ipset test foo 192.168.1.1,udp:53,10.0.0.1
const HashIpPortIp SetType = "hash:ip,port,ip"

// HashIpPortNet set type uses a hash to store IP address,
// port number and IP network address triples. The port number
// is interpreted together with a protocol (default TCP) and
// zero protocol number cannot be used. Network address with
// zero prefix size cannot be stored either.
//
//      CREATE-OPTIONS := [ family { inet | inet6 } ] |
//                        [ hashsize value ] [ maxelem value ]
//                        [ timeout value ] [ counters ]
//                        [ comment ] [ skbinfo ]
//
//      ADD-ENTRY := ipaddr,[proto:]port,netaddr
//
//      ADD-OPTIONS := [ timeout value ]  [ nomatch ]
//                     [ packets value ] [ bytes value ]
//                     [ comment string ] [ skbmark value ]
//                     [ skbprio value ] [ skbqueue value ]
//
//      DEL-ENTRY := ipaddr,[proto:]port,netaddr
//
//      TEST-ENTRY := ipaddr,[proto:]port,netaddr
//
//      where netaddr := ip[/cidr]
//
// For the ipaddr and [proto:]port parts of the elements see
// the descriptions at the hash:ip,port set type. For the
// netaddr part of the elements see the description at the
// HashNet set type.
//
// From the set netfilter match point of view the searching
// for a match always starts from the smallest size of
// netblock (most specific cidr) to the largest one (least
// specific cidr) added to the set. When adding/deleting
// triples to the set by the SET netfilter target, it will be
// added/deleted by the most specific cidr which can be found
// in the set, or by the host cidr value if the set is empty.
//
// The lookup time grows linearly with the number of the
// different cidr values added to the set.
//
// The HashIpPortNet type of sets require three src/dst
// parameters of the set match and SET target kernel modules.
//
// Examples:
//
//      ipset create foo hash:ip,port,net
//
//      ipset add foo 192.168.1,80,10.0.0/24
//
//      ipset add foo 192.168.2,25,10.1.0.0/16
//
//      ipset test foo 192.168.1,80.10.0.0/24
const HashIpPortNet SetType = "hash:ip,port,net"

// HashIpMark set type uses a hash to store IP address and
// packet mark pairs.
//
//      CREATE-OPTIONS := [ family { inet | inet6 } ] |
//                        [ markmask value ]
//                        [ hashsize value ] [ maxelem value ]
//                        [ timeout value ] [ counters ]
//                        [ comment ] [ skbinfo ]
//
//      ADD-ENTRY := ipaddr,mark
//
//      ADD-OPTIONS := [ timeout value ] [ packets value ]
//                     [ bytes value ] [ comment string ]
//                     [ skbmark value ] [ skbprio value ]
//                     [ skbqueue value ]
//
//      DEL-ENTRY := ipaddr,mark
//
//      TEST-ENTRY := ipaddr,mark
//
// Optional create options:
//
//       markmask value
//
// Allows you to set bits you are interested in the packet
// mark. This values is then used to perform bitwise AND
// operation for every mark added. markmask can be any value
// between 1 and 4294967295, by default all 32 bits are set.
//
// The mark can be any value between 0 and 4294967295.
//
// The HashIpMark type of sets require two src/dst parameters
// of the set match and SET target kernel modules.
//
// Examples:
//
//      ipset create foo hash:ip,mark
//
//      ipset add foo 192.168.1.0/24,555
//
//      ipset add foo 192.168.1.1,0x63
//
//      ipset add foo 192.168.1.1,111236
const HashIpMark SetType = "hash:ip,mark"

// HashNetPortNet set type behaves similarly to HashIpPortNet
// but accepts a cidr value for both the first and last
// parameter. Either subnet is permitted to be a /0 should you
// wish to match port between all destinations.
//
//      CREATE-OPTIONS := [ family { inet | inet6 } ] |
//                        [ hashsize value ] [ maxelem value ]
//                        [ timeout value ] [ counters ]
//                        [ comment ] [ skbinfo ]
//
//      ADD-ENTRY := netaddr,[proto:]port,netaddr
//
//      ADD-OPTIONS := [ timeout value ]  [ nomatch ]
//                     [ packets value ] [ bytes value ]
//                     [ comment string ] [ skbmark value ]
//                     [ skbprio value ] [ skbqueue value ]
//
//      DEL-ENTRY := netaddr,[proto:]port,netaddr
//
//      TEST-ENTRY := netaddr,[proto:]port,netaddr
//
//      where netaddr := ip[/cidr]
//
// For the [proto:]port part of the elements see the
// description at the hash:ip,port set type. For the netaddr
// part of the elements see the description at the HashNet
// set type.
//
// From the set netfilter match point of view the searching
// for a match always starts from the smallest size of
// netblock (most specific cidr) to the largest one (least
// specific cidr) added to the set. When adding/deleting
// triples to the set by the SET netfilter target, it will be
// added/deleted by the most specific cidr which can be found
// in the set, or by the host cidr value if the set is empty.
// The first subnet has precedence when performing the
// most-specific lookup, just as for HashNetNet.
//
// The lookup time grows linearly with the number of the
// different cidr values added to the set and by the number of
// secondary cidr values per primary.
//
// The HashNetPortNet type of sets require three src/dst
// parameters of the set match and SET target kernel modules.
//
// Examples:
//
//      ipset create foo hash:net,port,net
//
//      ipset add foo 192.168.1.0/24,0,10.0.0/24
//
//      ipset add foo 192.168.2.0/24,25,10.1.0.0/16
//
//      ipset test foo 192.168.1.1,80,10.0.0.1
const HashNetPortNet SetType = "hash:net,port,net"

// HashNetIface set type uses a hash to store different sized
// IP network address and interface action pairs.
//
//      CREATE-OPTIONS := [ family { inet | inet6 } ] |
//                        [ hashsize value ] [ maxelem value ]
//                        [ timeout value ] [ counters ]
//                        [ comment ] [ skbinfo ]
//
//      ADD-ENTRY := netaddr,[physdev:]iface
//
//      ADD-OPTIONS := [ timeout value ] [ nomatch ]
//                     [ packets value ] [ bytes value ]
//                     [ comment string ] [ skbmark value ]
//                     [ skbprio value ] [ skbqueue value ]
//
//      DEL-ENTRY := netaddr,[physdev:]iface
//
//      TEST-ENTRY := netaddr,[physdev:]iface
//
//      where netaddr := ip[/cidr]
//
// For the netaddr part of the elements see the description at
// the HashNet set type.
//
// When adding/deleting/testing entries, if the cidr prefix
// parameter is not specified, then the host prefix value is
// assumed. When adding/deleting entries, the exact element is
// added/deleted and overlapping elements are not checked by
// the kernel. When testing entries, if a host address is
// tested, then the kernel tries to match the host address in
// the networks added to the set and reports the result
// accordingly.
//
// From the set netfilter match point of view the searching
// for a match always starts from the smallest size of
// netblock (most specific prefix) to the largest one (least
// specific prefix) added to the set. When adding/deleting IP
// addresses to the set by the SET netfilter target, it will
// be added/deleted by the most specific prefix which can be
// found in the set, or by the host prefix value if the set is
// empty.
//
// The second direction parameter of the set match and SET
// target modules corresponds to the incoming/outgoing
// interface: src to the incoming one (similar to the -i flag
// of iptables), while dst to the outgoing one (similar to the
// -o flag of iptables). When the interface is flagged with
// physdev:, the interface is interpreted as the incoming
// or outgoing bridge port.
//
// The lookup time grows linearly with the number of the
// different prefix values added to the set.
//
// The internal restriction of the hash:net,iface set type is
// that the same network prefix cannot be stored with more
// than 64 different interfaces in a single set.
//
// Examples:
//
//      ipset create foo hash:net,iface
//
//      ipset add foo 192.168.0/24,eth0
//
//      ipset add foo 10.1.0.0/16,eth1
//
//      ipset test foo 192.168.0/24,eth0
const HashNetIface SetType = "hash:net,iface"

// ListSet type uses a simple list in which you can store set
// names.
//
//      CREATE-OPTIONS := [ size value ] [ timeout value ]
//                        [ counters ] [ comment ]
//                        [ skbinfo ]
//
//      ADD-ENTRY := setname [ { before | after } setname ]
//
//      ADD-OPTIONS := [ timeout value ] [ packets value ]
//                     [ bytes value ] [ comment string ]
//                     [ skbmark value ] [ skbprio value ]
//                     [ skbqueue value ]
//
//      DEL-ENTRY := setname [ { before | after } setname ]
//
//      TEST-ENTRY := setname [ { before | after } setname ]
//
// Optional create options:
//
//       size value
//
// The size of the list, the default is 8.
//
// By the ipset command you can add, delete and test set names
// in a ListSet type of set.
//
// By the set match or SET target of netfilter you can test,
// add or delete entries in the sets added to the ListSet type
// of set. The match will try to find a matching entry in the
// sets and the target will try to add an entry to the first
// set to which it can be added. The number of direction
// options of the match and target are important: sets which
// require more parameters than specified are skipped, while
// sets with equal or less parameters are checked, elements
// added/deleted. For example if a and b are ListSet type of
// sets then in the command
//
//      iptables -m set --match-set a src,dst -j SET
//                      --add-set b src,dst
//
// the match and target will skip any set in a and b which
// stores data triples, but will match all sets with single or
// double data storage in a set and stop matching at the
// first successful set, and add src to the first single or
// src,dst to the first double data storage set in b to which
// the entry can be added. You can imagine a ListSet type of
// set as an ordered union of the set elements.
//
// Please note: by the ipset command you can add, delete and
// test the setnames in a ListSet type of set, and not the
// presence of a set's member (such as an IP address).
//
// Examples:
//      ipset create foo hash:ip
//
//      ipset create bar hash:ip
//
//      ipset create baz hash:ip
//
//      ipset create list list:set size 6
//
//      ipset add list foo
//
//      ipset add list bar after foo
//
//      ipset add list baz before bar
const ListSet SetType = "list:set"
