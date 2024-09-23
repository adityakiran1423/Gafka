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
	fmt.Printf("the message_length byte array is %b\n", message_length)
	correlation_id := append(message_length, buff[8:12]...)

	byte_api_version := make([]byte, 2)
	copy(byte_api_version, buff[6:8])
	fmt.Println(correlation_id)
	api_version := parseApiVersion(byte_api_version)
	fmt.Printf("api_version recieved from func is %s", api_version)
	// int_api_verion := binary.BigEndian.Uint64(byte_api_version)

	conn.Write(correlation_id)
	conn.Write(api_version)
}

func parseApiVersion(message []byte) []byte {
	// func parseApiVersion(message []byte) int{
	fmt.Printf("first hex's decimal value is %d\n", message[0])
	fmt.Printf("second hex's decimal value is %d\n", message[1])
	api_version := message[0] + message[1]
	fmt.Println(api_version)
	if api_version > 0 && api_version <= 4 {
		return message
	}
	error_api_return_message := make([]byte, 4)
	copy(error_api_return_message, []byte{'0','1','0','1','0','0','1','1'})
	fmt.Printf("The api return if api version isn't 0 to 4 is %d", error_api_return_message)
	return error_api_return_message
}
