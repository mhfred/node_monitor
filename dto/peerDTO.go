package dto

type PeerDTO struct {
	NodeId     string `json:"node_id"`
	Agent      string `json:"agent"`
	P2pAddress string `json:"p2p_address"`
	Os         string `json:"os"`
	Location   string `json:"location"`
	Status     string `json:"status"`
}