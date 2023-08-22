package messageService

import (
	"fmt"
	"log"
	"time"

	"github.com/casmeyu/micro-message/conversationService"
	"github.com/casmeyu/micro-message/structs"
	"github.com/casmeyu/micro-message/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SendMessage(fromUser uint, toUser uint, messageText string, db *gorm.DB) structs.ServiceResponse {
	res := structs.ServiceResponse{}

	convResponse := conversationService.StartConversation(fromUser, toUser, db)
	if convResponse.Success == false {
		log.Println("[MessageService] (SendMessage) - Error occurred trying to send a message", convResponse.Err)
		return convResponse
	}
	conv, ok := convResponse.Result.(types.Conversation)
	if ok == false {
		log.Println("[MessageService] (SendMessage) - Error occurred trying to parse conversation data")
		res.Err = "Error while getting coversation info"
	}
	fmt.Println("Got the conversation", conv)
	message := &types.Message{
		CreatedAt:      time.Now(),
		Sender:         fromUser,
		Message:        messageText,
		ConversationId: conv.ID,
	}
	tx := db.Create(message)
	if tx.Error != nil {
		log.Println("[MessageService] (SendMessage) - Error occurred inserting message to the database", tx.Error)
		res.Err = tx.Error.Error()
		res.Status = fiber.StatusInternalServerError
		return res
	}
	res.Success = true
	res.Status = fiber.StatusCreated
	res.Result = message
	return res
}
