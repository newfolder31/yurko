package daemons

import (
	"fmt"
	"net/http"
	"yurko/interfaces"
	"yurko/usecases"
)

func Run() error {
	interfaces.Initialize()

	registrationInteractor := new(usecases.RegistrationInteractor)
	registrationInteractor.UserRepository = interfaces.UserInMemoryRepo{}

	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.RegistrationInteractor = registrationInteractor

	http.HandleFunc("/registration/fast", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.FastRegistration(res, req)
	})

	http.HandleFunc("/registration", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.Registration(res, req)
	})

	//http.Handle("/registration", registrationHandler())
	http.Handle("/login", loginHandler())
	http.Handle("/logout", logoutHandler())

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
func loginHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//if r.Method == http.MethodPost {
		interfaces.Login(w, r)
		w.WriteHeader(http.StatusOK)
		//} else {
		//	w.WriteHeader(http.StatusMethodNotAllowed)
		//}
	})
}

func logoutHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//if r.Method == http.MethodPost {
		interfaces.Logout(w, r)
		w.WriteHeader(http.StatusOK)
		//} else {
		//	w.WriteHeader(http.StatusMethodNotAllowed)
		//}
	})
}
