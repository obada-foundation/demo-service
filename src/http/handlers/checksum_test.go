package handlers

import (
	"encoding/json"
	"testing"

	"github.com/obada-foundation/sdkgo"
)

func TestDecodingJsonFromFrontendToDto(t *testing.T) {
	data := []byte(`
{
	"SerialNumberHash":"qw",
	"Manufacturer":"11",
	"PartNumber":"44",
	"OwnerDid":"666",
	"ObdDid":"888",
	"Status":"DISPOSED",
	"Metadata":{"ss1":"ff"}
}`)

	var dto sdkgo.ObitDto

	if err := json.Unmarshal(data, &dto); err != nil {
		t.Fatal(err)
	}
}
