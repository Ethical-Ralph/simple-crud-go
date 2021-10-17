package response

import (
	"encoding/json"
	"net/http"
)


func Success(rw http.ResponseWriter, data interface{}, msg string) {

	if msg == "" {
		msg = "Success"
	}

	json.NewEncoder(rw).Encode(map[string]interface{}{
		"success": true,
		"message": msg,
		"data": data,
	})
}