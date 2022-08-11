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

type multiplicationTcpHandler struct{}

func NewMultiplicationTcpHandler() Handler {
	return &multiplicationTcpHandler{}
}

func (h *multiplicationTcpHandler) Register(router *mux.Router) {
	router.HandleFunc(MultiplicationTcpPath, h.handleMultiplication)
}

func (h *multiplicationTcpHandler) handleMultiplication(w http.ResponseWriter, r *http.Request) {
	requestDto := &dto3.MultiplicationRequestDto{}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		pkg.Respond(w, pkg.Message(false, "Invalid request"))
		return
	}

	if err = json.Unmarshal(bytes, &requestDto.Multipliers); err != nil {
		pkg.Respond(w, pkg.Message(false, "Invalid request"))
		return
	}

	responseString := getResultMultiplication(generateStringToSendRequest(requestDto))

	response := getResultFromStringResponseTcp(responseString, requestDto)

	pkg.Respond(w, response)
}

func getResultMultiplication(data string) string {
	return sendTcpRequest("127.0.0.1:4545", data)
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
	EOF := "\r\n"
	result := ""

	for _, item := range m.Multipliers {
		result += strings.Join([]string{item.A, item.B}, ",") + EOF
	}

	return result + EOF
}

func getResultFromStringResponseTcp(data string, m *dto3.MultiplicationRequestDto) map[string]interface{} {
	response := make(map[string]interface{})
	results := pkg.GetNumbersFromString(data)

	for index, item := range results {
		response[m.Multipliers[index].Key] = item
	}

	return response
}
