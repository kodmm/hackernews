package database

import(
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"fmt"
	"log"
)

var Db *sql.DB

func InitDB() {
	db, err := sql.Open("pgx", "postgres://uno:thanxaaalot@localhost:5432/hackernews")
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	Db = db

}


func CloseDB() error {
	return Db.Close()
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := pgx.WithInstance(Db, &pgx.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/postgres",
		"pgx",
		driver,
	); if err != nil {
		fmt.Println(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

}