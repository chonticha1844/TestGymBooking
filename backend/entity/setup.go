package entity

import (
	//"fmt"
	//"go/format"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("project.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database.AutoMigrate(

		&Equipment{},
		&Gender{},
		&User{},
		&Reservation{},
	)

	db = database

	//Equipment
	torso_twist := Equipment{
		Equipments: "Torso Twist",
	}
	db.Model(&Equipment{}).Create(&torso_twist)

	ABD := Equipment{
		Equipments: "ABD.Machine",
	}
	db.Model(&Equipment{}).Create(&ABD)

	leg_press := Equipment{
		Equipments: "Leg Press",
	}
	db.Model(&Equipment{}).Create(&leg_press)

	leg_extension := Equipment{
		Equipments: "Leg Extension",
	}
	db.Model(&Equipment{}).Create(&leg_extension)

	leg_curl := Equipment{
		Equipments: "Leg Curl",
	}
	db.Model(&Equipment{}).Create(&leg_curl)

	calf_raise := Equipment{
		Equipments: "Calf Raise",
	}
	db.Model(&Equipment{}).Create(&calf_raise)

	in_thing := Equipment{
		Equipments: "Inner Thigh",
	}
	db.Model(&Equipment{}).Create(&in_thing)

	o_thing := Equipment{
		Equipments: "Outer Thigh",
	}
	db.Model(&Equipment{}).Create(&o_thing)

	pec_deck := Equipment{
		Equipments: "Pec Deck",
	}
	db.Model(&Equipment{}).Create(&pec_deck)

	in_press := Equipment{
		Equipments: "Incline Press",
	}
	db.Model(&Equipment{}).Create(&in_press)

	chest_press := Equipment{
		Equipments: "Chest Press",
	}
	db.Model(&Equipment{}).Create(&chest_press)

	sh_press := Equipment{
		Equipments: "Shoulder Press",
	}
	db.Model(&Equipment{}).Create(&sh_press)

	lat_pull := Equipment{
		Equipments: "Lat Pulldown",
	}
	db.Model(&Equipment{}).Create(&lat_pull)

	seated_row := Equipment{
		Equipments: "Seated Row",
	}
	db.Model(&Equipment{}).Create(&seated_row)

	back_ex := Equipment{
		Equipments: "Back Extention",
	}
	db.Model(&Equipment{}).Create(&back_ex)

	lat_row := Equipment{
		Equipments: "Lat/High Row",
	}
	db.Model(&Equipment{}).Create(&lat_row)

	arm_curl := Equipment{
		Equipments: "Arm Curl",
	}
	db.Model(&Equipment{}).Create(&arm_curl)

	tri_ex := Equipment{
		Equipments: "Triceps Extension",
	}
	db.Model(&Equipment{}).Create(&tri_ex)

	//gender
	male := Gender{
		Gender: "Male",
	}
	db.Model(&Gender{}).Create(&male)

	female := Gender{
		Gender: "Female",
	}
	db.Model(&Gender{}).Create(&female)

	//User
	user1 := User{
		Username: "B1234567",
		Gmail:    "B1234567@g.sut.ac.th",
		Password: "1234567890123",
		Fullname: "Tom Highway",
		Age:      21,
		Weight:   70,
		Height:   182,

		Gender: male,
	}
	db.Model(&User{}).Create(&user1)

	user2 := User{
		Username: "D1472583",
		Gmail:    "D1472583@g.sut.ac.th",
		Password: "2345678901234",
		Fullname: "Malisa Somalia",
		Age:      30,
		Weight:   56,
		Height:   160,

		Gender: female,
	}
	db.Model(&User{}).Create(&user2)

	//Reservation
	resv1 := Reservation{
		User:      user1,
		Datetime:  time.Date(2023, 5, 10, 15, 30, 00, 00, time.Now().Local().Location()),
		Equipment: leg_press,
	}
	db.Model(&Reservation{}).Create(&resv1)

	resv2 := Reservation{
		User:      user2,
		Datetime:  time.Date(2023, 5, 10, 16, 00, 00, 00, time.Now().Local().Location()),
		Equipment: leg_extension,
	}
	db.Model(&Reservation{}).Create(&resv2)
}
