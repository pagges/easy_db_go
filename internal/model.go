package internal

import "time"

// easy db entry
type Entry struct {
	Checksum uint32     `json:"checkSum"`
	Key      []byte     `json:"key"`
	Value    []byte     `json:"value"`
	Offset   int64      `json:"offset"`
	Expiry   *time.Time `json:"expiry"`
}

// in-memory index container
type MemIndex struct {
	IndexEntryMap *map[string]Item
}

type Item struct {
	Key    string `json:"key"`
	FileID int    `json:"fileid"`
	Offset int64  `json:"offset"`
	Size   int64  `json:"size"`
}

type FileEngine struct {
	Path string
}
