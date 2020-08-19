package geoip

import (
	"fmt"
	"math/big"
	"net"
)

// AddressRange returns the first and last addresses in the given CIDR range.
func AddressRange(network *net.IPNet) (net.IP, net.IP) {
	// the first IP is easy
	firstIP := network.IP

	// the last IP is the network address OR NOT the mask address
	prefixLen, bits := network.Mask.Size()
	if prefixLen == bits {
		// Easy!
		// But make sure that our two slices are distinct, since they
		// would be in all other cases.
		lastIP := make([]byte, len(firstIP))
		copy(lastIP, firstIP)
		return firstIP, lastIP
	}

	firstIPInt, bits := IPToInt(firstIP)
	hostLen := uint(bits) - uint(prefixLen)
	lastIPInt := big.NewInt(1)
	lastIPInt.Lsh(lastIPInt, hostLen)
	lastIPInt.Sub(lastIPInt, big.NewInt(1))
	lastIPInt.Or(lastIPInt, firstIPInt)

	return firstIP, IntToIP(lastIPInt, bits)
}

// AddressCount returns the number of distinct host addresses within the given
// CIDR range.
//
// Since the result is a uint64, this function returns meaningful information
// only for IPv4 ranges and IPv6 ranges with a prefix size of at least 65.
func AddressCount(network *net.IPNet) uint64 {
	prefixLen, bits := network.Mask.Size()
	return 1 << (uint64(bits) - uint64(prefixLen))
}

// IPToInt todo
func IPToInt(ip net.IP) (*big.Int, int) {
	val := &big.Int{}
	val.SetBytes([]byte(ip))

	switch len(ip) {
	case net.IPv4len:
		return val, 32
	case net.IPv6len:
		return val, 128
	default:
		panic(fmt.Errorf("Unsupported address length %d", len(ip)))
	}
}

// IntToIP todo
func IntToIP(ipInt *big.Int, bits int) net.IP {
	ipBytes := ipInt.Bytes()
	ret := make([]byte, bits/8)
	// Pack our IP bytes into the end of the return array,
	// since big.Int.Bytes() removes front zero padding.
	for i := 1; i <= len(ipBytes); i++ {
		ret[len(ret)-i] = ipBytes[len(ipBytes)-i]
	}
	return net.IP(ret)
}
