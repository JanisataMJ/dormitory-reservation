package entity

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	RoomNumber 		uint
	Available 		string
	Confirmation 	string

	// DormID ทำหน้าที่เป็น FK
	DormID 	*uint
	Dorm	Dorm `gorm:"foriegnKey:DormID"`

	Reservations []Reservation `gorm:"foreignKey:RoomID"`
}