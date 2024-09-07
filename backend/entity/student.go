package entity

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	SID  	string
	Password  	string
	FirstName 	string
	LastName  	string
	Year      	uint
	Birthday  	time.Time
	Major 		string

	
	GenderID *uint
	Gender   Gender `gorm:"foriegnKey:GenderID"`

}