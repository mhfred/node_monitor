package model

import "encoding/json"

type PeerInfo struct {
	Peer          string        `json:"peer"`
	Connecting    bool          `json:"connecting"`
	Syncing       bool          `json:"syncing"`
	LastHandshake LastHandshake `json:"last_handshake"`
}

type LastHandshake struct {
	NetworkVersion           int64  `json:"network_version"`
	ChainId                  string `json:"chain_id"`
	NodeId                   string `json:"node_id"`
	Key                      string `json:"key"`
	Time                     json.Number `json:"time"`
	Token                    string `json:"token"`
	Sig                      string `json:"sig"`
	P2pAddress               string `json:"p2p_address"`
	LastIrreversibleBlockNum int64  `json:"last_irreversible_block_num"`
	LastIrreversibleBlockId  string `json:"last_irreversible_block_id"`
	HeadNum                  int64  `json:"head_num"`
	HeadId                   string `json:"head_id"`
	Os                       string `json:"os"`
	Agent                    string `json:"agent"`
	Generation               int64  `json:"generation"`
}

func (peer *PeerInfo) String() string{
	data,err:= json.Marshal(peer)
	if err!=nil{
		return ""
	}
	return string(data)
}

func PeerInfoFromString(data string) *PeerInfo {
	var peer PeerInfo
	json.Unmarshal([]byte(data),&peer)
	return &peer
}
