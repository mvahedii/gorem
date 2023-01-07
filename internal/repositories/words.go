package repositories

import (
	"time"
)

type Word struct {
	ID          int
	Word        string
	Description string
	Created     time.Time
}

type WordRepository interface {
	Insert(title, description string) (int, error)
	Get(id int) (*Word, error)
}
