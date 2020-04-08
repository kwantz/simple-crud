package handler

import (
	"encoding/json"
	"log"
)

type response struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (an response) createJSON() []byte {
	res, err := json.Marshal(an)
	if err != nil {
		log.Println("Handler.Product.createJSON ... Error")
		log.Fatal(err.Error())
	}
	return res
}
