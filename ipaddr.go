package main

import (
	"fmt"
	"github.com/tech10/go-ipify"
)

func main() {

	// Initialize variables.
	// We will be resolving to IPV4 and IPV6.
	var ip4 string
	var ip6 string
	var err error

	ip4, err = ipify.GetIp4()
	if err != nil {
		fmt.Println("IPV4 address retrieval failure:", err)
	} else {
		fmt.Println("IPV4 address:", ip4)
	}

	ip6, err = ipify.GetIp6()
	if err != nil {
		fmt.Println("IPV6 address retrieval failure:", err)
	} else {
		fmt.Println("IPV6 address:", ip6)
	}
fmt.Println("Press enter to continue.")
fmt.Scanln()
}
