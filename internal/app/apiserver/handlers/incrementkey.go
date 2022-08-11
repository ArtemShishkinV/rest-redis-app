package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rest-redis-app/internal/app/apiserver/dto"
	"rest-redis-app/internal/app/store"
	"rest-redis-app/pkg"
)

type incrementKeyHandler struct {
	repository store.Repository
}

func NewIncKeyHandler(repository store.Repository) Handler {
	return &incrementKeyHandler{
		repository: repository,
	}
}

func (h *incrementKeyHandler) Register(router *mux.Router) {
	router.HandleFunc(IncKeyPath, h.handleIncrementKey)
}

func (h *incrementKeyHandler) handleIncrementKey(w http.ResponseWriter, r *http.Request) {
	requestDto := &dto.IncrementKeyRequestDto{}

	err := json.NewDecoder(r.Body).Decode(requestDto)
	if err != nil {
		pkg.Respond(w, pkg.Message(false, "Invalid request"))
		return
	}

	val, _ := h.repository.IncrementKeyByValue(requestDto.Key, requestDto.Val)

	response := make(map[string]interface{})
	response[requestDto.Key] = val

	pkg.Respond(w, response)
}