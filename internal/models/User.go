package models

import "time"

type User struct {
	Id               int
	Login            string
	Password         string
	Email            string
	Name             string
	LastName         string
	Type             int
	DateTimeCreation time.Time
	DateOfBirth      time.Time
	IsDeleted        bool
	IsBanned         bool
	IsConfirmed      bool
}

func NewUser() *User {
	return &User{}
}
