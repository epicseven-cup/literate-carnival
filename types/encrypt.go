package types

type EncryptedObject struct {
	Object []byte
	Tag    []byte
}

type SignedObject struct {
	Object    []byte
	Signature []byte
	PublicKey []Multihash
}
