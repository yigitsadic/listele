package main

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/yigitsadic/listele/database"
	"github.com/yigitsadic/listele/handlers"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://myexampleuser:myexample@localhost:6500/listele?sslmode=disable")
	if err != nil {
		log.Fatalf("unable to open connection due to=%q", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("unable to ping database")
	}

	repo := database.PeopleRepository{Database: db}

	app := fiber.New()

	app.Get("/hello", handlers.HandleList(&repo))

	log.Fatalln(app.Listen(":3035"))
}
