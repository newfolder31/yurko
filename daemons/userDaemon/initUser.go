package userDaemon

import (
	userRepository "github.com/newfolder31/yurko/interfaces/user/repositories"
	userWebHandler "github.com/newfolder31/yurko/interfaces/user/webHandlers"
	"github.com/newfolder31/yurko/usecases/userUsecases"
	"net/http"
)

func InitAuthModule() {
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
}
