package entity

import (
	//"time"

	//"regexp"

	//"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// ---Education---
type Education struct {
	gorm.Model
	Education_degree string
}
