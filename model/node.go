package model

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	PublicKey        string `json:"public_key"`
	PrivateKey       string `json:"private_key"`
	ProducerName     string `json:"producer_name"`
	Domain           string `json:"domain"`
	OrganizationName string `json:"organization_name"`
	Location         string `json:"location"`
	HttpPort         string `json:"http_port"`
	P2pPort          string `json:"p2p_port"`
	Status           int    `json:"status"`
	Logo             string `json:"logo"`
}

func (node *Node) ToJson() string {
	data, err := json.Marshal(node)
	if err != nil {
		return ""
	}
	return string(data)
}

func NodeFromJson(data string) Node {
	var node Node
	json.Unmarshal([]byte(data), &node)
	return node
}

func PeerAddresses(nodes []Node) []string {
	peers := make([]string, 0)
	for _, node := range nodes {
		if node.Status == 1 {
			addr := fmt.Sprintf("%s:%s", node.Domain, node.P2pPort)
			peers = append(peers, addr)
		}
	}
	return peers
}
