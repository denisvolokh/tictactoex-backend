package handlers

import (
	"go-tic-tac-toe-api/store"
	"net/http"
)

type IEventhandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	store store.IGameStore
}

func NewEventHandler(store store.IGameStore) IEventhandler {
	return &handler{store: store}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
