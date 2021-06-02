package properties

import (
	"fmt"
	"github.com/obada-foundation/sdk-go/hash"
)

type StringProperty struct {
	value string
	hash  hash.Hash
}

func NewStringProperty(value string) (StringProperty, error) {
	var sp StringProperty

	h, err := hash.NewHash(value)

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
