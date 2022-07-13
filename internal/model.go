package internal

import "time"

// easy db entry
type Entry struct {
	Checksum uint32     `json:"Checksum"`
	Key      []byte     `json:"Key"`
	Value    []byte     `json:"Value"`
	Offset   int64      `json:"Offset"`
	Expiry   *time.Time `json:"Expiry"`
}

// in-memory index
type IndexEntry struct {
	Key   string `json:"Key"`
	Entey *Entry `json:"Entey"`
}
