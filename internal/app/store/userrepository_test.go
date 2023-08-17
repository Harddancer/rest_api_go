package store_test

import (
	"fmt"
	"testing"

	"github.com/Harddancer/rest_api_go/internal/app/model"
	"github.com/Harddancer/rest_api_go/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T, ){
	s, teardown := store.TestStore(t, databaseUrl)
	fmt.Printf("store %v",s)
	defer teardown("users")

	u,err := s.User().Create(model.TestUser((t)))

	fmt.Print(u)

	assert.NoError(t, err)
	assert.NotNil(t,u)
}

func TestUserRepository_FindByEmail(t *testing.T,){
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	email := "user123457@email.com"

	_,err := s.User().FindByEmail(email)
	assert.Error(t,err)

	

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)

	u,err = s.User().FindByEmail(email)
	assert.NoError(t,err)
	assert.NotNil(t,u)


}



func TestUserRepository_FindById(t *testing.T,){
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	id,_:= s.User().Create(model.TestUser((t)))

	u,err := s.User().FindById(id.ID)
	fmt.Print(u)
	assert.NoError(t,err)
	assert.NotNil(t,u)


}