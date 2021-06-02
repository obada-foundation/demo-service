package properties

import (
	"fmt"
	"github.com/obada-foundation/sdk-go/hash"
	"strconv"
)

type KvProperty struct {
	records []Record
	hash    hash.Hash
}

type Record struct {
	key   StringProperty
	value StringProperty
	hash  hash.Hash
}

func NewRecord(key string, value string) (Record, error) {
	var r Record

	k, err := NewStringProperty(key)
	v, err := NewStringProperty(value)

	if err != nil {
		return r, err
	}

	kh := k.GetHash()
	vh := k.GetHash()
	kvDec := kh.GetDec() + vh.GetDec()

	h, err := hash.NewHash(strconv.FormatUint(kvDec, 10))

	if err != nil {
		return r, err
	}

	r.key = k
	r.value = v
	r.hash = h

	return r, nil
}

func (r *Record) GetKey() StringProperty {
	return r.key
}

func (r *Record) GetValue() StringProperty {
	return r.value
}

func (r *Record) GetHash() hash.Hash {
	return r.hash
}

func NewMapProperty(kv map[string]string) (KvProperty, error) {
	var mp KvProperty

	var kvDec uint64

	for key, value := range kv {
		r, err := NewRecord(key, value)

		if err != nil {
			return mp, err
		}

		rh := r.GetHash()
		kvDec += rh.GetDec()

		mp.records = append(mp.records, r)
	}

	h, err := hash.NewHash(strconv.FormatUint(kvDec, 10))

	if err != nil {
		return mp, fmt.Errorf("cannot hash %q: %w", kvDec, err)
	}

	mp.hash = h

	return mp, nil
}

func (mp *KvProperty) GetAll() []Record {
	return mp.records
}

func (mp *KvProperty) GetHash() hash.Hash {
	return mp.hash
}
