package service

import (
	"demoGo/apps/handler"
	"demoGo/apps/handler/exception"
	"demoGo/apps/models"
)

//example for registration service
func Register(user *models.Users) (*models.Users, *exception.ErrorException) {
	err := handler.ValidateRequest(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
