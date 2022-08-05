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

//func (s *server) Start() error {
//	if err := s.configureLogger(); err != nil {
//		return err
//	}
//
//	s.configureRouter()
//
//	if err := s.configureStore(); err != nil {
//		return err
//	}
//
//	s.logger.Info("starting api server...")
//
//	return http.ListenAndServe(s.config.BindAddr, s.router)
//}

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

//func (s *server) configureStore() error {
//	st := store.newServer(s.config.Store)
//	if err := st.Open(); err != nil {
//		return err
//	}
//
//	s.store = st
//
//	return nil
//}

func (s *server) handleIncrementKey(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	key := &dto.IncrementKeyRequestDto{}

	err := json.NewDecoder(r.Body).Decode(key)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	s.logger.Info("in the handler, " + path)

	val, _ := s.store.Repository().IncrementKeyByValue(key.Key, key.Val)

	//response := make(map[string]int)
	//response[key.Key] = val

	//marshal, err := json.Marshal(response)
	//if err != nil {
	//	utils.Respond(w, utils.Message(false, "Invalid response"))
	//	return
	//}
	fmt.Printf("%s = %d", key.Key, val)
}

func (s *server) handleMakeSign(w http.ResponseWriter, r *http.Request) {

}

func (s *server) handleMultiplication(ц http.ResponseWriter, r *http.Request) {

}
