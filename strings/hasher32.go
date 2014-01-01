package strings

import "hash"

/**********************************************************************
 ** Public
 **********************************************************************/

// Hasher32
type Hasher32 interface {
	Hash(s string) uint32
}

// NewHasher32 creates a new Hasher32 from a hash.Hash32
func NewHasher32(h hash.Hash32) Hasher32 {
	return &hasher32{h}
}

// Hash32 is a convenience method for hashing a string against a hash.Hash32
func Hash32(s string, h hash.Hash32) uint32 {
	h.Reset()
	h.Write([]byte(s))
	return h.Sum32()
}

/**********************************************************************
 ** Private
 **********************************************************************/

// hasher32
type hasher32 struct {
	hash.Hash32
}

// hasher32:Hash
func (h *hasher32) Hash(s string) uint32 {
	h.Reset()
	h.Write([]byte(s))
	return h.Sum32()
}
