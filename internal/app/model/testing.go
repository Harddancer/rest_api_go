package model

import "testing"

func TestUser(t *testing.T) *User{
	return &User{
	Email: "user@example.org",
	Password_user: "Pass123",
}
}