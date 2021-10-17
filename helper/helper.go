package helper 

import (
	"net/http"
	"github.com/Ethical-Ralph/simple-crud-go/response"
)

func HandleError(err error, rw http.ResponseWriter) {
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		response.Error(rw, err.Error())
		return
	}
}

func ContentJSON(rw http.ResponseWriter) {
	rw.Header().Add("Content-Type", "application/json")
}

