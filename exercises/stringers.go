package exercises

import "fmt"

//Exercise: Stringers

// Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.
// For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".

// TODO: Add a "String() string" method to IPAddr.

type IPAddr [4]byte

func (p IPAddr) String() string {
	var ipString string

	for i, octet := range p {
		if i == 0 {
			ipString = fmt.Sprintf("%d", octet)
		} else {
			ipString = fmt.Sprintf("%s.%d", ipString, octet)
		}
	}
	return ipString
}

func MakeIPAddr() {
	hosts := map[string]IPAddr{
		"loopback":      {127, 0, 0, 1},
		"googleDNS":     {8, 8, 8, 8},
		"someRandomDNS": {0, 0, 0, 11},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
