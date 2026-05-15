package protocol

// CacheBlob represents a blob as used in the client side blob cache protocol. It holds a hash of its data and
// the full data of it.
//
// Added: v1.12
type CacheBlob struct {
	// Hash is the hash of the blob. The hash is computed using xxHash, and must be deterministic for the same
	// chunk data.
	//
	// Added: v1.12
	Hash uint64
	// Payload is the data of the blob. When sent, the client will associate the Hash of the blob with the
	// Payload in it.
	//
	// Added: v1.12
	Payload []byte
}

// Marshal encodes/decodes a CacheBlob.
func (x *CacheBlob) Marshal(r IO) {
	r.Uint64(&x.Hash)
	r.ByteSlice(&x.Payload)
}
