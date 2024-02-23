package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, model interface{}) {
	if body, err := io.ReadAll(r.Body); err != nil {
		if err := json.Unmarshal([]byte(body), model); err != nil {
			return 
		}
	}
}