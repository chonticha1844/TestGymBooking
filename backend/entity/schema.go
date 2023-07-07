package entity

import (
	//"regexp"

	//"github.com/asaskevich/govalidator"
	"time"

	"gorm.io/gorm"
)

// Equipment
type Equipment struct {
	gorm.Model
	Equipments string

	Reservation []Reservation `gorm:"foreignKey:EquipmentID"`
}

// Gender
type Gender struct {
	gorm.Model
	Gender string

	User []User `gorm:"foreignKey:GenderID"`
}

// User
type User struct {
	gorm.Model
	Username string
	Gmail    string
	Password string
	Fullname string
	Age      int32
	Weight   int32
	Height   int32

	GenderID *uint
	Gender   Gender

	Reservation []Reservation `gorm:"foreignKey:UserID"`
}

// Reservation
type Reservation struct {
	gorm.Model
	Datetime time.Time

	UserID      *uint
	User        User
	EquipmentID *uint
	Equipment   Equipment
}
