package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func PostGetJson(w http.ResponseWriter, r *http.Request, obj interface{}) {
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//To allocate slice for request body
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Read body data to parse json
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body[:length], obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
