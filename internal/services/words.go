package services

import (
	"github.com/mvahedii/gorem/internal/repositories/database"
)

type WordsService interface {
	CreateWord(title, description string) (int, error)
	GetWord(id int) (*database.Word, error)
}
