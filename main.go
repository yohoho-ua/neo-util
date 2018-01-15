package main

import (
	//"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/neo", AccountInfoHandler).Methods("GET")
	r.HandleFunc("/neo", TransferHandler).Methods("POST")
	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}
