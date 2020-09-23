package net

import (
	"net"
)

// LocalIP : Get preferred outbound ip of this machine
func LocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if warnOnErr("%v", err) != nil {
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return fSf("%v", localAddr.IP)
}
