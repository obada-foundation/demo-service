package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
)

type Hash struct {
	hash string
	dec  uint64
}

// NewHash ...
func NewHash(value string) (Hash, error) {
	var hash Hash

	h := sha256.New()

	if _, err := h.Write([]byte(value)); err != nil {
		return hash, fmt.Errorf("cannot wite bytes %v to hasher: %w", []byte(value), err)
	}

	hashStr := hex.EncodeToString(h.Sum(nil))

	hashDec, err := hashToDec(hashStr)

	if err != nil {
		return hash, err
	}

	hash.hash = hashStr
	hash.dec = hashDec

	return hash, nil
}

// hashToDec ...
func hashToDec(hash string) (uint64, error) {
	match, err := regexp.MatchString(`^[0-9a-fA-F]+$`, hash)

	if err != nil {
		return 0, fmt.Errorf("cannot check if given string %q is valid hex: %w", hash, err)
	}

	if !match {
		return 0, fmt.Errorf("given string string %q is not valid hex", hash)
	}

	if len(hash) > 8 {
		hash = hash[:8]
	}

	decimal, err := strconv.ParseUint(hash, 16, 32)

	if err != nil {
		return 0, err
	}

	return decimal, nil
}

// GetHash ...
func (h Hash) GetHash() string {
	return h.hash
}

// GetDec ...
func (h Hash) GetDec() uint64 {
	return h.dec
}
