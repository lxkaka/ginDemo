package models

import "time"

type Author struct {
	ID       uint64 `json:"id,omitempty" gorm:"primary_key; auto_increment"`
	Username string `json:"username" binding:"required" gorm:"type:varchar(20);unique_index"`
	Email    string `json:"email" binding:"required,email" gorm:"type:varchar(100)"`
	Password string `json:"password" binding:"required" gorm:"type:varchar(10)"`
}

type Video struct {
	ID          uint64    `json:"id,omitempty" gorm:"primary_key; auto_increment"`
	Title       string    `json:"title" binding:"min=2,max=50" gorm:"type:varchar(50)"`
	Description string    `json:"description" binding:"min=10,max=100" gorm:"type:varchar(100)"`
	Url         string    `json:"url" binding:"min=10,max=100" gorm:"type:varchar(100)"`
	Author      Author    `json:"author,omitempty" binding:"-" gorm:"foreignkey:AuthorID"`
	AuthorID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
