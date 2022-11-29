package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (m *MediaObject) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = uuid.New()
	return
}

type MediaObject struct {
	Id            uuid.UUID `gorm:"primaryKey" gorm:"type:uuid;default:uuid_generate_v4()"`
	FilePath      string    `json:"filePath"`
	DirectoryName string    `json:"directoryName"`
}
