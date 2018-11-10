package core

// Image encapsulates image file information
type Image struct {
	Key   string
	Type  string
	Data  []byte
	Model Model
}
