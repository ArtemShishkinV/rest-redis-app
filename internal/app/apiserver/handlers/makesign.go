package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	dto3 "rest-redis-app/internal/app/apiserver/dto"
	"rest-redis-app/pkg"
)

type makeSignHandler struct{}

func NewMakeSignHandler() Handler {
	return &makeSignHandler{}
}

func (h *makeSignHandler) Register(router *mux.Router) {
	router.HandleFunc(MakeSignPath, h.handleMakeSign)
}

func (h *makeSignHandler) handleMakeSign(w http.ResponseWriter, r *http.Request) {
	requestDto := &dto3.ComputeHmacDto{}

	if err := json.NewDecoder(r.Body).Decode(requestDto); err != nil {
		pkg.Respond(w, pkg.Message(false, "Invalid request"))
		return
	}

	hmac512 := pkg.ComputeHmac512(requestDto.S, requestDto.Key)

	response := make(map[string]interface{})
	response["hmac512"] = hmac512

	pkg.Respond(w, response)
}
