package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"node_monitor/dto"
	"node_monitor/persist"
	"sort"
	"node_monitor/monitor"
)

func nodesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST,OPTIONS")

	if r.Method == "OPTIONS"{
		w.WriteHeader(http.StatusOK)
	}else if r.Method == "GET" {
		getNodesHandler(w, r)
	} else if r.Method == "POST" {
		saveNodeHandler(w, r)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func getNodesHandler(w http.ResponseWriter, r *http.Request) {
	nodes := persist.GetNodes()
	nodeDTOs := dto.NodeDTOsFromNodes(nodes)

	sort.Slice(nodeDTOs, func(i, j int) bool {
		return nodeDTOs[i].ProducerName < nodeDTOs[j].ProducerName
	})

	data, _ := json.Marshal(nodeDTOs)
	w.Write(data)
}

func saveNodeHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading body", http.StatusInternalServerError)
	}
	var nodeDTO dto.NodeDTO
	err = json.Unmarshal(body, &nodeDTO)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "invalid json", http.StatusBadRequest)
	} else {
		nodes := persist.GetNodes()
		for _, node := range nodes {
			if nodeDTO.Domain == node.Domain && (nodeDTO.HttpPort == node.HttpPort || nodeDTO.P2pPort == node.P2pPort){
				e := fmt.Sprintf(`{"error_code":3,"error_message":"Duplicated domain"}`)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(e))
				return
			}
		}

		publicKey := r.URL.Query().Get("publicKey")
		privateKey := ""
		if publicKey == "" {
			privateKey, publicKey ,err= monitor.CreateKey()
		}

		if err !=nil{
			e := fmt.Sprintf(`{"error_code":2,"error_message":"Create key failed"}`)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(e))
			return
		}

		err = monitor.CreateAccount(nodeDTO.ProducerName,publicKey)
		if err !=nil {
			e := fmt.Sprintf(`{"error_code":1,"error_message":"Duplicated producer name"}`)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(e))
		}else{
			node := nodeDTO.ToNode()
			node.PublicKey = publicKey
			node.PrivateKey = privateKey
			persist.SaveNode(node)
			path:=fmt.Sprintf(`{"path":"/bios/%s/eosiosg.tar"}`,nodeDTO.ProducerName)
			w.Write([]byte(path))
		}
	}
}

func secureNodesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST,OPTIONS")

	if r.Method == "OPTIONS"{
		w.WriteHeader(http.StatusOK)
	}else if r.Method == "GET" {
		getNodesHandler(w, r)
	} else if r.Method == "POST" {
		saveSecureNodeHandler(w, r)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func saveSecureNodeHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading body", http.StatusInternalServerError)
	}
	var nodeDTO dto.NodeDTO
	err = json.Unmarshal(body, &nodeDTO)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "invalid json", http.StatusBadRequest)
	} else {
		nodes := persist.GetNodes()
		for _, node := range nodes {
			if nodeDTO.Domain == node.Domain && (nodeDTO.HttpPort == node.HttpPort || nodeDTO.P2pPort == node.P2pPort){
				e := fmt.Sprintf(``)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(e))
				return
			}
		}

		publicKey := r.URL.Query().Get("publicKey")
		privateKey := ""
		if publicKey == "" {
			privateKey, publicKey ,err= monitor.CreateKey()
		}

		if err !=nil{
			e := fmt.Sprintf(``)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(e))
			return
		}

		err = monitor.CreateAccount(nodeDTO.ProducerName,publicKey)
		if err !=nil {
			e := fmt.Sprintf(``)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(e))
		}else{
			node := nodeDTO.ToNode()
			node.PublicKey = publicKey
			node.PrivateKey = privateKey
			persist.SaveNode(node)
			path:=fmt.Sprintf(`/bios/%s/eosiosg.tar`,nodeDTO.ProducerName)
			w.Write([]byte(path))
		}
	}
}