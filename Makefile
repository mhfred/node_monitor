all: build

build:
	go build -o node_monitor cmd/main.go
	chmod +x node_monitor

run: build
	./node_monitor -redisHost 127.0.0.1 -cleosHost 127.0.0.1 run


build-linux:
	env GOOS=linux GOARCH=amd64 go build -o .build/linux/node_monitor cmd/main.go

stage: build-linux
#care api url
	#scp -i ~/aws/key/eospro.pem .build/linux/node_monitor ubuntu@13.251.11.39:/home/ubuntu/monitor


release: build-linux
	scp -i ~/aws/key/eospro.pem .build/linux/node_monitor ubuntu@13.251.6.13:/home/ubuntu/monitor


PHONY: build build-linux stage release