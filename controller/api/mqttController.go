package api

import (
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ekharisma/web-service-pp/constant"
	"github.com/ekharisma/web-service-pp/controller/db"
	"github.com/ekharisma/web-service-pp/model"
)

type MqttController interface {
	StartSubscribe() bool
	ParseMessage(message []byte) model.Temperature
}

type MqttHandler struct {
	client         mqtt.Client
	messageHandler mqtt.MessageHandler
	db             db.Database
}

func CreateNewMqttController(client mqtt.Client, database db.Database) MqttController {

	obj := MqttHandler{
		client: client,
		db:     database,
	}
	obj.messageHandler = func(c mqtt.Client, m mqtt.Message) {
		switch m.Topic() {
		case constant.Topic:
			{
				payload := obj.ParseMessage(m.Payload())
				database.StoreTemperature(payload)
			}
		}
	}
	return &obj
}

func (mh *MqttHandler) StartSubscribe() bool {
	token := mh.client.Subscribe(constant.Topic, constant.DefaultQOS, func(c mqtt.Client, m mqtt.Message) {
		payload := mh.ParseMessage(m.Payload())
		mh.db.StoreTemperature(model.Temperature(payload))
	})
	if token.Wait() {
		return false
	}
	return true
}

func (mh *MqttHandler) ParseMessage(message []byte) model.Temperature {
	fmt.Println("Message Received")
	var payload model.Temperature
	err := json.Unmarshal(message, &payload)
	if err != nil {
		panic(err.Error())
	}
	return payload
}
