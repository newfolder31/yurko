package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/newfolder31/yurko/infrastructures"
	userUsecases "github.com/newfolder31/yurko/usecases/user"
	"net/http"
)

/**
http handlers for:
- /login
- /logout
*/
func (webservice UserWebserviceHandler) Login(w http.ResponseWriter, r *http.Request) {
	form := new(userUsecases.LoginForm)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&form); err != nil {
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
}

func (webservice UserWebserviceHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie(infrastructures.COOKIE_SESSION_NAME)
	infrastructures.InMemorySession.Delete(cookie.Value)
	infrastructures.DeleteCookie(w)

	w.WriteHeader(http.StatusOK)
}
