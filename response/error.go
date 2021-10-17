package response

import (
	"encoding/json"
	"net/http"
)


func Error(rw http.ResponseWriter, msg string) {

	if msg == "" {
		msg = "Success"
	}

	json.NewEncoder(rw).Encode(map[string]interface{}{
		"success": false,
		"message": msg,
		"data": nil,
	})
}