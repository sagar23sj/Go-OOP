package counter

// Count is an exported identifier of integer type
type Count int

// internalCounter is an unexported identifier of integer type
type internalCounter int

// NewInternalCounter creates an object of internalCounter
// with value passed to it and returns
// an object of type internalCounter
func NewInternalCounter(val int) internalCounter {
	return internalCounter(val)
}
