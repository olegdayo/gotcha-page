package requesters

// Requester interface used in ParserContainer to include all parsers which implement it.
type Requester interface {
	// GetName gets name of requester.
	GetName() (name string)
	// GetNickname gets nickname of a user.
	GetNickname() (nickname string)
	// IsSelected shows if requester is selected.
	IsSelected() (selected bool)
	// SetAvailability sets availability condition.
	SetAvailability(cond bool)
	// GetInfo gets url and name of user by their nickname.
	GetInfo() (url string, name string, err error)
}
