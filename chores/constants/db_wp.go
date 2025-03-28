package constants

import "time"

type WpResourceType struct {
	ID        string    `json:"id" gorm:"type:varchar(255);not null;primaryKey"`
	Name      string    `json:"name" binding:"required" gorm:"type:varchar(255);not null;unique"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

type WpResourceDetail struct {
	ID        string    `json:"id" gorm:"type:varchar(255);not null;primaryKey"`
	Name      string    `json:"name" binding:"required" gorm:"type:varchar(255);not null;unique"`
	Link      string    `json:"link" binding:"required" gorm:"type:varchar(255);not null"`
	TypeId    string    `json:"typeId" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

type WpResourceTypeIdOpt struct {
	ID string
}

type WpResourceIdOpt struct {
	ID string
}

func (WpResourceType) TableName() string {
	return "wp_resource_type"
}

func (WpResourceDetail) TableName() string {
	return "wp_resource_detail"
}
