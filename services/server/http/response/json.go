package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	blob, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		Error(w, err)
		return
	}
	w.Write(blob)
}
