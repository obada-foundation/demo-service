package properties

import (
	"fmt"
	"github.com/obada-foundation/sdk-go/hash"
	"log"
)

type StatusProperty struct {
	value string
	hash  hash.Hash
}

func NewStatusProperty(status string, log *log.Logger, debug bool) (StatusProperty, error) {
	var sp StatusProperty

	if debug {
		log.Printf("\nNewStatusProperty(%q)", status)
	}

	h, err := hash.NewHash(status, log, debug)

	if err != nil {
		return sp, fmt.Errorf("cannot hash %q: %w", status, err)
	}

	sp.hash = h
	sp.value = status

	return sp, nil
}

func (sp StatusProperty) GetValue() string {
	return sp.value
}

func (sp StatusProperty) GetHash() hash.Hash {
	return sp.hash
}
