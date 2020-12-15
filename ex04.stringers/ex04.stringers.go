/*
Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.
For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4"
*/
package main

import "fmt"

//IPAddr interfase to 4 byte array that holds ip adress and could be handeled as a string
type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("\"%d.%d.%d.%d\"", ip[0], ip[1], ip[2], ip[3])
}

// TODO: Add a "String() string" method to IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
