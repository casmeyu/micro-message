package conversationService

import (
	"log"
	"time"

	"github.com/casmeyu/micro-message/structs"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Conversation struct {
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:null"`
	DeletedAt time.Time `json:"deletedAt" gorm:"default:null"`
	User1     int       `json:"user1" gorm:"primaryKey; not null"`
	User2     int       `json:"user2" gorm:"primaryKey; not null"`
}

func StartConversation(userId1 int, userId2 int, db *gorm.DB) structs.ServiceResponse {
	getConvResponse := GetConversation(userId1, userId2, db)
	if getConvResponse.Success == true {
		return getConvResponse
	} else {
		return CreateConversation(userId1, userId2, db)
	}
}

func GetConversation(userId1 int, userId2 int, db *gorm.DB) structs.ServiceResponse {
	var conv Conversation
	res := structs.ServiceResponse{}

	tx := db.Model(&Conversation{}).Where("user1 = ?", userId1).Where("user2 = ?", userId2).First(&conv)
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

func CreateConversation(userId1 int, userId2 int, db *gorm.DB) structs.ServiceResponse {
	var conv Conversation
	res := structs.ServiceResponse{}
	tx := db.Model(&Conversation{}).Where("user1 = ?", userId1).Where("user2 = ?", userId2).First(&conv)
	if tx.Error != nil {
		// If there is no conversation create it
		if tx.Error.Error() == "record not found" {
			conv.CreatedAt = time.Now()
			conv.User1 = userId1
			conv.User2 = userId2
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
