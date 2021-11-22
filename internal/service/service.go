package service

type Services struct {
	Generator GeneratorService
	Analyzer  AnalyzerService
}

func NewServices(generator GeneratorService, analyzer AnalyzerService) *Services {
	return &Services{
		Generator: generator,
		Analyzer:  analyzer,
	}
}
