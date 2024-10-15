package main

import (
	"MQTT_Middleware/Executor"
	"MQTT_Middleware/config"
	"MQTT_Middleware/connection"
	"MQTT_Middleware/util"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.Init()

	util.IdCreator, _ = util.NewWorker(1)
	cliOpt := connection.SubClientOpt{
		Executor:              Executor.DatabaseExecFunc,
		ClientIdWithTimestamp: true,
		Database:              "mysql",
	}
	cli := connection.NewSubClient(cliOpt)

	ic := make(chan os.Signal, 1)
	signal.Notify(ic, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ic
		cli.Close()
		log.Println("signal received, exiting")
		os.Exit(0)
	}()

	go cli.RunSubClient()
	time.Sleep(1 * time.Second)
	select {}

}
