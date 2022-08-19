package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"tcpserver/pkg"
)

const EOL = "\r\n"

func main() {
	listener, err := net.Listen("tcp", "server:4545")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()
	fmt.Println("Server is listening...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			if err := conn.Close(); err != nil {
				return
			}
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	response, err := readRequest(conn)
	if err != nil {
		_, err = conn.Write([]byte("Invalid request data"))
	}

	response = getResponse(response)

	_, err = conn.Write([]byte(response + EOL))
}

func readRequest(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	result := ""

	for {
		source, err := reader.ReadString('\n')
		if source == EOL || err != nil && err != io.EOF {
			return result, err
		}
		result += source
	}
}

func getResponse(data string) string {
	multipliers, err := pkg.GetNumbersFromString(data)

	if len(multipliers)%2 != 0 || err != nil {
		return "Invalid request data"
	}

	var results []string

	product := 1

	for i := 0; i < len(multipliers); i += 2 {
		product = multipliers[i] * multipliers[i+1]
		results = append(results, strconv.Itoa(product))
	}

	return strings.Join(results, "\r\n") + EOL
}
