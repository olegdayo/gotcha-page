package requesters

// Requester interface used in ParserContainer to include all parsers which implement it.
type Requester interface {
	// GetName - Getting name of the parser.
	GetName() string
	// GetInfo - Getting link to user's and their name profile via their nickname.
	GetInfo() (string, string, error)
}
