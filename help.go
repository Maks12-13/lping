package main

import (
	"fmt"
)

func help() {
	fmt.Println(`Usage: lping [options] <ip>
lping is a lite ping utility.
<ip>          IP address to ping
Options:
  -h, --help    Show this help message
  -v, --version Show version information
  -d, --debug   Enable debug mode
  -W <timeout>  Time to wait for response
  -c <count>    Number of packets to send
 reply from <ip> is OK - exit code 0
 reply from <ip> is not OK - exit code 1`)

}

func version() {
	fmt.Printf("%s version %s\n", NAME, VERSION)
	fmt.Println("lping is a lite ping utility.")
	fmt.Println("Author: MakselPr")
	fmt.Println("License: MIT")
}
