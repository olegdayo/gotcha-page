package requesters

// Requester interface used in ParserContainer to include all parsers which implement it.
type Requester interface {
	// GetName gets name of requester.
	GetName() string
	// GetInfo gets url and name of user by their nickname.
	GetInfo() (string, string, error)
}
