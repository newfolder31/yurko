package userInterfaces

import (
	"fmt"
	"github.com/gorilla/schema"
	"infrastructures"
	"net/http"
	"usecases/user"
)

func (webservice UserWebserviceHandler) Login(w http.ResponseWriter, r *http.Request) {
	//if r.Method == http.MethodPost {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	form := new(userUsecases.LoginForm)
	if err := schema.NewDecoder().Decode(form, r.Form); err != nil {
		fmt.Fprintf(w, "some error in parse request params: %s!", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if err := webservice.AuthorizationInteractor.ValidateCredentials(form); err != nil {
			fmt.Fprintf(w, "authorization is failed: %s!", err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			sessionId := infrastructures.InMemorySession.Init(form.Email)

			infrastructures.AddCookie(w, sessionId)
			w.WriteHeader(http.StatusOK)
		}
	}
	//} else {
	//	w.WriteHeader(http.StatusMethodNotAllowed)
	//}
}

func (webservice UserWebserviceHandler) Logout(w http.ResponseWriter, r *http.Request) {
	//if r.Method == http.MethodPost {

	cookie, _ := r.Cookie(infrastructures.COOKIE_SESSION_NAME)
	infrastructures.InMemorySession.Delete(cookie.Value)
	infrastructures.DeleteCookie(w)

	w.WriteHeader(http.StatusOK)
	//} else {
	//	w.WriteHeader(http.StatusMethodNotAllowed)
	//}
}
