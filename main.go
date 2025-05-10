package main

import (
	"fmt"
	"net"
	"os"
)

var DEBUG = false
var VERSION = "1.0.0"
var NAME = "lping"

func main() {
	if len(os.Args) < 2 || (os.Args[1] == "-h") || (os.Args[1] == "--help") {
		help()
		os.Exit(0) // Exit with code 0 if help is requested
	}
	if (os.Args[1] == "-v") || (os.Args[1] == "--version") {
		version()
		os.Exit(0) // Exit with code 0 if version is requested
	}
	if (os.Args[1] == "-d") || (os.Args[1] == "--debug") {
		os.Args = append(os.Args[:1], os.Args[2:]...) // Remove the first argument (-d)
		DEBUG = true
		//fmt.Println("DEBUG mode is enabled")
		//os.Exit(0) // Exit with code 0 if debug mode is enabled
	}

	if len(os.Args) != 2 {
		help()
		os.Exit(1) // Exit with code 1 if no argument is provided
	}

	ip := os.Args[1]
	if net.ParseIP(ip) != nil {
		if DEBUG {
			fmt.Printf("DEBUG: IP address is  %s valid\n", ip)
		}
		if lping(ip) {
			os.Exit(0) // Exit with code 0 if gping returns true
		} else {
			os.Exit(1) // Exit with code 1 if gping returns false
		}
		os.Exit(0) // Exit with code 0 if the IP is valid
	} else {
		if DEBUG {
			println("DEBUG: IP address is not valid")
		}
		os.Exit(1) // Exit with code 1 if the IP is invalid
	}
}
