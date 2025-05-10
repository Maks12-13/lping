package main

import (
	"fmt"
	"net"
	"os"
)

var DEBUG = false
var VERSION = "1.0.1"
var NAME = "lping"
var countInt int = 3
var waitTime int = 2

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(0) // Exit with code 0 if help is requested
	}

	switch os.Args[1] {
	case "-h", "--help":
		help()
		os.Exit(0)
	case "-v", "--version":
		version()
		os.Exit(0)
	}

	for i := 1; i < len(os.Args)-1; i++ {
		switch os.Args[i] {
		case "-d", "--debug":
			DEBUG = true
		case "-c":
			if i+1 < len(os.Args)-1 {
				// Parse the count value
				count := os.Args[i+1]
				if _, err := fmt.Sscanf(count, "%d", &countInt); err != nil || countInt < 1 || countInt > 999 {
					fmt.Println("Error: Invalid value for -c. Expected a number between 1 and 999.")
					os.Exit(1)
				}
				if DEBUG {
					fmt.Printf("Count: %s\n", count) // Replace with actual handling logic
				}
				i++ // Skip the next argument as it is part of this flag
			} else {
				fmt.Println("Error: Missing value for -c")
				os.Exit(1)
			}
		case "-W":
			if i+1 < len(os.Args)-1 {
				// Parse the wait time value
				wait := os.Args[i+1]
				if _, err := fmt.Sscanf(wait, "%d", &waitTime); err != nil || waitTime < 1 || waitTime > 30 {
					fmt.Println("Error: Invalid value for -W. Expected a number between 1 and 30.")
					os.Exit(1)
				}
				if DEBUG {
					fmt.Printf("Wait time: %s\n", wait) // Replace with actual handling logic
				}

				i++ // Skip the next argument as it is part of this flag
			} else {
				fmt.Println("Error: Missing value for -W")
				os.Exit(1)
			}
		default:
			fmt.Printf("Unknown option: %s\n", os.Args[i])
			os.Exit(1)
		}
	}

	// The last argument should be the IP address
	if len(os.Args) < 2 {
		fmt.Println("Error: Missing or IP address")
		os.Exit(1)
	}

	ip := os.Args[len(os.Args)-1]
	if net.ParseIP(ip) != nil {
		if DEBUG {
			fmt.Printf("DEBUG: IP address is  %s valid\n", ip)
		}
		if lping(ip, countInt, waitTime) {
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
