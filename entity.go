package core

// Entity is an abstraction for data entity transported through message queue
type Entity interface {
	Pack(Intent, OperationType) (*Message, error)
	Unpack(Message) error
}
