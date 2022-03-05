package parsers

type Parser interface {
	GetName() string
	GetLink() (string, error)
}
