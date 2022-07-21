package internal

import (
	"hash/crc32"
	"time"
)

// NewEntry creates a new `Entry` with the given `key` and `value`
func NewEntry(key, value []byte, expiry *time.Time) Entry {
	checksum := crc32.ChecksumIEEE(value)

	return Entry{
		Checksum: checksum,
		Key:      key,
		Value:    value,
		Expiry:   expiry,
	}
}
