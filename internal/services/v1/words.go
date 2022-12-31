package v1

import (
	"github.com/mvahedii/gorem/internal/repositories"
	"github.com/mvahedii/gorem/internal/services"
)

type WordsService struct {
	wordRepository repositories.WordRepository
}

func NewWordService(wordRepository repositories.WordRepository) services.WordsService {
	return &WordsService{
		wordRepository: wordRepository,
	}
}

func (wordsService *WordsService) GetWord(id int) (*repositories.Word, error) {
	word, err := wordsService.wordRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return word, nil
}

func (wordsService *WordsService) CreateWord(title, description string) (int, error) {
	res, err := wordsService.wordRepository.Insert(title, description)
	if err != nil {
		return 0, err
	}
	return res, nil
}
