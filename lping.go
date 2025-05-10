package main

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func lping(ip string, countInt int, waitTime int) bool {
	const (
		icmpEchoRequest = 8
		protocolICMP    = 1
	)

	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		fmt.Println("Error creating ICMP connection:", err)
		return false
	}
	defer conn.Close()

	dst, err := net.ResolveIPAddr("ip4", ip)
	if err != nil {
		fmt.Println("Error resolving IP address:", err)
		return false
	}

	for i := 0; i < countInt; i++ {
		msg := icmp.Message{
			Type: ipv4.ICMPTypeEcho,
			Code: 0,
			Body: &icmp.Echo{
				ID:   i + 1,
				Seq:  i + 1,
				Data: []byte("ping"),
			},
		}

		msgBytes, err := msg.Marshal(nil)
		if err != nil {
			fmt.Println("Error marshaling ICMP message:", err)
			return false
		}

		_, err = conn.WriteTo(msgBytes, dst)
		if err != nil {
			fmt.Println("Error sending ICMP request:", err)
			return false
		}
		conn.SetReadDeadline(time.Now().Add(time.Duration(waitTime) * time.Second))
		reply := make([]byte, 1500)
		n, _, err := conn.ReadFrom(reply)
		if err == nil {
			parsedMsg, err := icmp.ParseMessage(protocolICMP, reply[:n])
			if err != nil {
				if DEBUG {
					fmt.Println("Error parsing ICMP message:", err)
				}
				continue
			}

			// Проверяем, что это ICMP Echo Reply
			if parsedMsg.Type == ipv4.ICMPTypeEchoReply {
				if DEBUG {
					fmt.Printf("Received Echo Reply: %x\n", reply[:n])
				}
				return true
			} else if DEBUG {
				fmt.Printf("Received unexpected ICMP message: Type=%v, Code=%v\n", parsedMsg.Type, parsedMsg.Code)
			}
		} else if DEBUG {
			fmt.Println("No reply received or error:", err)
		}
	}

	return false
}
