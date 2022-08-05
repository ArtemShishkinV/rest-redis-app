package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"rest-redis-app/internal/app/apiserver/http/dto"
	"rest-redis-app/internal/app/store"
	"rest-redis-app/utils"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting api server...")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/inc", s.handleInc)
	s.router.HandleFunc("/dec", s.handleDec)
	s.router.HandleFunc("/test1", s.handleIncrementKey)
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIServer) handleIncrementKey(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	key := &dto.IncrementKeyRequestDto{}

	err := json.NewDecoder(r.Body).Decode(key)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	s.logger.Info("in the handler, " + path)

	val, _ := s.store.IncrementKeyByValue(key.Key, key.Val)

	//response := make(map[string]int)
	//response[key.Key] = val

	//marshal, err := json.Marshal(response)
	//if err != nil {
	//	utils.Respond(w, utils.Message(false, "Invalid response"))
	//	return
	//}
	fmt.Printf("%s = %d", key.Key, val)
}

func (s *APIServer) handleInc(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	s.logger.Info("in the handler, " + path)

	val, _ := s.store.IncrementKeyByValue("key", 3)

	fmt.Fprintf(w, "%s = %d", "key", val)
}

func (s *APIServer) handleDec(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	s.logger.Info("in the handler, " + path)

	val, _ := s.store.Dec("key")

	fmt.Fprintf(w, "%s = %s", "key", val)
}
