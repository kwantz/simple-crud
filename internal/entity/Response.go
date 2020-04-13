package entity

import (
	"encoding/json"
	"log"
)

// SetSuccessResponse - create HTTP 200
func SetSuccessResponse(data interface{}) []byte {
	return (&response{true, "200", "Success", data}).toJSON()
}

// SetCreatedResponse - create HTTP 201
func SetCreatedResponse(data interface{}) []byte {
	return (&response{true, "201", "Created", data}).toJSON()
}

// SetServerErrorResponse - create HTTP 500
func SetServerErrorResponse(message string) []byte {
	return (&response{false, "500", message, nil}).toJSON()
}

// SetClientErrorResponse - create HTTP 400
func SetClientErrorResponse(message string) []byte {
	return (&response{false, "400", message, nil}).toJSON()
}

// SetNotFoundResponse - create HTTP 404
func SetNotFoundResponse() []byte {
	return (&response{false, "404", "Not found", nil}).toJSON()
}

type response struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (res *response) toJSON() []byte {
	jsonData, err := json.Marshal(res)
	if err != nil {
		log.Println("entity.Response.toJSON ... Error")
		log.Fatal(err.Error())
	}
	return jsonData
}
