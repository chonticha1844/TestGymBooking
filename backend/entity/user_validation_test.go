package entity

import (
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestWriterCorrect(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Check format user", func(t *testing.T) {
		user := User{

			Username: "B1234567",
			Email:    "B1234567@g.sut.ac.th",
			Password: "1234567890123",
			Fullname: "Tom Highway",
		}
		//ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(user)

		//เช็คว่ามันเป็นค่าจริงไหม
		g.Expect(ok).To(BeTrue())

		//เช็คว่ามันว่างไหม
		g.Expect(err).To((BeNil()))

		fmt.Println(err)
	})
}

func TestUser(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Check username not blank ", func(t *testing.T) {

		user := User{

			Username: "",
			Email:    "B1234567@g.sut.ac.th",
			Password: "1234567890123",
			Fullname: "Tom Highway",
		}
		//ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณากรอก username"))
	})

	t.Run("Check username must be 8 digit ", func(t *testing.T) {

		user := User{

			Username: "k12345",
			Email:    "B1234567@g.sut.ac.th",
			Password: "1234567890123",
			Fullname: "Tom Highway",
		}
		//ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("username ต้องมี 8 ตัว"))
	})

	t.Run("Check email not blank ", func(t *testing.T) {

		user := User{

			Username: "B1234567",
			Email:    "",
			Password: "1234567890123",
			Fullname: "Tom Highway",
		}
		//ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณากรอกอีเมล์"))
	})

	t.Run("Check email pattern correct ", func(t *testing.T) {

		user := User{

			Username: "B1234567",
			Email:    "B1234567@",
			Password: "1234567890123",
			Fullname: "Tom Highway",
		}
		//ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("รูปแบบอีเมล์ไม่ถูกต้อง"))
	})

	t.Run("Check password not blank ", func(t *testing.T) {

		user := User{

			Username: "B1234567",
			Email:    "B1234567@g.sut.ac.th",
			Password: "",
			Fullname: "Tom Highway",
		}
		//ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณากรอกรหัสผ่าน"))
	})

	t.Run("Check password must be =13 digit", func(t *testing.T) {

		user := User{

			Username: "B1234567",
			Email:    "B1234567@g.sut.ac.th",
			Password: "12345",
			Fullname: "Tom Highway",
		}
		//ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("password ต้องมี 13 ตัว"))
	})

	t.Run("Check fullname not blank ", func(t *testing.T) {

		user := User{

			Username: "B1234567",
			Email:    "B1234567@g.sut.ac.th",
			Password: "1234567890123",
			Fullname: "",
		}
		//ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณากรอกชื่อ-นามสกุล"))
	})
}
