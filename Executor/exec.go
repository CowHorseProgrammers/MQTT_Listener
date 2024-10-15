package Executor

import (
	"MQTT_Middleware/databaseConnection"
	"errors"
	"log"
	"strings"
)

type ExecFunc func(string, string) error

func DatabaseExecFunc(jsonStr, topic string) error {
	if databaseConnection.MysqlClient == nil {
		return errors.New("mysql not connection")
	}
	slice := strings.Split(topic, "/")
	if slice[len(slice)-1] == "status" {
		if err := SaveStatusToDatabase(jsonStr, slice[len(slice)-3]); err != nil {
			return err
		}
	} else if slice[len(slice)-1] == "control" {
		if err := SaveHistoryToDatabase(jsonStr, slice[len(slice)-3]); err != nil {
			return err
		}
	} else {
		return errors.New("Unknown Topic ")
	}
	return errors.New("I dont know why it can come here :(")
}

func DefaultExecFunc(jsonStr, topic string) error {
	log.Printf("receive msg: %v from topic: %v", jsonStr, topic)
	return nil
}
