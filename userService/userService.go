package userService

import (
	"log"

	"github.com/casmeyu/micro-message/structs"
	"github.com/casmeyu/micro-message/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllContacts(userId uint, db *gorm.DB) structs.ServiceResponse {
	res := structs.ServiceResponse{}
	var contacts types.User
	var err error

	err = db.Model(&types.User{}).Preload("Contacts").Where("id = ?", userId).Find(&contacts).Error
	if err != nil {
		log.Println("[UserService] (GetAllContacts) - Error occured while getting user contacts", err.Error())
		res.Err = "An error occurred loading user contacts"
		res.Status = fiber.StatusInternalServerError
		return res
	}

	res.Success = true
	res.Status = fiber.StatusOK
	res.Result = contacts
	return res
}

func AddContact(userId uint, contactId uint, db *gorm.DB) structs.ServiceResponse {
	res := structs.ServiceResponse{}
	var contact types.User
	var err error

	err = db.Model(&types.User{}).Where("id = ?", contactId).First(&contact).Error
	if err != nil {
		log.Println("[UserService] (AddContact) - Contact trying to be added does not exist")
		res.Status = fiber.StatusBadRequest
		res.Err = "User does not exist"
		return res
	}

	err = db.Model(&types.User{}).Association("Contacts").Append(&contact)
	if err != nil {
		log.Println("[UserService] (AddContact) - Error occurred while adding contact", err.Error())
		res.Status = fiber.StatusInternalServerError
		res.Err = err.Error()
		return res
	}

	return res
}

func DeleteContact(userId uint, contactId uint, db *gorm.DB) structs.ServiceResponse {
	res := structs.ServiceResponse{}
	var err error

	err = db.Model(&types.User{}).Where("id = ?", userId).Association("Contacts").Delete(&types.User{ID: contactId})
	if err != nil {
		log.Println("[UserService] (DeleteContact) - Error occurred while deleting contact", err.Error())
		res.Status = fiber.StatusInternalServerError
		res.Err = err.Error()
		return res
	}

	res.Success = true
	res.Status = fiber.StatusOK
	res.Result = "Contact deleted"
	return res
}
