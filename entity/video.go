package entity

import "time"

type Person struct {
	ID        uint64 `gorm:"primary_key; auto_increment" json:"id"`
	FirstName string `json::"firstname" binding:"required" gorm:"type:nvarchar(32)"`
	LastName  string `json::"lastname" binding:"required" gorm:"type:nvarchar(32)"`
	Age       int8   `json::"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email" gorm:"type:nvarchar(256)"`
}

type Video struct {
	ID          uint64    `gorm:"primary_key; auto_increment" json:"id"`
	Title       string    `json:"title" binding:"min=2,max=100" validate:"is-cool" gorm:"type:nvarchar(100)"`
	Description string    `json:"description" binding:"max=20" gorm:"type:nvarchar(200)"`
	Url         string    `json:"url" binding:"required,url" gorm:"type:nvarchar(256); UNIQUE"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignkey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreateAt    time.Time `json:"-" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt    time.Time `json:"-" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}
