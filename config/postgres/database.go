package postgres

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Postgres struct {
	DB *gorm.DB
}

func NewPostgres() (*Postgres, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_DATABASE_HOST"),
		os.Getenv("POSTGRES_DATABASE_USERNAME"),
		os.Getenv("POSTGRES_DATABASE_NAME"),
		os.Getenv("POSTGRES_DATABASE_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE_PORT"),
	)

	fmt.Println(dsn)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return &Postgres{}, err
	}

	return &Postgres{DB: db}, err
}
