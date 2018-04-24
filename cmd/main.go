package main

import (
	"log"
	"node_monitor/api"
	"node_monitor/config"
	"node_monitor/monitor"
	"os"

	"github.com/urfave/cli"
)



func main() {
	app := cli.NewApp()
	app.Name = "node_monitor"
	app.HideVersion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "redisHost",
			Value:       "127.0.0.1",
			Usage:       "redis host",
			Destination: &config.Flags.RedisHost,
		},
		cli.StringFlag{
			Name:        "redisPort",
			Value:       "6379",
			Usage:       "redis port",
			Destination: &config.Flags.RedisPort,
		},
		cli.StringFlag{
			Name:        "cleosHost",
			Value:       "127.0.0.1",//13.251.6.13
			Usage:       "cleos host",
			Destination: &config.Flags.CleosHost,
		},
		cli.StringFlag{
			Name:        "cleosPort",
			Value:       "8800",
			Usage:       "cleos host",
			Destination: &config.Flags.CleosPort,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "run",
			Action: run,
		},
	}

	err := app.Run(os.Args)
	if err!=nil {
		log.Fatal(err)
	}

}


func run(c *cli.Context) error {
	monitor.StartMonitorLoop()
	api.Serve()
	return nil
}