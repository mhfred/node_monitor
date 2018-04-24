package api

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func initServer() {
	initConfigApi()
}

func Serve() {
	initServer()
	r := mux.NewRouter()
	r.HandleFunc("/chainInfo", chainInfoHandler)
	r.HandleFunc("/bios/{name}/eosiosg.tar", generateZipHandler)
	r.HandleFunc("/nodes", nodesHandler)
	r.HandleFunc("/secureNode", secureNodesHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
