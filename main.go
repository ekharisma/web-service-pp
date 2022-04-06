package main

import (
	"log"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/ekharisma/web-service-pp/constant"
	"github.com/ekharisma/web-service-pp/controller/api"
	"github.com/ekharisma/web-service-pp/controller/db"
	mqttPkg "github.com/ekharisma/web-service-pp/controller/mqtt"
)

func main() {
	mqttPkg.CreateMqttClient(constant.Broker, constant.Port)
	var (
		mqttClient mqtt.Client = mqttPkg.GetMqttClient()
		database   db.Database = db.NewMySQLDatabase(constant.UsernameDB, constant.PasswordDB, constant.HostDB, constant.NameDB, constant.PortDB)
		// database       db.Database        = db.NewInMemoryDatabase()
		controller     api.Controller     = api.NewController(mqttClient, database)
		mqttController api.MqttController = api.CreateNewMqttController(mqttClient, database)
	)
	if mqttController.StartSubscribe() {
		log.Panic("Can't connect")
	}
	http.HandleFunc("/thermometer", controller.GetTemperature)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
