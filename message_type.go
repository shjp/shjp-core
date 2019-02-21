package core

// MessageType is types of operations
type MessageType string

const (
	// ModelType indicates a message transporting a model info
	ModelType MessageType = "model"

	// Relationship indicates a message transporting a relationship info
	RelationshipType MessageType = "relationship"

	// FileType indicates a message transporting a file info
	FileType MessageType = "file"

	// UnknownType indicates a message whose content could not be determined
	UnknownType MessageType = "unknown"
)
