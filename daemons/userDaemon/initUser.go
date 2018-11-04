package userDaemon

import (
	"github.com/newfolder31/yurko/interfaces/repositories/userRepository"
	"github.com/newfolder31/yurko/interfaces/webservices/userWebservice"
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

	//initialize webservices
	webserviceHandler := userWebservice.UserWebserviceHandler{}
	webserviceHandler.RegistrationInteractor = registrationInteractor
	webserviceHandler.AuthorizationInteractor = authorizationInteractor

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
