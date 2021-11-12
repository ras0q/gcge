package repository

type Repositories struct {
	Parser ParserRepository
}

func NewRepositories(parser ParserRepository) *Repositories {
	return &Repositories{
		Parser: parser,
	}
}
