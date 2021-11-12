package repository

type Repositories struct {
	Generator GeneratorRepository
	Parser    ParserRepository
}

func NewRepositories(generator GeneratorRepository, parser ParserRepository) *Repositories {
	return &Repositories{
		Generator: generator,
		Parser:    parser,
	}
}
