package dto

import (
	"node_monitor/model"
)

type NodeDTO struct {
	ProducerName     string `json:"producer_name"`
	Domain           string `json:"domain"`
	OrganizationName string `json:"organization_name"`
	Location         string `json:"location"`
	HttpPort         string `json:"http_port"`
	P2pPort          string `json:"p2p_port"`
	Status           int    `json:"status"`
	Logo             string `json:"logo"`
}

func NodeDTOFromNode(node model.Node) NodeDTO {
	return NodeDTO{
		ProducerName:     node.ProducerName,
		Domain:           node.Domain,
		OrganizationName: node.OrganizationName,
		Location:         node.Location,
		HttpPort:         node.HttpPort,
		P2pPort:          node.P2pPort,
		Status:           node.Status,
		Logo:             node.Logo,
	}
}

func NodeDTOsFromNodes(nodes []model.Node) []NodeDTO {
	nodeDTOs := make([]NodeDTO, 0)
	for _, node := range nodes {
		nodeDTOs = append(nodeDTOs, NodeDTOFromNode(node))
	}
	return nodeDTOs
}

func (nodeDTO NodeDTO) ToNode() model.Node {
	return model.Node{
		ProducerName:     nodeDTO.ProducerName,
		Domain:           nodeDTO.Domain,
		OrganizationName: nodeDTO.OrganizationName,
		Location:         nodeDTO.Location,
		HttpPort:         nodeDTO.HttpPort,
		P2pPort:          nodeDTO.P2pPort,
		Status:           0,
		Logo:             nodeDTO.Logo,
	}
}
