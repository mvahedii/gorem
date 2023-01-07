package services

import (
	"github.com/mvahedii/gorem/internal/repositories/database"
)

type WordsService interface {
	CreateWord(title, description string) error
	GetWord(id int) (*database.Word, error)
}
