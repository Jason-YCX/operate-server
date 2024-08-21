package constants

import (
	"time"
)

type Mp3 struct {
	ID        string    `json:"id" gorm:"type:varchar(255);not null;primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null;unique"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

func (Mp3) TableName() string {
	return "mp3"
}
