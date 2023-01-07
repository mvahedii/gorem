package database

import (
	"database/sql"
	"errors"
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

type WordModel struct {
	DB *sql.DB
}

func NewWordRepository(db *sql.DB) WordRepository {
	return &WordModel{
		DB: db,
	}
}

func (m *WordModel) Insert(title string, description string) (int, error) {

	stmt := `INSERT INTO words (word,description,created) 
	VALUES (?,?,UTC_TIMESTAMP())`

	res, err := m.DB.Exec(stmt, title, description)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	return int(id), nil

}

func (m *WordModel) Get(id int) (*Word, error) {
	stmt := `SELECT id,word, description, created FROM words
	WHERE id=?`

	row := m.DB.QueryRow(stmt, id)
	w := &Word{}

	err := row.Scan(&w.ID, &w.Word, &w.Description, &w.Created)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		} else {
			return nil, err
		}
	}
	return w, nil
}
