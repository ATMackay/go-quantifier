package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func doPost(path string, request interface{}) (*http.Response, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(path, "application/json", bytes.NewReader(b))
	if e := handleResponseErr(resp, err); e != nil {
		return resp, e
	}
	return resp, nil
}

func handleResponseErr(resp *http.Response, err error) error {
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		var v jsonErr
		if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
			return fmt.Errorf("cannot parse JSON body from error response: %w", err)
		}
		return fmt.Errorf(v.Err)
	}
	return nil
}

type jsonErr struct {
	Err string `json:"error"`
}

func decodeJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
