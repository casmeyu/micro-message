package conversationService

import (
	"log"
	"time"

	"github.com/casmeyu/micro-message/structs"
	"github.com/casmeyu/micro-message/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func StartConversation(userId1 uint, userId2 uint, db *gorm.DB) structs.ServiceResponse {
	getConvResponse := GetConversation(userId1, userId2, db)
	if getConvResponse.Success == true {
		return getConvResponse
	} else {
		return CreateConversation(userId1, userId2, db)
	}
}

func GetConversation(userId1 uint, userId2 uint, db *gorm.DB) structs.ServiceResponse {
	var conv types.Conversation
	res := structs.ServiceResponse{}

	tx := db.Model(&types.Conversation{}).Where("(user1 = ? AND user2 = ?) OR (user1 = ? AND user2 = ?)", userId1, userId2, userId2, userId1).First(&conv)
	if tx.Error != nil {
		log.Println("[ConversationService] (GetConversation) - Error occurred while getting conversation", tx.Error.Error())
		res.Err = tx.Error.Error()
		res.Status = fiber.StatusBadRequest
		return res
	}

	res.Success = true
	res.Status = fiber.StatusOK
	res.Result = conv
	return res
}

func CreateConversation(userId1 uint, userId2 uint, db *gorm.DB) structs.ServiceResponse {
	var conv types.Conversation
	res := structs.ServiceResponse{}

	tx := db.Model(&types.Conversation{}).Where("(user1 = ? AND user2 = ?) OR (user1 = ? AND user2 = ?)", userId1, userId2, userId2, userId1).First(&conv)
	if tx.Error != nil {
		// If there is no conversation create it
		if tx.Error.Error() == "record not found" {
			conv.CreatedAt = time.Now()
			conv.User1 = userId1
			conv.User2 = userId2
			conv.ID = 2

			tx := db.Create(conv) // Adding conversation to database
			if tx.Error != nil {
				log.Println("Error while creating conversation", tx.Error)
				res.Err = tx.Error.Error()
				res.Status = fiber.StatusInternalServerError
				return res
			}
			res.Success = true
			res.Status = fiber.StatusCreated
			res.Result = conv
			return res
		}
		res.Success = false
		res.Status = fiber.StatusBadRequest
		res.Err = tx.Error.Error()
		return res
	}

	res.Success = false
	res.Status = fiber.StatusBadRequest
	res.Err = "Conversation already exists"
	return res
}

func GetMessages(convId uint, db *gorm.DB) structs.ServiceResponse {
	var res structs.ServiceResponse
	var conv types.Conversation
	var messages []types.Message
	var tx *gorm.DB

	tx = db.Model(&types.Conversation{}).Where("id = ?", convId).First(&conv)
	if tx.Error != nil {
		log.Println("[ConversationService] (GetMessages) - Error occurred while getting conversation by id", tx.Error.Error())
		res.Err = "Error getting conversation"
		res.Status = fiber.StatusInternalServerError
		return res
	}
	tx = db.Model(&types.Message{}).Where("conversation_id = ?", convId).Find(&messages)
	if tx.Error != nil {
		log.Println("[ConversationService] (GetMessages) - Error occurred while getting conversation messages", tx.Error.Error())
		res.Err = "Error getting conversation messages"
		res.Status = fiber.StatusInternalServerError
		return res
	}

	res.Success = true
	res.Status = fiber.StatusOK
	res.Result = messages
	return res
}
