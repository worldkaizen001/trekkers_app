package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Profile Profile `gorm:"foreignKey:id"`
	Steps []Step `gorm:"foreignKey:id"`

	// Invites
	SentInvites     []Invite `gorm:"foreignKey:SenderID" json:"sent_invites,omitempty"`
	ReceivedInvites []Invite `gorm:"foreignKey:ReceiverID" json:"received_invites,omitempty"`

}


