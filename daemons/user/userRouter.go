package user

import (
	"github.com/go-chi/chi"
	userWebHandler "github.com/newfolder31/yurko/interfaces/user/handlers"
	userRepository "github.com/newfolder31/yurko/interfaces/user/repositories"
	userUsecases "github.com/newfolder31/yurko/usecases/user"
	"net/http"
)

func InitUserModule(r *chi.Mux) {

	//initialize repositories
	userInMemoryRepo := userRepository.NewUserInMemoryRepo()
	addressInMemoryRepo := userRepository.NewAddressInMemoryRepo()

	//initialize db repositories
	//var postgresHandler = infrastructures.NewPostgresHandler()
	//userInDbRepo := interfaces.UserInDbRepo{DbHandler: postgresHandler}

	//initialize interactors
	registrationInteractor := new(userUsecases.RegistrationInteractor)
	registrationInteractor.UserRepository = userInMemoryRepo
	registrationInteractor.AddressRepository = addressInMemoryRepo

	authorizationInteractor := new(userUsecases.AuthorizationInteractor)
	authorizationInteractor.UserRepository = userInMemoryRepo

	profileInteractor := new(userUsecases.ProfileInteractor)
	profileInteractor.UserRepository = userInMemoryRepo
	profileInteractor.AddressRepository = addressInMemoryRepo

	//initialize webservices
	webserviceHandler := userWebHandler.UserWebserviceHandler{}
	webserviceHandler.RegistrationInteractor = registrationInteractor
	webserviceHandler.AuthorizationInteractor = authorizationInteractor
	webserviceHandler.ProfileInteractor = profileInteractor

	//set api handlers
	r.Post("/api/v0/registration/fast", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.FastRegistration(res, req)
	})

	r.Post("/api/v0/registration", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.Registration(res, req)
	})

	r.Post("/api/v0/login", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.Login(res, req)
	})

	r.Post("/api/v0/logout", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.Logout(res, req)
	})

	r.Get("/api/v0/profile/get", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.GetUser(res, req)
	})

	r.Post("/api/v0/profile/update", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.UpdateUser(res, req)
	})
}
