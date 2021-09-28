package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Answer struct {
	ID        string `json:"id" gorm:"primary_key;type:uuid;"`
	ReviewId  string `jsom:"-" gorm:"type:uuid REFERENCES review(id)"`
	Answer    string `json:"answer"`
	CreatedAt string `json:"created_at"`
	SourceId  string `json:"source_id"`
	StatusId  string `json:"status_id"`
}

func (a *Answer) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.NewString()
	return
}
