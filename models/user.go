package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id           int
	Name         string
	Email        sql.NullString
	Password     string
	PhoneNumber  sql.NullString
	Address      string
	ProfileImage string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (user *User) TableName() string {
	return "users"
}
