package strings

import "hash"

/**********************************************************************
 ** Public
 **********************************************************************/

// Hasher64
type Hasher64 interface {
	Hash(s string) uint64
}

// NewHasher64 creates a new Hasher64 from a hash.Hash64
func NewHasher64(h hash.Hash64) Hasher64 {
	return &hasher64{h}
}

// Hash64 is a convenience method for hashing a string against a hash.Hash64
func Hash64(s string, h hash.Hash64) uint64 {
	h.Reset()
	h.Write([]byte(s))
	return h.Sum64()
}

/**********************************************************************
 ** Private
 **********************************************************************/

// hasher64
type hasher64 struct {
	hash.Hash64
}

// hasher64:hash
func (h *hasher64) Hash(s string) uint64 {
	h.Reset()
	h.Write([]byte(s))
	return h.Sum64()
}
