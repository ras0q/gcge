package repository

type Repositories struct {
	Generator GeneratorRepository
	Analyzer  AnalyzerRepository
}

func NewRepositories(generator GeneratorRepository, analyzer AnalyzerRepository) *Repositories {
	return &Repositories{
		Generator: generator,
		Analyzer:  analyzer,
	}
}
