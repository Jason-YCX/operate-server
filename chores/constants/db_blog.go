package constants

import (
	"time"
)

type Category struct {
	ID        string    `json:"id" gorm:"type:varchar(255);not null;primaryKey"`
	Name      string    `json:"name" binding:"required" gorm:"type:varchar(255);not null;unique"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

type Tag struct {
	ID        string    `json:"id" gorm:"type:varchar(255);not null;primaryKey"`
	Name      string    `json:"name" binding:"required" gorm:"type:varchar(255);not null;unique"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

type Article struct {
	ID          string    `json:"id" gorm:"type:varchar(255);not null;primaryKey"`
	Title       string    `json:"title" binding:"required" gorm:"type:varchar(255);not null"`
	Content     string    `json:"content" binding:"required" gorm:"type:text;not null"`
	CategoryIds string    `json:"categoryIds" gorm:"type:varchar(255)"`
	TagIds      string    `json:"tagIds" gorm:"type:varchar(255)"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

type IdOpt struct {
	Id string
}

func (Category) TableName() string {
	return "category"
}

func (Tag) TableName() string {
	return "tag"
}

func (Article) TableName() string {
	return "article"
}
