package core

// Intent represents the communication type of a message
type Intent string

const (
	// IntentRequest indicates a request to perform action
	IntentRequest Intent = "request"

	// IntentSuccess indicates a success result of a performed action
	IntentSuccess Intent = "success"

	// IntentFailure indicates a failure result of a performed action
	IntentFailure Intent = "failure"
)
