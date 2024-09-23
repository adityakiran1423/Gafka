package main

import (
	"fmt"
	"net"
	"os"
	"encoding/binary" // added externally
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
	// fmt.Printf("message_length is %x\n", message_length)
	correlation_id := append(message_length, buff[8:12]...)



	request_api_version := binary.BigEndian.Uint16(buff[6:8])


	if request_api_version > 0 && request_api_version <= 4{
		kafka_response := append(correlation_id, buff[6:8]...)
		conn.Write(kafka_response)
	}
	api_error_code := make([]byte, 8)
	copy(api_error_code, []byte{'0','0','1','0','0','0','1','1'})
	kafka_response :=append(correlation_id, api_error_code...)
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

// func parseApiVersion(message []byte) []byte {

// 	decimal_num, err := strconv.ParseInt(message, 16, 64)

// 	// in case of any error
// 	if err != nil {
// 		panic(err)
// 	}

// 	// print the decimal number
// 	fmt.Println("The decimal conversion of", hexadecimal_num, "is ", decimal_num)

// 	// func parseApiVersion(message []byte) int{
// 	fmt.Printf("api from the request message in hex is %x\n", message)
// 	// fmt.Printf("second hex's decimal value is %d\n", message[1])
// 	api_version := message[0] + message[1]
// 	fmt.Printf("api_version which I calculated %x\n", api_version)
// 	// fmt.Println(api_version)
// 	if api_version > 0 && api_version <= 4 {
// 		return message
// 	}
// 	error_api_return_message := make([]byte, 4)
// 	copy(error_api_return_message, []byte{'0','1','0','1','0','0','1','1'})
// 	fmt.Printf("The api return if api version isn't 0 to 4 is %x\n", error_api_return_message)
// 	return error_api_return_message
// }
