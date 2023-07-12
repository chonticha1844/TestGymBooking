package entity

import (
	"regexp"

	"time"

	"github.com/asaskevich/govalidator"

	"gorm.io/gorm"
)

// Equipment
type Equipment struct {
	gorm.Model
	Equipments string
	picture    string

	Reservation []Reservation `gorm:"foreignKey:EquipmentID" valid:"-"`
}

// Gender
type Gender struct {
	gorm.Model
	Gender string

	User []User `gorm:"foreignKey:GenderID" valid:"-"`
}

// User
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex" valid:"required~กรุณากรอก username,matches(^(B|D|M)([0-9]{7}$))~username ต้องมี 8 ตัว" `
	Email    string `gorm:"uniqueIndex" valid:"email~รูปแบบอีเมล์ไม่ถูกต้อง,required~กรุณากรอกอีเมล์"`
	Password string `valid:"required~กรุณากรอกรหัสผ่าน,matches(^[1-9]([0-9]{12}$))~password ต้องมี 13 ตัว" `
	Fullname string `valid:"required~กรุณากรอกชื่อ-นามสกุล"`
	Age      int32
	Weight   int32
	Height   int32

	GenderID *uint
	Gender   Gender `gorm:"references:id" valid:"-"`

	Reservation []Reservation `gorm:"foreignKey:UserID" valid:"-"`
}

// Reservation
type Reservation struct {
	gorm.Model
	Datetime time.Time

	UserID      *uint
	User        User `gorm:"references:id" valid:"-"`
	EquipmentID *uint
	Equipment   Equipment `gorm:"references:id" valid:"-"`
}

// ฟังก์ชันที่จะใช่ในการ validation ตัวอักษรพิเศษและตัวเลข
func init() {
	govalidator.CustomTypeTagMap.Set("checkuserpattern", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		s, ok := i.(string)
		if !ok {
			return false
		}
		match, _ := regexp.MatchString("^[ก-๛a-zA-Z\\s]+$", s)
		return match
	}))
}
