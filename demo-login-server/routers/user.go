package routers

import (
	"TaskManager/controllers"

	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST", "OPTIONS")
	// logOutRouter := mux.NewRouter()
	// logOutRouter.HandleFunc("/users/logout", controllers.LogOut).Methods("DELETE", "OPTIONS")
	// router.PathPrefix("/users/logout").Handler(negroni.New(
	// 	negroni.HandlerFunc(common.Authorize),
	// 	negroni.Wrap(logOutRouter),
	// ))
	return router
}
