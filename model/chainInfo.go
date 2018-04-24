package model

import "encoding/json"

type ChainInfo struct {
	ServerVersion            string `json:"server_version"`
	HeadBlockNum             int64  `json:"head_block_num"`
	LastIrreversibleBlockNum int64  `json:"last_irreversible_block_num"`
	HeadBlockId              string `json:"head_block_id"`
	HeadBlockTime            string `json:"head_block_time"`
	HeadBlockProducer        string `json:"head_block_producer"`
}

func (chainInfo *ChainInfo) String() string {
	data, err := json.Marshal(chainInfo)
	if err != nil {
		return ""
	}
	return string(data)
}

func ChainInfoFromString(data string) *ChainInfo {
	var chainInfo ChainInfo
	json.Unmarshal([]byte(data), &chainInfo)
	return &chainInfo
}
