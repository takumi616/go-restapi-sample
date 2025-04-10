package persistence

import "database/sql"

type VocabPersistence struct {
	DB *sql.DB
}

func New(db *sql.DB) *VocabPersistence {
	return &VocabPersistence{db}
}
