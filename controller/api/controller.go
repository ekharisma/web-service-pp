package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ekharisma/web-service-pp/model"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func Temperature(w http.ResponseWriter, r *http.Request) {
	temperature := GetLastTemperatures()
	payload := model.Payload{
		Timestamp:   time.Now(),
		Temperature: [2]float32{temperature[0], temperature[1]},
	}
	message, err := json.Marshal(payload)
	if err != nil {
		fmt.Fprint(w, "Error", err.Error())
	}
	fmt.Fprintf(w, "%v", string(message))
}
