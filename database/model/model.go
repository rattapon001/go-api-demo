package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (m *MediaObject) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = uuid.New()

	// if u.Role == "admin" {
	// 	return errors.New("invalid role")
	// }
	return
}

type User struct {
	Id        uint   `gorm:"primary_key" json:"id" gorm:"autoIncrement"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}

type MediaObject struct {
	Id            uuid.UUID `gorm:"primaryKey" gorm:"type:uuid;default:uuid_generate_v4()"`
	FilePath      string    `json:"filePath"`
	DirectoryName string    `json:"directoryName"`
	// LastName  string `json:"lastName"`
	// Age       int    `json:"age"`
	// Email     string `json:"email"`
}
