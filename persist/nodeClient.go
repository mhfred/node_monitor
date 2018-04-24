package persist

import (
	"fmt"
	"node_monitor/model"
)

const NodeKey = "nodes"

func GetNodes() []model.Node {
	redisClient := newRedisClient()
	defer redisClient.Close()

	results,err:= redisClient.HGetAll(NodeKey).Result()
	if err!=nil{
		fmt.Println(err)
		return []model.Node{}
	}
	var nodes []model.Node
	for _,result :=range results{
		nodes = append(nodes,model.NodeFromJson(result))
	}
	return nodes
}

func GetActiveNodes() []model.Node{
	redisClient := newRedisClient()
	defer redisClient.Close()

	results,err:= redisClient.HGetAll(NodeKey).Result()
	if err!=nil{
		fmt.Println(err)
		return []model.Node{}
	}
	var nodes []model.Node
	for _,result :=range results{
		node := model.NodeFromJson(result)
		if node.Status==1{
			nodes = append(nodes,node)
		}
	}
	return nodes
}

func SaveNode(node model.Node){
	redisClient := newRedisClient()
	defer redisClient.Close()

	redisClient.HSet(NodeKey,node.PublicKey,node.ToJson()).Result()
}