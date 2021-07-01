package controllers

import "TaskManager/models"

type (
	UserResource struct {
		Data models.User `json:"data"`
	}
	LoginModel struct {
		Email    string `json:"email`
		Password string `json:"password"`
	}
	LoginResource struct {
		Data LoginModel `json:"data"`
	}
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}
)
