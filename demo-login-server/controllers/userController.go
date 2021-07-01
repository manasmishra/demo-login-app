package controllers

import (
	"encoding/json"
	"net/http"

	"TaskManager/common"
	"TaskManager/data"
	"TaskManager/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	// common.SetCors(&w)
	var dataSource UserResource

	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User Data",
			500,
		)
		return
	}
	user := &dataSource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}
	repo.CreateUser(user)
	user.HashedPassword = ""
	if j, err := json.Marshal(UserResource{Data: *user}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An Unexpected error occured",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

// Handlers for HTTP Post - users Login
func Login(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	var dataSource LoginResource
	var token string
	err := json.NewDecoder(r.Body).Decode(&dataSource)

	if err != nil {
		// err := errors.New("Invalid Login Data")
		common.DisplayAppError(
			w,
			err,
			"Invalid Login Data",
			500,
		)
		return
	}
	loginModel := dataSource.Data
	loginUser := models.User{
		Email:    loginModel.Email,
		Password: loginModel.Password,
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}
	if user, err := repo.Login(loginUser); err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid login Credentials",
			401,
		)
		return
	} else { // Login successful
		// Generate JWT Token
		token, err = common.GenerateJWT(user.Email, "member")
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"Error while generating the Access Token",
				500,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		user.HashedPassword = ""
		authUser := AuthUserModel{
			User:  user,
			Token: token,
		}
		j, err := json.Marshal(AuthUserResource{Data: authUser})
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"An Unexpected Error Occured",
				500,
			)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
		http.Redirect(w, r, "/users/logout", 302)
	}
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully Loged out"))
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
