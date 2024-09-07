package config

import (
	"fmt"
	"time"

	"github.com/tanapon395/sa-67-example/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {
	db.AutoMigrate(
		&entity.Student{},
		&entity.Gender{},
		&entity.Dorm{},
		&entity.Room{},
		&entity.Reservation{},
	)

	GenderMale := entity.Gender{Gender_type: "Male"}
	GenderFemale := entity.Gender{Gender_type: "Female"}

	db.FirstOrCreate(&GenderMale, &entity.Gender{Gender_type: "Male"})
	db.FirstOrCreate(&GenderFemale, &entity.Gender{Gender_type: "Female"})

	// Hash the password and handle errors
	hashedPassword, err := HashPassword("123456")
	if err != nil {
		panic(fmt.Sprintf("failed to hash password: %v", err))
	}

	// Parse the birthday and handle errors
	BirthDay, err := time.Parse("2006-01-02", "1988-11-12")
	if err != nil {
		panic(fmt.Sprintf("failed to parse birthdate: %v", err))
	}
	gid := uint(1)
	User := &entity.Student{
		SID:       "A1234567",
		FirstName: "Software",
		LastName:  "Analysis",
		Year:      1, // Changed to uint type
		Major:     "Engineering",
		Password:  hashedPassword,
		Birthday:  BirthDay, // Use Birthday instead of BirthDay
		GenderID:  &gid, // Make sure GenderID is of type *uint
	}

	db.FirstOrCreate(User, &entity.Student{
		SID: "A1234567",
	})
}

