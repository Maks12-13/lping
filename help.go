package main

import (
	"fmt"
)

func help() {
	fmt.Println("Usage: lping [options] <ip>")
	fmt.Println("lping is a lite ping utility.")
	fmt.Println("<ip>          IP address to ping")
	fmt.Println("Options:")
	fmt.Println("  -h, --help    Show this help message")
	fmt.Println("  -v, --version Show version information")
	fmt.Println("  -d, --debug   Enable debug mode")
	fmt.Println(" reply from <ip> is OK - exit code 0")
	fmt.Println(" reply from <ip> is not OK - exit code 1")

}

func version() {
	fmt.Printf("%s version %s\n", NAME, VERSION)
	fmt.Println("lping is a lite ping utility.")
	fmt.Println("Author: MakselPr")
	fmt.Println("License: MIT")
}
