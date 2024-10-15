package connection

import (
	"MQTT_Middleware/Executor"
	"MQTT_Middleware/config"
	"MQTT_Middleware/databaseConnection"
	"context"
	"fmt"
	"github.com/eclipse/paho.golang/paho"
	"log"
	"net"
	"strconv"
	"time"
)

type SubClientOpt struct {
	Executor              Executor.ExecFunc
	ClientIdWithTimestamp bool
	Database              string
}

type SubClient struct {
	SubClientOpt
	closeQuery chan int
}

func NewSubClient(opts SubClientOpt) *SubClient {
	if opts.Database == "mysql" || opts.Database == "Mysql" || opts.Database == "MYSQL" {
		databaseConnection.MysqlInit()
	} else {
		panic("Unsupported database,now only can use mysql")
	}
	if opts.Executor == nil {
		opts.Executor = Executor.DefaultExecFunc
	}
	return &SubClient{
		SubClientOpt: opts,
		closeQuery:   make(chan int),
	}
}

func (client *SubClient) RunSubClient() {
	msgChan := make(chan *paho.Publish)
	conn, err := net.Dial("tcp", config.GlobalConfig.MQTT.Server.In_Server)
	if err != nil {
		log.Fatalf("Failed to connect to %s: %s", config.GlobalConfig.MQTT.Server.In_Server, err)
	}

	c := paho.NewClient(paho.ClientConfig{
		Router: paho.NewSingleHandlerRouter(func(m *paho.Publish) {
			msgChan <- m
		}),
		Conn: conn,
	})
	ClientID := ""
	if client.ClientIdWithTimestamp == false {
		ClientID = config.GlobalConfig.MQTT.Client.ClientId
	} else {
		ClientID = config.GlobalConfig.MQTT.Client.ClientId + strconv.FormatInt(time.Now().Unix(), 10)
	}
	cp := &paho.Connect{
		KeepAlive:  60,
		ClientID:   ClientID,
		CleanStart: true,
		Username:   config.GlobalConfig.MQTT.Client.Username,
		Password:   []byte(config.GlobalConfig.MQTT.Client.Password),
	}

	if config.GlobalConfig.MQTT.Client.Username != "" {
		cp.UsernameFlag = true
	}
	if config.GlobalConfig.MQTT.Client.Password != "" {
		cp.PasswordFlag = true
	}

	log.Println(cp.UsernameFlag, cp.PasswordFlag)

	ca, err := c.Connect(context.Background(), cp)
	if err != nil {
		log.Fatalln(err)
	}
	if ca.ReasonCode != 0 {
		log.Fatalf("Failed to connect to %s : %d - %s", config.GlobalConfig.MQTT.Server.In_Server, ca.ReasonCode, ca.Properties.ReasonString)
	}

	fmt.Printf("Connected to %s\n", config.GlobalConfig.MQTT.Server.In_Server)

	sa, err := c.Subscribe(context.Background(), &paho.Subscribe{
		Subscriptions: []paho.SubscribeOptions{
			{Topic: config.GlobalConfig.MQTT.Topic.In_Topic, QoS: byte(config.GlobalConfig.MQTT.Message.QOS)},
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	if sa.Reasons[0] != byte(config.GlobalConfig.MQTT.Message.QOS) {
		log.Fatalf("Failed to subscribe to %s : %d", config.GlobalConfig.MQTT.Topic.In_Topic, sa.Reasons[0])
	}
	log.Printf("Subscribed to %s", config.GlobalConfig.MQTT.Topic.In_Topic)

	for m := range msgChan {
		go func() {
			if err = client.Executor(string(m.Payload), m.Topic); err != nil {
				log.Printf("error happened:%v", err)
			}
		}()
	}
}

func (client *SubClient) Close() {
	client.closeQuery <- 0
}
