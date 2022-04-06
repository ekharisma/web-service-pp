package api_test

import (
	"net/http/httptest"
	"testing"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ekharisma/web-service-pp/controller/api"
	"github.com/ekharisma/web-service-pp/controller/db"
)

func TestController(t *testing.T) {
	mqttClient := mqtt.NewClient(mqtt.NewClientOptions())
	db := db.NewInMemoryDatabase()
	c := api.NewController(mqttClient, db)
	r := httptest.NewRequest("GET", "localhost:8000/thermometer", nil)
	w := httptest.NewRecorder()
	c.GetTemperature(w, r)
	if w.Result().StatusCode != 200 {
		t.Error("Error, status code is not 200")
	}
}
