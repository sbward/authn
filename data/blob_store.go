package data

type BlobStore interface {
	// Read fetches a blob from the store.
	Read(name string) ([]byte, error)

	// WriteNX will write the blob into the store only if the name does not exist.
	WriteNX(name string, blob []byte) (bool, error)
}
