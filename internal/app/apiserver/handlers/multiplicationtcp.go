package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net"
	"net/http"
	dto3 "rest-redis-app/internal/app/apiserver/dto"
	"rest-redis-app/pkg"
	"strings"
)

const TcpAddr = "server:4545"

type multiplicationTcpHandler struct{}

func NewMultiplicationTcpHandler() Handler {
	return &multiplicationTcpHandler{}
}

func (h *multiplicationTcpHandler) Register(router *mux.Router) {
	router.HandleFunc(MultiplicationTcpPath, h.handleMultiplication)
}

func (h *multiplicationTcpHandler) handleMultiplication(w http.ResponseWriter, r *http.Request) {
	requestDto := &dto3.MultiplicationRequestDto{}

	if err := readDataFromRequest(r, requestDto); err != nil {
		pkg.Respond(w, pkg.Message(false, "Invalid request"))
		return
	}

	response := getResultMultiplication(requestDto)

	if len(response) == 0 {
		pkg.Respond(w, pkg.Message(false, "Invalid request"))
		return
	}

	pkg.Respond(w, response)
}

func readDataFromRequest(r *http.Request, dto *dto3.MultiplicationRequestDto) error {
	bytes, err := io.ReadAll(r.Body)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(bytes, &dto.Multipliers); err != nil {
		return err
	}

	return nil
}

func getResultMultiplication(data *dto3.MultiplicationRequestDto) map[string]interface{} {
	requestString := generateStringToSendRequest(data)
	tcpResponse := sendTcpRequest(TcpAddr, requestString)

	return getResultFromStringTcpResponse(tcpResponse, data)
}

func sendTcpRequest(address string, message string) string {
	network := "tcp"

	conn, err := net.Dial(network, address)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer conn.Close()

	if n, err := conn.Write([]byte(message)); n == 0 || err != nil {
		fmt.Println(err)
		return ""
	}

	buff := make([]byte, 1024*4)
	n, err := conn.Read(buff)

	return string(buff[0:n])
}

func generateStringToSendRequest(m *dto3.MultiplicationRequestDto) string {
	EOL := "\r\n"
	result := ""

	for _, item := range m.Multipliers {
		result += strings.Join([]string{item.A, item.B}, ",") + EOL
	}

	return result + EOL
}

func getResultFromStringTcpResponse(data string, m *dto3.MultiplicationRequestDto) map[string]interface{} {
	response := make(map[string]interface{})
	results, err := pkg.GetNumbersFromString(data)

	fmt.Println(results)

	if err != nil || len(results) == 0 {
		return response
	}

	for index, item := range results {
		response[m.Multipliers[index].Key] = item
	}

	return response
}
