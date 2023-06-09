package logger

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func LogHTTPRequest(entry *logrus.Entry) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		entry := entry
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			body, err := readBody(req)
			if err != nil {
				entry.WithError(err)
			}
			entry = entry.WithField("body", body)
			statusRecorder := &responseRecorder{ResponseWriter: w}
			start := time.Now()
			h.ServeHTTP(statusRecorder, req)
			elapsed := time.Since(start)
			httpCode := statusRecorder.statusCode
			entry = entry.WithField("response", string(statusRecorder.response))
			entry = entry.WithFields(logrus.Fields{
				"http_route":  req.URL.Path,
				"http_method": req.Method,
				"http_code":   httpCode,
				"elapsed":     elapsed.Milliseconds(),
			})
			if httpCode > 399 {
				entry.Level = logrus.ErrorLevel
			}
			entry.Print()
		})
	}
}

type responseRecorder struct {
	http.ResponseWriter

	statusCode int
	response   []byte
}

func (w *responseRecorder) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseRecorder) Write(b []byte) (int, error) {
	w.response = b
	return w.ResponseWriter.Write(b)
}

func readBody(r *http.Request) (map[string]interface{}, error) {
	body := make(map[string]interface{})
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &body); err != nil {
		return nil, err
	}
	defer func() {
		r.Body.Close()
		r.Body = io.NopCloser(bytes.NewBuffer(b))
		r.ContentLength = int64(bytes.NewBuffer(b).Len())
	}()
	return body, nil
}
