package database

import (
	"database/sql"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Word struct {
	Model       gorm.Model
	ID          int
	Word        string
	Description string
	Created     time.Time
}

type WordRepository interface {
	Insert(title, description string) error
	Get(id int) (*Word, error)
}

type WordModel struct {
	DB *gorm.DB
}

func NewWordRepository(db *gorm.DB) WordRepository {
	return &WordModel{
		DB: db,
	}
}

func (m *WordModel) Insert(title string, description string) (err error) {

	word := Word{
		Model:       gorm.Model{},
		Word:        title,
		Description: description,
	}

	err = m.DB.Create(&word).Error

	if err != nil {
		return err
	}

	return nil

}

func (m *WordModel) Get(id int) (*Word, error) {

	w := &Word{}
	row := m.DB.First(w, id)
	err := row.Row().Scan(&w.ID, &w.Word, &w.Description, &w.Created)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		} else {
			return nil, err
		}
	}
	return w, nil
}
