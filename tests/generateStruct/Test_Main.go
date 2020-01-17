package main

import (
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s host:port", os.Args[0])
	}
	//os.Args[0] = "127.0.0.1"
	//os.Args[1] = "127.0.0.1:8000"
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	n, err := conn.Write([]byte("hhhhhhh"))
	n, err = conn.Write([]byte("HEAD / HTTP/1.1\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	n, err =conn.Write([]byte("abcdefg"))
	log.Fatal(n)
}


//func main() {
//	if len(os.Args) != 2 {
//		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
//		os.Exit(1)
//	}
//	service := os.Args[1]
//	udpAddr, err := net.ResolveUDPAddr("udp4", service)
//	checkError(err)
//	conn, err := net.DialUDP("udp", nil, udpAddr)
//	checkError(err)
//	_, err = conn.Write([]byte("anything"))
//	checkError(err)
//	var buf [512]byte
//	n, err := conn.Read(buf[0:])
//	checkError(err)
//	fmt.Println(string(buf[0:n]))
//	os.Exit(0)
//}
//func checkError(err error) {
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
//		os.Exit(1)
//	}
//}


