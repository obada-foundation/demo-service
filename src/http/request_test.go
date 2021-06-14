package http

import (
	"bytes"
	"log"
	"net/http"
	"testing"
)

func TestDecodingDynamicKV(t *testing.T) {
	reader := bytes.NewReader([]byte(`
{
	"SerialNumberHash":"qw",
	"Manufacturer":"11",
	"PartNumber":"44",
	"OwnerDid":"666",
	"ObdDid":"888",
	"Status":"DISPOSED",
	"Metadata":{"color":"red"}
}`))

	log.Printf("%v", reader)

	_, err := http.NewRequest(http.MethodPost, "/foo", reader)
	if err != nil {
		t.Fatalf("Failed request")
	}
}
