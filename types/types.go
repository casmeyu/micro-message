package types

import "time"

type User struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Contacts []User `json:"contacts" gorm:"many2many:contacts"`
}

type Conversation struct {
	ID        uint      `json:"id" gorm:"primarykey; autoIncrement"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:null"`
	DeletedAt time.Time `json:"deletedAt" gorm:"default:null"`
	User1     uint      `json:"user1" gorm:"primaryKey; not null"`
	User2     uint      `json:"user2" gorm:"primaryKey; not null"`
}

type Message struct {
	ID             uint      `json:"id" gorm:"primarykey"`
	CreatedAt      time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"default:null"`
	DeletedAt      time.Time `json:"deletedAt" gorm:"default:null"`
	Sender         uint      `json:"from" gorm:"primaryKey"`
	Message        string    `json:"message" gorm:"not null"`
	ConversationId uint      `json:"conversationId" gorm:"foreignKey; not null"`
}
