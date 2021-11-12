package repository

import (
	"github.com/Ras96/gcg/internal/model"
)

type ParserRepository interface {
	ParseFile(filename string) (*model.File, error)
}
