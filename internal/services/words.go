package services

import (
	"github.com/mvahedii/gorem/internal/repositories"
)

type WordsService interface {
	CreateWord(title, description string) (int, error)
	GetWord(id int) (*repositories.Word, error)
}
