package main

import (
	"encoding/binary" // added externally
	"fmt"
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:9092") // creates a server
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

	// var buff = make([]byte, 256)
	var buff = make([]byte, 1024)
	conn.Read(buff)

	fmt.Printf("recieved buff is %x and in string form -> %s", buff,buff)

	message_length := make([]byte, 4)
	copy(message_length, []byte{'0', '0', '0', '0'})
	// fmt.Printf("message_length is %x\n", message_length)
	correlation_id := append(message_length, buff[8:12]...)

	request_api_version := binary.BigEndian.Uint16(buff[6:8])

	if request_api_version > 0 && request_api_version <= 4 {
		kafka_response := append(correlation_id, buff[6:8]...)
		fmt.Println("in if")
		fmt.Printf("kafka response is %x", kafka_response)
		conn.Write(kafka_response)
	}
	api_error_code := make([]byte, 2)
	//copy()
	copy(api_error_code, []byte{0x00,0x23})
	kafka_response := append(correlation_id, api_error_code...)
	fmt.Println("in else ffs")
	fmt.Printf("kafka response is %x", kafka_response)
	conn.Write(kafka_response)
	// byte_api_version := make([]byte, 2)
	// copy(byte_api_version, buff[6:8])
	// fmt.Println(correlation_id)
	// fmt.Printf("hex val of api from request %x\n", byte_api_version)
	// api_version := parseApiVersion(request_api_version)
	// fmt.Printf("api_version recieved from func is %s\n", api_version)
	// fmt.Printf("api_version recieved from func in hex is %x\n", api_version)
	// int_api_verion := binary.BigEndian.Uint64(byte_api_version)

	// fmt.Printf("eventually this is sent (string format) %s \n", kafka_response)
	// fmt.Printf("eventually this is sent (byte format) %b \n", kafka_response)
	// fmt.Printf("eventually this is sent (hex format) %x \n", kafka_response)
	// conn.Write(kafka_response)
	// conn.Write(api_version)
}
