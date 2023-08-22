package main

import (
	"log"
	"strconv"

	"github.com/casmeyu/micro-message/configuration"
	"github.com/casmeyu/micro-message/conversationService"
	"github.com/casmeyu/micro-message/storage"
	"github.com/casmeyu/micro-message/structs"
	"github.com/casmeyu/micro-message/types"
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
		var res structs.ServiceResponse
		var conv types.Conversation
		var messages []types.Message

		user1 := 2
		user2, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Invalid user id")
		}

		db, err := storage.Open(Config)
		convRes := conversationService.StartConversation(uint(user1), uint(user2), db)
		if convRes.Success == false {
			return c.Status(res.Status).JSON(res.Err)
		}

		conv = convRes.Result.(types.Conversation)
		convMessagesRes := conversationService.GetMessages(conv.ID, db)
		if convMessagesRes.Success == false {
			return c.Status(res.Status).JSON(res.Err)
		}
		messages = convMessagesRes.Result.([]types.Message)
		return c.Status(fiber.StatusOK).JSON(map[string]interface{}{"conversation": conv, "messages": messages})
	})
}

func main() {
	Setup()

	storage.MakeMigration(Config, types.User{})
	storage.MakeMigration(Config, types.Conversation{})
	storage.MakeMigration(Config, types.Message{})

	app := fiber.New()
	SetRoutes(app)

	app.Listen(":3000")
}
