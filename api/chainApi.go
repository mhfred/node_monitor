package api

import (
	"net/http"
	"node_monitor/persist"
)

func chainInfoHandler(w http.ResponseWriter, r *http.Request){
	chainInfo := persist.GetChainInfo()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(chainInfo))
}