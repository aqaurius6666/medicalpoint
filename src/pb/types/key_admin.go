package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AdminKeyPrefix is the prefix to retrieve all Admin
	AdminKeyPrefix = "Admin/value/"
)

// AdminKey returns the store key to retrieve a Admin from the index fields
func AdminKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
