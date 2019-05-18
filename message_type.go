package core

// MessageType is types of operations
type MessageType string

// Message types that indicate the type of entity a message is transporting
const (
	ModelType        MessageType = "model"
	RelationshipType MessageType = "relationship"
	FileType         MessageType = "file"
	UnknownType      MessageType = "unknown"
)
