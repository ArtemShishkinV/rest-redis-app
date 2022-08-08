package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"rest-redis-app/internal/app/apiserver/http/dto"
	"rest-redis-app/internal/app/store"
	"rest-redis-app/utils"
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

	response := make(map[string]interface{})
	response["result"] = generateStringSendToApp(requestDto)

	utils.Respond(w, response)
}

func generateStringSendToApp(m *dto.MultiplicationRequestDto) string {
	EOF := "\r\n"
	result := ""

	for _, item := range m.Multipliers {
		result += strings.Join([]string{item.A, item.B}, ",") + EOF
	}

	return result + EOF
}
