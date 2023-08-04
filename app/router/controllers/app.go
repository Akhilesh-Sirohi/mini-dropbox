package controllers

import (
	"encoding/json"
	"github.com/mini-dropbox/app/providers"
	"net/http"
)

type ResponseBody struct {
	Success bool        `json:"success"`
	Error   error       `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type app struct{}

var (
	// App will hold the instance of Mandate Manager app
	App app
)

// NewApp will initialize the new application controller to handle the method defined
func NewApp() {
	App = app{}
}

// Get will return the welcome message
func (controller app) Get(w http.ResponseWriter, r *http.Request) {
	responseContent := "Welcome To Mini-Dropbox"
	json.NewEncoder(w).Encode(responseContent)
}

func (controller app) Status(w http.ResponseWriter, r *http.Request) {
	err := providers.DB.Ping()

	w.WriteHeader(getStatus(err))
	w.Write(getResponseBody(nil, err))
}

func getResponseBody(data interface{}, err error) []byte {
	response := ResponseBody{
		Success: err == nil,
		Error:   err,
		Data:    data,
	}

	res, _ := json.Marshal(response)
	return res
}
func getStatus(err error) int {
	if err == nil {
		return 200
	}
	return 500
}
