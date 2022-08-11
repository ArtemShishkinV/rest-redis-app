package handlers

import "github.com/gorilla/mux"

const (
	IncKeyPath            = "/test1"
	MakeSignPath          = "/test2"
	MultiplicationTcpPath = "/test3"
)

type Handler interface {
	Register(router *mux.Router)
}
