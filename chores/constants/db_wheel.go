package constants

import (
	"encoding/json"
	"time"
)

type Wheel struct {
	ID         string          `json:"id" gorm:"type:varchar(255);not null;primaryKey"`
	Name       string          `json:"name" binding:"required" gorm:"type:varchar(255);not null;unique"`
	Options    json.RawMessage `json:"options" binding:"required" gorm:"type:json;not null"`
	SharedBy   string          `json:"sharedBy" gorm:"type:varchar(255)"`
	SharedFrom string          `json:"sharedFrom" gorm:"type:varchar(255)"`
	CreatedBy  string          `json:"createdBy" binding:"required" gorm:"type:varchar(255);not null;unique"`
	CreatedAt  time.Time       `json:"createdAt" gorm:"autoCreateTime"`
}

type WheelIdOpt struct {
	Id string
}

func (Wheel) TableName() string {
	return "wheel"
}
