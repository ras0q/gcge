package repository

import (
	"github.com/Ras96/gcg/internal/model"
)

type AnalyzerRepository interface {
	AnalyzeFile(filename string) (*model.File, error)
}
