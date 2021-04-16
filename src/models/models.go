package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	UserId    string    `json:"userid" gorm:"primaryKey"`
	StartDate time.Time `json:"startdate"`
	EndDate   time.Time `json:"enddate"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
}

type User struct {
	gorm.Model
	UserId   string `gorm:"primaryKey;unique"`
	Password string
	Email    string
}

type RequestTodo struct {
	Version  string `json:"version"`
	TodoList []Todo `json:"todolist"`
}
