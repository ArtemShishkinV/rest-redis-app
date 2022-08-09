package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"net/http"
	"regexp"
	"rest-redis-app/internal/app/apiserver/http/dto"
	"rest-redis-app/internal/app/store"
	"rest-redis-app/utils"
	"strconv"
	"strings"
)

type server struct {
	logger *logrus.Logger
	router *mux.Router
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureLogger() error {
	level, err := logrus.ParseLevel("debug")
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/test1", s.handleIncrementKey)
	s.router.HandleFunc("/test2", s.handleMakeSign)
	s.router.HandleFunc("/test3", s.handleMultiplication)
}

func (s *server) handleIncrementKey(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	s.logger.Info("in the handler, " + path)

	requestDto := &dto.IncrementKeyRequestDto{}

	err := json.NewDecoder(r.Body).Decode(requestDto)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	val, _ := s.store.Repository().IncrementKeyByValue(requestDto.Key, requestDto.Val)

	response := make(map[string]interface{})
	response[requestDto.Key] = val

	utils.Respond(w, response)
}

func (s *server) handleMakeSign(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	s.logger.Info("in the handler, " + path)

	requestDto := &dto.ComputeHmacDto{}
	err := json.NewDecoder(r.Body).Decode(requestDto)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	hmac512 := utils.ComputeHmac512(requestDto.S, requestDto.Key)

	response := make(map[string]interface{})
	response["hmac512"] = hmac512

	utils.Respond(w, response)
}

func (s *server) handleMultiplication(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	s.logger.Info("in the handler, " + path)

	requestDto := &dto.MultiplicationRequestDto{}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	err = json.Unmarshal(bytes, &requestDto.Multipliers)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	responseString := getResultMultiplication(generateStringToSendRequest(requestDto))

	response := getResultFromStringResponseTcp(responseString, requestDto)

	utils.Respond(w, response)
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

func generateStringToSendRequest(m *dto.MultiplicationRequestDto) string {
	EOF := "\r\n"
	result := ""

	for _, item := range m.Multipliers {
		result += strings.Join([]string{item.A, item.B}, ",") + EOF
	}

	return result + EOF
}

func getResultFromStringResponseTcp(data string, m *dto.MultiplicationRequestDto) map[string]interface{} {
	response := make(map[string]interface{})
	results := getNumbersFromRequest(data)

	for index, item := range results {
		response[m.Multipliers[index].Key] = item
	}

	return response
}

func getNumbersFromRequest(data string) []int {
	var numbers []int

	re := regexp.MustCompile("-?\\d+")
	matchNumbers := re.FindAllString(data, -1)

	for _, element := range matchNumbers {
		i, _ := strconv.Atoi(element)
		numbers = append(numbers, i)
	}

	return numbers
}
