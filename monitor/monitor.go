package monitor

import (
	"encoding/json"
	"errors"
	"fmt"
	"node_monitor/config"
	"node_monitor/model"
	"node_monitor/persist"
	"os/exec"
	"strings"
	"time"
	"bytes"
	"context"
)

const AppCleos = "cleos"

func StartMonitorLoop() {
	host := config.Flags.CleosHost
	port := config.Flags.CleosPort

	fmt.Println(host, port)

	ticker1 := time.NewTicker(1 * time.Second)
	ticker2 := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker1.C:
				doOnceChainInfo(host, port)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ticker2.C:
				doOnceHeartBeat()
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ticker2.C:
				setProducer(host, port)
			}
		}
	}()
}

func doOnceHeartBeat() {
	nodes := persist.GetNodes()
	for _, node := range nodes {
		if node.Domain == "localhost" ||
			node.Domain == "0.0.0.0" ||
			node.Domain == "127.0.0.1" {
			node.Status = 0
		} else {
			_, err := getChainInfo(node.Domain, node.HttpPort)
			if err != nil {
				node.Status = 0
			} else {
				node.Status = 1
			}
		}

		persist.SaveNode(node)
	}
}

func doOnceChainInfo(host string, port string) {
	chainInfo, err := getChainInfo(host, port)
	if err != nil {
		return
	}
	persist.SaveChainInfo(chainInfo)
}

func getChainInfo(host string, port string) (*model.ChainInfo, error) {
	argString := fmt.Sprintf("-H %s -p %s get info", host, port)
	args := strings.Fields(argString)

	//cmd := exec.Command(AppCleos, args...)
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, AppCleos, args...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("chain info error ", stderr.String())
		return nil, errors.New(err.Error())
	}

	return dtoMapChainInfo([]byte(out.String()))
}

func dtoMapChainInfo(data []byte) (*model.ChainInfo, error) {
	var chainInfo model.ChainInfo
	err := json.Unmarshal(data, &chainInfo)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &chainInfo, nil
}

func CreateAccount(name string, publicKey string) error {
	host := config.Flags.CleosHost
	port := config.Flags.CleosPort
	argString := fmt.Sprintf("-H %s -p %s --wallet-port 8000 create account eosio %s %s %s", host, port, name, publicKey, publicKey)
	args := strings.Fields(argString)

	cmd := exec.Command(AppCleos, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("set producer error " + fmt.Sprint(err) + ": " + stderr.String())
		return errors.New("create account error")
	}
	fmt.Println("Result: " + out.String())
	return nil
}

func CreateKey() (privateKey string, publicKey string, err error) {
	host := config.Flags.CleosHost
	port := config.Flags.CleosPort
	argString := fmt.Sprintf("-H %s -p %s  create key", host, port)
	args := strings.Fields(argString)

	cmd := exec.Command(AppCleos, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("set producer error " + fmt.Sprint(err) + ": " + stderr.String())
		return "", "", errors.New("create account error")
	}
	fmt.Println("Result: " + out.String())
	s := out.String()
	splitFunc := func(r rune) bool {
		return r == ':' || r == '\n'
	}
	items := strings.FieldsFunc(s, splitFunc)
	if len(items) != 4 {
		return "", "", errors.New("key parse error")
	}
	fmt.Println(items)
	return strings.TrimSpace(items[1]), strings.TrimSpace(items[3]), nil
}

func setProducer(host string, port string) {
	producerData := generateProducerJson()
	//export PATH=PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:$PATH ;\n echo $PATH > /home/ubuntu/root_node/PATH ;\n
	//argString := fmt.Sprintf("cleos -H %s -p %s push action eosio setprods '%s' -p eosio -j > /home/ubuntu/root_node/result",host,port,producerData)
	argString := fmt.Sprintf("-H %s -p %s --wallet-port 8000 push action eosio setprods %s -p eosio -j ", host, port, producerData)
	args := strings.Fields(argString)
	//err := ioutil.WriteFile("/home/ubuntu/root_node/xiaofeng.sh", []byte(argString), 0744)

	//fmt.Println(err)

	fmt.Println(argString)
	cmd := exec.Command(AppCleos, args...)
	//cmd.Env = append(cmd.Env, "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("set producer error " + fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}

func generateProducerJson() string {
	activeNodes := persist.GetActiveNodes()
	producers := make([]ProducerData, 0)
	for _, n := range activeNodes {
		p := ProducerData{ProducerName: n.ProducerName, BlockSigningKey: n.PublicKey}
		producers = append(producers, p)
	}
	producers = append(producers, ProducerData{ProducerName: "eosio", BlockSigningKey: "EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV"})

	producerData := ProducersData{Version: int32(time.Now().Unix()), Producers: producers}

	data, err := json.Marshal(producerData)
	if err != nil {
		return ""
	}
	return string(data)
}

type ProducersData struct {
	Version   int32          `json:"version"`
	Producers []ProducerData `json:"producers"`
}

type ProducerData struct {
	ProducerName    string `json:"producer_name"`
	BlockSigningKey string `json:"block_signing_key"`
}
