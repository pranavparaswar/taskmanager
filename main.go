package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pranavparaswar/taskmanager/bootstrap"
	"github.com/pranavparaswar/taskmanager/repository"
)

type Repository repository.Repository

func main() {
	app := fiber.New()
	bootstrap.InitializeApp(app)
}
