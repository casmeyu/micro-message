package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/casmeyu/micro-message/configuration"
	"github.com/casmeyu/micro-message/conversationService"
	"github.com/casmeyu/micro-message/storage"
	"github.com/casmeyu/micro-message/structs"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Config structs.Config
var validate = validator.New()

// Executes LoadConfig() function and sets up initial information for the backend app
func Setup() error {
	err := configuration.LoadConfig(&Config)
	if err != nil {
		log.Println("Error while setting up config", err.Error())
		return err
	}
	return nil
}

func SetRoutes(app *fiber.App) {
	// Setup User Routes
	userRoutes := app.Group("/messages")
	userRoutes.Get("/:id<int>", func(c *fiber.Ctx) error {
		user1 := 1
		user2, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Invalid user id")
		}
		db, err := storage.Open(Config)
		res := conversationService.StartConversation(user1, user2, db)
		if res.Success == true {
			return c.Status(res.Status).JSON(res.Result)
		} else {
			return c.Status(res.Status).JSON(res.Err)
		}
	})
}

func main() {
	fmt.Println("hello there")
	Setup()
	app := fiber.New()
	SetRoutes(app)
	storage.MakeMigration(Config, &conversationService.Conversation{})
	app.Listen(":3000")
}
