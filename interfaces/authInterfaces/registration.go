package authInterfaces

import (
	"fmt"
	"github.com/gorilla/schema"
	"net/http"
	"yurko/usecases/authUsecases"
)

func (webservice WebserviceHandler) FastRegistration(w http.ResponseWriter, r *http.Request) {
	//if r.Method == http.MethodPost {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	form := new(authUsecases.RegistrationForm)
	if err := schema.NewDecoder().Decode(form, r.Form); err != nil {
		fmt.Fprintf(w, "some error in parse request params: %s!", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if err := webservice.RegistrationInteractor.ValidateFastRegistrationRequest(form); err != nil {
			fmt.Fprintf(w, "some error in request params: %s!", err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			webservice.RegistrationInteractor.Registration(form)
			w.WriteHeader(http.StatusOK)
		}
	}
	//} else {
	//	w.WriteHeader(http.StatusMethodNotAllowed)
	//}
}

func (webservice WebserviceHandler) Registration(w http.ResponseWriter, r *http.Request) {
	//if r.Method == http.MethodPost {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	form := new(authUsecases.RegistrationForm)
	if err := schema.NewDecoder().Decode(form, r.Form); err != nil {
		fmt.Fprintf(w, "some error in parse request params: %s!", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if err := webservice.RegistrationInteractor.ValidateRegistrationRequest(form); err != nil {
			fmt.Fprintf(w, "some error in request params: %s!", err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			webservice.RegistrationInteractor.Registration(form)
			w.WriteHeader(http.StatusOK)
		}
	}
	//} else {
	//	w.WriteHeader(http.StatusMethodNotAllowed)
	//}
}
