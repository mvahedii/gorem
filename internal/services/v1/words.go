package v1

import (
	"github.com/mvahedii/gorem/internal/repositories/database"
	"github.com/mvahedii/gorem/internal/services"
)

type WordsService struct {
	wordRepository database.WordRepository
}

func NewWordService(wordRepository database.WordRepository) services.WordsService {
	return &WordsService{
		wordRepository: wordRepository,
	}
}

func (wordsService *WordsService) GetWord(id int) (*database.Word, error) {
	word, err := wordsService.wordRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return word, nil
}

func (wordsService *WordsService) CreateWord(title, description string) error {
	err := wordsService.wordRepository.Insert(title, description)
	if err != nil {
		return err
	}
	return nil
}
