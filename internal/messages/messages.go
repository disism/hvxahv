package messages

import (
	"github.com/disism/hvxahv/pkg/cockroach"
	"gorm.io/gorm"
	"log"
)

type Messages struct {
	gorm.Model
	Actor     string `gorm:"type:varchar(999);actor"`
	EventType string `gorm:"type:varchar(999);event_type"`
	ID        string `gorm:"type:varchar(999);id"`
	Username  string `gorm:"primaryKey;type:varchar(999);username"`
}

func (m *Messages) New() {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Messages{}); err != nil {
		log.Printf("failed to automatically create database: %v", err)
	}

	if err1 := db.Debug().Table("messages").Create(&m).Error; err1 != nil {
		log.Printf("an error occurred while creating the messages: %v", err1)
	}
}

func NewMessages(actor string, types string, ID string, username string) *Messages {
	return &Messages{Actor: actor, EventType: types, ID: ID, Username: username}
}

type Message interface {
	New()
}
