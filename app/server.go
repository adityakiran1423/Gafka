package main

import (
	"fmt"
	"net"
	"os"
	// "encoding/binary" // added externally
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	var buff = make([]byte, 256)

	conn.Read(buff)

	message_length := make([]byte, 4)
	copy(message_length, []byte{'0', '0', '0', '0'})
	correlation_id := append(message_length, buff[8:12]...)

	byte_api_version := make([]byte, 2)
	copy(byte_api_version, buff[6:8])

	parseApiVersion(byte_api_version)

	// int_api_verion := binary.BigEndian.Uint64(byte_api_version)

	conn.Write(correlation_id)
}


func parseApiVersion(message []byte) []byte{
	fmt.Println(message)
	return message
}