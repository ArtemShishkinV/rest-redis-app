package apiserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"rest-redis-app/internal/app/store"
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
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIServer) handleInc(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	s.logger.Info("in the handler, " + path)

	val, _ := s.store.Inc("key")

	fmt.Fprintf(w, "%s = %s", "key", val)
}

func (s *APIServer) handleDec(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	s.logger.Info("in the handler, " + path)

	val, _ := s.store.Dec("key")

	fmt.Fprintf(w, "%s = %s", "key", val)
}
