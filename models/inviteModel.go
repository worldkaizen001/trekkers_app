package models

import "gorm.io/gorm"

type Invite struct {
	gorm.Model
	SenderID   uint   `json:"sender_id"`
	ReceiverID uint   `json:"receiver_id"`
	Status     string `json:"status" gorm:"default:pending"`
	Message     string `json:"message"`


	// Relationships
	Sender   User `gorm:"foreignKey:SenderID" json:"-"`
	Receiver User `gorm:"foreignKey:ReceiverID" json:"-"`
}
