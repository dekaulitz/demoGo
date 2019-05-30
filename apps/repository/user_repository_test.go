package repository

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndexUsers(t *testing.T) {
	_ = assert.New(t)
	users, err := GetUserRepository().Index()
	if err != nil {
	}
	if len(users) == 0 {
		assert.True(t, len(users) == 0, "Data users not match")
	} else {
		assert.True(t, len(users) > 0, "users is empty")
	}
}

func TestStoreUsers(t *testing.T) {
	_ = assert.New(t)
	user := &UsersEntity{Name: "john", Email: "jon@email"}
	_, err := GetUserRepository().Store(user)
	if err != nil {
		println(err.Info.InternalMessage)
	}
	assert.NotEmpty(t, user.ID, "Data not stored")
}

func TestUsers_Show(t *testing.T) {
	user, err := GetUserRepository().Show("1")
	if err != nil {
		println(err.Info.Message)
	}
	assert.NotEmpty(t, user, "users empty")
}

func TestUser_Delete(t *testing.T) {
	err := GetUserRepository().Delete("7")
	println(err)
	assert.Nil(t, err, "error detected ")
}

func TestUser_Update(t *testing.T) {
	user := &UsersEntity{Name: "AJO"}
	err := GetUserRepository().Update("8", user)
	if err != nil {
		fmt.Println(err.Info.InternalMessage)
	}
}
