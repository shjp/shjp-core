package core

// Model is the interface that every model should implement
type Model interface {
	GetID() string
	Pack(Intent, OperationType) (*Message, error)
}
