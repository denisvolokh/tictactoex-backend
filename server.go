package main

import (
	"go-tic-tac-toe-api/handlers"
	"go-tic-tac-toe-api/store"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Args struct {
	conn string
	port string
}

func Run(args Args) error {
	router := mux.NewRouter().
		PathPrefix("/api/v1/").
		Subrouter()

	st := store.NewSqliteGameStore(args.conn)
	hnd := handlers.NewEventHandler(st)
	RegisterAllRouters(router, hnd)

	log.Println("Starting server at port: ", args.port)
	return http.ListenAndServe(args.port, router)
}

func middleware(next http.Handler) http.Handler {
	return &middlewareHandler{next}
}

type middlewareHandler struct {
	next http.Handler
}

func (m *middlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	m.next.ServeHTTP(w, r)
}

func RegisterAllRouters(router *mux.Router, hnd handlers.IEventhandler) {

	// router.Use(func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		w.Header().Set("Content-Type", "application/json")
	// 		next.ServeHTTP(w, r)
	// 	})
	// })

	// Set Middleware
	router.Use(middleware)

	router.HandleFunc("/game", hnd.List).Methods(http.MethodGet)
	router.HandleFunc("/games", hnd.Get).Methods(http.MethodGet)
	router.HandleFunc("/game", hnd.Create).Methods(http.MethodPost)
	router.HandleFunc("/game", hnd.Update).Methods(http.MethodPut)
}
