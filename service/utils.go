package service

import (
	"encoding/json"
	"log"
	"net/http"
)

//SendJSONResponse sends a struct as a json response
func SendJSONResponse(w http.ResponseWriter, data interface{}, status int) {
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to encode a JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	_, err = w.Write(body)
	if err != nil {
		log.Printf("Failed to write the response body: %v", err)
		return
	}
}

//ReceiveAsJSON parses the http request and extracts the body
func ReceiveAsJSON(req *http.Request, data interface{}) error {
	decoder := json.NewDecoder(req.Body)
	return decoder.Decode(data)
}
