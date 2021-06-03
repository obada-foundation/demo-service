package properties

import (
	"github.com/obada-foundation/sdk-go/hash"
	"log"
	"time"
)

type TimeProperty struct {
	value time.Time
	hash  hash.Hash
}

func NewTimeProperty(time time.Time, log *log.Logger, debug bool) (TimeProperty, error) {
	var tp TimeProperty

	if debug {
		log.Printf("\nNewTimeProperty(%v)", time)
	}

	h, err := hash.NewHash(time.String(), log, debug)

	if err != nil {
		return tp, err
	}

	tp.hash = h
	tp.value = time

	return tp, nil
}

func (sp TimeProperty) GetValue() time.Time {
	return sp.value
}

func (sp TimeProperty) GetHash() hash.Hash {
	return sp.hash
}
