package types

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	From         int    `json:"from" gorm:"primaryKey"`
	To           int    `json:"to" gorm:""`
	Text         string `json:"text" gorm:""`
	Conversation int    `json:"conversationId" gorm:"foreignKey"`
}
