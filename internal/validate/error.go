package validate

// ValidationError ...
type ValidationError struct {
	reason string
}

func (e *ValidationError) Error() string {
	return e.reason
}
