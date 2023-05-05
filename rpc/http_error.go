package rpc

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func RespondWithError(w http.ResponseWriter, code int, msg interface{}) {
	var message string
	switch m := msg.(type) {
	case error:
		message = m.Error()
	case string:
		message = m
	}
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func HandleResponseErr(resp *http.Response) error {
	if resp.StatusCode != 200 {
		var v jsonErr
		if err := DecodeJSON(resp.Body, &v); err != nil {
			return fmt.Errorf("cannot parse JSON body from error response: %w", err)
		}
		return fmt.Errorf(v.Err)
	}
	return nil
}

type jsonErr struct {
	Err string `json:"error"`
}
