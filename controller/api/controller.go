package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ekharisma/web-service-pp/controller/db"
)

type APIController struct {
	client mqtt.Client
	db     db.Database
}

type Controller interface {
	GetTemperature(w http.ResponseWriter, r *http.Request)
}

func NewController(client mqtt.Client, db db.Database) Controller {
	return &APIController{
		client: client,
		db:     db,
	}
}

func (c *APIController) GetTemperature(w http.ResponseWriter, r *http.Request) {
	temperature, err := c.db.GetLastTemperatures()
	message, err := json.Marshal(temperature)
	if err != nil {
		fmt.Fprint(w, "Error", err.Error())
	}
	fmt.Fprintf(w, "%v", string(message))
}
