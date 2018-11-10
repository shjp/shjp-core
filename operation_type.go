package core

type OperationType string

const (
	// UpsertOperation indicates an intent to upsert
	UpsertOperation = "upsert"

	// CreateOperation indicates an intent to create
	CreateOperation = "create"

	// UpdateOperation indicates an intent to update
	UpdateOperation = "update"

	// DeleteOperation indicates an intent to delete
	DeleteOperation = "delete"

	// UnknownOperation indicates an unknown operation
	UnknownOperation = "unknown"
)
