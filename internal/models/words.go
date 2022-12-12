package models

import (
	"database/sql"
	"errors"
)

type Words struct {
	ID          int
	word        string
	description string
}

type WordModel struct {
	DB *sql.DB
}

func (m *WordModel) Insert(word string, description string) (int, error) {

	stmt := `INSERT INTO words (word,description,created) 
	VALUES (?,?,UTC_TIMESTAMP())`

	_, err := m.DB.Exec(stmt, word, description)

	if err != nil {
		return 0, err
	}

	return 1, nil

}

func (m *WordModel) Get(id int) (*Words, error) {
	stmt := `SELECT id,word,description FROM words
	WHERE id=?`

	row := m.DB.QueryRow(stmt, id)
	w := &Words{}

	err := row.Scan(&w.ID, &w.word, &w.description)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return w, nil
}
