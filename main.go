package main

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-migrate/migrate/v4"
	"github.com/yigitsadic/listele/database"
	"github.com/yigitsadic/listele/handlers"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	dataSource := os.Getenv("DATASOURCE")
	if dataSource == "" {
		dataSource = "postgres://listele_user:lorems@database:5432/listele_project?sslmode=disable"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3035"
	}

	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatalf("unable to open connection due to=%q", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("unable to ping database, err=", err)
	}

	m, err := migrate.New("file://db/migrations", dataSource)
	if err != nil {
		log.Fatal("unable to run migrations due to ", err)
	}

	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("error occurred during execution of migrations ", err)
	}

	repo := database.PeopleRepository{Database: db}

	app := fiber.New()

	app.Get("/", handlers.HandleList(&repo))

	log.Fatalln(app.Listen(":" + port))
}
