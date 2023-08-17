package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct{
	ID int
	Email string
	Password_user string
	Password string
}

// 
func (u *User) BeforCreate() error{
	if len(u.Password_user) > 0 {
		enc,err := encryptString(u.Password_user)
		if err != nil{
			return err
		}
		u.Password = enc
	}
	return nil
}

func encryptString (s string) (string,error){
	b,err := bcrypt.GenerateFromPassword([]byte(s),bcrypt.MinCost)
	if err != nil{
		return "",err
	}
return string(b),nil

}

func (u *User) Validate() error{
	return validation.ValidateStruct(
		u, 
		validation.Field(&u.Email, validation.Required,is.Email),
		validation.Field(&u.Password_user,validation.By(requiredIf(u.Password == "")), validation.Length(6,100)),
	)
}