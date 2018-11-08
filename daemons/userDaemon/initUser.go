package userDaemon

import (
	"github.com/go-chi/chi"
	userRepository "github.com/newfolder31/yurko/interfaces/user/repositories"
	userWebHandler "github.com/newfolder31/yurko/interfaces/user/webHandlers"
	"github.com/newfolder31/yurko/usecases/userUsecases"
	"net/http"
)

func InitUserModule(r *chi.Mux) {

	//initialize repositories
	userInMemoryRepo := userRepository.NewUserInMemoryRepo()

	//initialize db repositories
	//var postgresHandler = infrastructures.NewPostgresHandler()
	//userInDbRepo := interfaces.UserInDbRepo{DbHandler: postgresHandler}

	//initialize interceptors
	registrationInteractor := new(userUsecases.RegistrationInteractor)
	registrationInteractor.UserRepository = userInMemoryRepo

	authorizationInteractor := new(userUsecases.AuthorizationInteractor)
	authorizationInteractor.UserRepository = userInMemoryRepo

	profileInteractor := new(userUsecases.ProfileInteractor)
	profileInteractor.UserRepository = userInMemoryRepo

	//initialize webservices
	webserviceHandler := userWebHandler.UserWebserviceHandler{}
	webserviceHandler.RegistrationInteractor = registrationInteractor
	webserviceHandler.AuthorizationInteractor = authorizationInteractor
	webserviceHandler.ProfileInteractor = profileInteractor

	//set api handlers todo question: prefix "api" ?
	r.Post("/registration/fast", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.FastRegistration(res, req)
	})
	//TODO old version
	//http.HandleFunc("/registration/fast", func(res http.ResponseWriter, req *http.Request) {
	//	webserviceHandler.FastRegistration(res, req)
	//})

	r.Post("/registration", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.Registration(res, req)
	})
	//TODO old version
	//http.HandleFunc("/registration", func(res http.ResponseWriter, req *http.Request) {
	//	webserviceHandler.Registration(res, req)
	//})

	r.Post("/login", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.Login(res, req)
	})
	//TODO old version
	//http.HandleFunc("/login", func(res http.ResponseWriter, req *http.Request) {
	//	webserviceHandler.Login(res, req)
	//})

	r.Post("/logout", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.Logout(res, req)
	})
	//TODO old version
	//http.HandleFunc("/logout", func(res http.ResponseWriter, req *http.Request) {
	//	webserviceHandler.Logout(res, req)
	//})

	r.Get("/profile/get", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.GetUser(res, req)
	})
	//TODO old version
	//http.HandleFunc("/profile/get", func(res http.ResponseWriter, req *http.Request) {
	//	webserviceHandler.GetUser(res, req)
	//})

	r.Post("/profile/update", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.UpdateUser(res, req)
	})
	//TODO old version
	//http.HandleFunc("/profile/update", func(res http.ResponseWriter, req *http.Request) {
	//	webserviceHandler.UpdateUser(res, req)
	//})
}
