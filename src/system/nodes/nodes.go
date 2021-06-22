package nodes

import (
	"net"
	"strings"
)

// GetNodes returns a slice of available nodes taken from txt records.
func GetNodes() ([]string, error) {
	var nodes []string

	txts, err := net.LookupTXT("nodes.obada.io")

	if err != nil {
		return nodes, err
	}

	for _, record := range txts {
		recordNodes := strings.Split(record, ";")

		if len(recordNodes) > 0 {
			nodes = append(nodes, recordNodes...)
		}
	}

	return nodes, nil
}
