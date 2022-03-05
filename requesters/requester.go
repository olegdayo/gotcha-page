package requesters

// Requester interface used in ParserContainer to include all parsers which implement it.
type Requester interface {
	// GetName - Getting name of the parser.
	GetName() string
	// GetLink - Getting link to user's profile via their nickname.
	GetLink() (string, error)
}
