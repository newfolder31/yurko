package daemons

import (
	"fmt"
	"net/http"
	"yurko/interfaces"
	"yurko/usecases"
)

func Run() error {
	interfaces.Initialize()
	userInMemoryRepo := interfaces.NewUserInMemoryRepo()

	registrationInteractor := new(usecases.RegistrationInteractor)
	registrationInteractor.UserRepository = userInMemoryRepo

	authorizationInteractor := new(usecases.AuthorizationInteractor)
	authorizationInteractor.UserRepository = userInMemoryRepo

	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.RegistrationInteractor = registrationInteractor
	webserviceHandler.AuthorizationInteractor = authorizationInteractor

	http.HandleFunc("/registration/fast", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.FastRegistration(res, req)
	})

	http.HandleFunc("/registration", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.Registration(res, req)
	})

	http.HandleFunc("/login", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.Login(res, req)
	})

	http.HandleFunc("/logout", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.Logout(res, req)
	})

	http.Handle("/", indexHandler())

	http.ListenAndServe(":8081", nil)

	return nil
}

func indexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		cookie, _ := r.Cookie("sessionId")
		if cookie != nil {
			userEmail := interfaces.GetCurrentUser(cookie.Value)
			fmt.Fprintf(w, "Current user is %s\n", userEmail)
		}
	})
}
