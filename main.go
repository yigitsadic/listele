package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigitsadic/listele/database"
	"github.com/yigitsadic/listele/handlers"
	"log"
)

type mockRepo struct {
}

func (m mockRepo) FindAll() ([]database.Person, error) {
	return []database.Person{
		{
			FullName: "Albert",
		},
		{
			FullName: "John Doe",
		},
	}, nil
}

func main() {
	app := fiber.New()

	repo := mockRepo{}
	app.Get("/hello", handlers.HandleList(repo))

	log.Fatalln(app.Listen(":3035"))
}
