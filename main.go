package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/caarlos0/env"
	_ "github.com/lib/pq"
)

type config struct {
	Port       string `env:"APP_CONTAINER_PORT"`
	DBHost     string `env:"POSTGRES_HOST"`
	DBPort     string `env:"POSTGRES_PORT"`
	DBUser     string `env:"POSTGRES_USER"`
	DBPassword string `env:"POSTGRES_PASSWORD"`
	DBName     string `env:"POSTGRES_DB"`
	DBSslmode  string `env:"POSTGRES_SSLMODE"`
}

type vocabulary struct {
	vocabularyNo int
	title        string
	meaning      string
	sentence     string
}

func getConfig() (*config, error) {
	config := &config{}
	if err := env.Parse(config); err != nil {
		return nil, err
	}
	return config, nil
}

func main() {
	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser,
		config.DBPassword, config.DBName, config.DBSslmode,
	)
	dbHandle, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	conn, err := dbHandle.Conn(ctx)
	if err != nil {
		log.Fatal(err)
	}

	r, err := conn.ExecContext(
		ctx,
		"INSERT INTO vocabularies(title, meaning, sentence) VALUES($1, $2, $3)",
		"test",
		"test meaning",
		"test sentence",
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = r.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := conn.QueryContext(ctx, "SELECT * FROM vocabularies")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	vocabularies := make([]vocabulary, 0)
	for rows.Next() {
		var vocabulary vocabulary
		if err := rows.Scan(
			&vocabulary.vocabularyNo, &vocabulary.title,
			&vocabulary.meaning, &vocabulary.sentence,
		); err != nil {
			log.Fatal(err)
		}
		vocabularies = append(vocabularies, vocabulary)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Result: %v", vocabularies)
}
