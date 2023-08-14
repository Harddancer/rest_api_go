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