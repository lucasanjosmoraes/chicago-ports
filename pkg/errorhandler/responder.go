package errorhandler

// Responder defines what's needed to create an error from a HTTP request.
type Responder interface {
	Status() int
	Response() []byte
	error
}
