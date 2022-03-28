package controller

import (
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ekharisma/web-service-pp/controller/db"
	"github.com/ekharisma/web-service-pp/model"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectionLostHandler mqtt.ConnectionLostHandler = func(c mqtt.Client, err error) {
	fmt.Println("Connection Lost due to", err.Error())
}

func MqttInit(broker string, port int) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.ClientID = "pp-2-web-service"
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectionLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

func ConsumeMqtt(client mqtt.Client, message mqtt.Message) {
	// storeto database
	payload := message.Payload()
	timestamp, temperature := parseMessage(payload)
	db.StoreTemperature(model.Temperature{
		timestamp,
		temperature,
	})
}

func parseMessage(message []byte) (time.Time, [2]float32) {
	fmt.Println("Message Received")
	var payload model.Payload
	err := json.Unmarshal(message, &payload)
	if err != nil {
		panic(err.Error())
	}
	return payload.Timestamp, payload.Temperature
}
