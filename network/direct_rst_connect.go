package main

import (
	"net"
	"syscall"
)

func NewRstDialer() *net.Dialer {
	return &net.Dialer{
		Control: func(network, address string, c syscall.RawConn) error {
			return c.Control(func(fd uintptr) {
				syscall.SetsockoptLinger(int(fd), syscall.SOL_SOCKET, syscall.SO_LINGER, &syscall.Linger{Onoff: 1, Linger: 0})
			})
		},
	}
}

func main() {

	rstDialer := NewRstDialer()
	rstConn, _ := rstDialer.Dial("tcp", "127.0.0.1:8080")
	_ = rstConn.Close()

	finDialer := net.Dialer{}
	finConn, _ := finDialer.Dial("tcp", "127.0.0.1:8080")
	_ = finConn.Close()

}
