package core

// OperationType is the enum type for operation type that message conveys
type OperationType string

// Operation types
const (
	UpsertOperation  = "upsert"
	CreateOperation  = "create"
	UpdateOperation  = "update"
	DeleteOperation  = "delete"
	UnknownOperation = "unknown"
)
