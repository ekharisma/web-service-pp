package main

import (
	"log"
	"net/http"

	"github.com/ekharisma/web-service-pp/controller"
)

const broker = "broker.emqx.io"
const port = 1883
const topic = "/demo/pp/3"
const defaultQOS = 2

func main() {
	http.HandleFunc("/", controller.Hello)
	http.HandleFunc("/thermometer", controller.Temperature)
	client := controller.MqttInit(broker, port)
	if token := client.Subscribe(topic, defaultQOS, controller.ConsumeMqtt); token.Wait() && token.Error() != nil {
		panic(token.Error().Error())
	}
	log.Fatal(http.ListenAndServe(":8000", nil))
}
