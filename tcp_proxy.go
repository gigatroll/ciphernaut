package ciphernaut

import (
	"fmt"
	"net"
	"strconv"
)

func StartTCPProxy(ip string, port int) (ok bool, err error) {

	laddr := net.TCPAddr{IP: net.ParseIP(ip), Port: port}

	ln, err := net.ListenTCP("tcp4", &laddr)
	if err != nil {
		return false, err
	}
	go tcpProxyAcceptLoop(ln)
	return true, nil
}

func tcpProxyAcceptLoop(ln *net.TCPListener) {
	var conn *net.TCPConn
	var err error

	for {
		conn, err = ln.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleTcpProxyClient(conn)
	}
}

func handleTcpProxyClient(conn *net.TCPConn) {
	defer conn.Close()
	fmt.Println("Incomming client")
	conn.Write([]byte("Hello!\n"))
	clientAddr := conn.RemoteAddr()

	socket, err := conn.File()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer socket.Close()

	dest, port, err := GetOriginalDestination(4, socket)
	if err != nil {
		fmt.Println(err)
		return
	}

	fromTo := clientAddr.String() + " => " + dest + ":" + strconv.Itoa(port)
	conn.Write([]byte(fromTo))
}
