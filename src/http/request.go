package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dimfeld/httptreemux/v5"
)

func Param(r *http.Request, key string) string {
	m := httptreemux.ContextParams(r.Context())
	return m[key]
}

func DecodeKV(r *http.Request, dynamicKv string) (map[string]string, error) {
	var kv map[string]string
	var jsonBytes []byte

	if _, err := r.Body.Read(jsonBytes); err != nil {
		log.Printf("negroni 2 %v", err)
		return kv, err
	}

	log.Printf("negroni1 %v", jsonBytes)

	return kv, nil
}

func Decode(r *http.Request, val interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(val); err != nil {
		return err
	}

	return nil
}
