package response

import (
	"log"
	"net/http"
)

type errorResponseBody struct {
	Error string `json:"error"`
}

func Error(w http.ResponseWriter, err error) {
	log.Println("Error:", err)
	status, message := errorStatusCodeAndMessage(err)
	w.WriteHeader(status)
	JSON(w, &errorResponseBody{
		Error: message,
	})
}

func errorStatusCodeAndMessage(err error) (int, string) {
	switch err {
	default:
		return http.StatusInternalServerError, "Something went wrong"
	}
}
