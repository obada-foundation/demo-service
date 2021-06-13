package handlers

import (
	"encoding/json"
	"fmt"
	sdk_go "github.com/obada-foundation/sdk-go"
	"testing"
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

	var dto sdk_go.ObitDto

	if err := json.Unmarshal(data, &dto); err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v", dto)
}
