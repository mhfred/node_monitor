package api

import (
	"node_monitor/dto"
	"text/template"
	"node_monitor/persist"
	"node_monitor/model"
	"bytes"
	"fmt"
	"net/http"
	"archive/tar"
	"github.com/gorilla/mux"
)

var tmpl *template.Template

func initConfigApi() {
	tmpl = template.Must(template.New("ConfigIniFile").Parse(dto.ConfigIniFile))
}

type configData struct {
	Domain        string
	P2pPort       string
	HttpPort      string
	ProducerName  string
	PublicKey     string
	PrivateKey    string
	PeerAddresses []string
}

func generateZipHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	tarWriter := tar.NewWriter(buf)

	d := generateConfigIniFile(name)
	dHeader := &tar.Header{Name: "eosiosg/config-dir/config.ini", Mode: 0666, Size: int64(len(d))}
	tarWriter.WriteHeader(dHeader)
	tarWriter.Write([]byte(d))

	d = generateGenesisJson()
	dHeader = &tar.Header{Name: "eosiosg/config-dir/genesis.json", Mode: 0666, Size: int64(len(d))}
	tarWriter.WriteHeader(dHeader)
	tarWriter.Write([]byte(d))

	d = generateKillShell()
	dHeader = &tar.Header{Name: "eosiosg/kill.sh", Mode: 0755, Size: int64(len(d))}
	tarWriter.WriteHeader(dHeader)
	tarWriter.Write([]byte(d))

	d = generateRunShell()
	dHeader = &tar.Header{Name: "eosiosg/run.sh", Mode: 0755, Size: int64(len(d))}
	tarWriter.WriteHeader(dHeader)
	tarWriter.Write([]byte(d))

	// Make sure to check the error on Close.
	err := tarWriter.Close()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Disposition", "attachment; filename=eosiosg.tar")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(buf.Bytes())
}

func generateConfigIniFile(name string) string {
	nodes := persist.GetNodes()
	found := false
	var node model.Node
	for _, n := range nodes {
		if n.ProducerName == name {
			found = true
			node = n
			break
		}
	}

	var data bytes.Buffer
	if found {
		peersAddresses := model.PeerAddresses(nodes)
		configData := configData{
			Domain:        node.Domain,
			P2pPort:       node.P2pPort,
			HttpPort:      node.HttpPort,
			ProducerName:  node.ProducerName,
			PublicKey:     node.PublicKey,
			PrivateKey:    node.PrivateKey,
			PeerAddresses: peersAddresses,
		}

		tmpl.Execute(&data, configData)
	}

	return data.String()
}

func generateGenesisJson() string {
	return dto.Genesis
}

func generateRunShell() string {
	return dto.Run
}

func generateKillShell() string {
	return dto.Kill
}
