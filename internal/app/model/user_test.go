package model_test

import (
	"testing"

	"github.com/Harddancer/rest_api_go/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_BeforCreate(t *testing.T){
	u := model.TestUser(t)
	assert.NoError(t, u.BeforCreate())
	assert.NotEmpty(t, u.Password)

}


func TestUserValidation (t *testing.T){
	testCases := []struct{
		name string
		u func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User{
				return model.TestUser(t)
			},
			isValid: true,


		},
		{
			name: "emptyEmail",
			u: func() *model.User{
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,


		},
		{
			name: "validate encrypted Password",
			u: func() *model.User{
				u := model.TestUser(t)
				u.Password_user = ""
				u.Password = "ecrypted"
				return u
			},
			isValid: true,


		},
		{
			name: "validEmail",
			u: func() *model.User{
				u := model.TestUser(t)
				u.Email = "emailcom."
				return u
			},
			isValid: false,


		},
		{
			name: "validPassw",
			u: func() *model.User{
				u := model.TestUser(t)
				u.Password_user= "em."
				return u
			},
			isValid: false,


		},
		{
			name: "empPassw",
			u: func() *model.User{
				u := model.TestUser(t)
				u.Password_user= ""
				return u
			},
			isValid: false,


		},
	}

	for _, tc := range testCases{
		t.Run(tc.name,func(t *testing.T){
			if tc.isValid{
				assert.NoError(t,tc.u().Validate())
			}else{
				assert.Error(t,tc.u().Validate())
			}
		})
	}

}