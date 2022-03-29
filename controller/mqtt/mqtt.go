package mqtt

import (
	"fmt"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectionLostHandler mqtt.ConnectionLostHandler = func(c mqtt.Client, err error) {
	fmt.Println("Connection Lost due to", err.Error())
}

var mqttClient mqtt.Client = nil
var mqttSingleton sync.Once

func CreateMqttClient(broker string, port int) {
	mqttSingleton.Do(func() {
		opts := mqtt.NewClientOptions()
		opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
		opts.OnConnect = connectHandler
		opts.ClientID = "pp-2-web-service"
		opts.OnConnectionLost = connectionLostHandler
		mqttClient = mqtt.NewClient(opts)
		if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	})
}

func GetMqttClient() mqtt.Client {
	return mqttClient
}
