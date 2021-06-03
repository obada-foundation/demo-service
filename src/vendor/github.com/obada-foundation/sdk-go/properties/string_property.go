package properties

import (
	"fmt"
	"github.com/obada-foundation/sdk-go/hash"
	"log"
)

type StringProperty struct {
	value string
	hash  hash.Hash
}

func NewStringProperty(value string, log *log.Logger, debug bool) (StringProperty, error) {
	var sp StringProperty

	if debug {
		log.Printf("\nNewStringProperty(%q)", value)
	}

	h, err := hash.NewHash(value, log, debug)

	if err != nil {
		return sp, fmt.Errorf("cannot hash %q: %w", value, err)
	}

	sp.hash = h
	sp.value = value

	return sp, nil
}

func (sp StringProperty) GetValue() string {
	return sp.value
}

func (sp StringProperty) GetHash() hash.Hash {
	return sp.hash
}
