package userWebservice

import (
	"fmt"
	"github.com/gorilla/schema"
	"infrastructures"
	"net/http"
	"usecases/user"
)

func (webservice UserWebserviceHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		form := new(userUsecases.LoginForm)
		if err := schema.NewDecoder().Decode(form, r.Form); err != nil {
			fmt.Fprintf(w, "some error in parse request params: %s!", err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			cookie, err := r.Cookie(infrastructures.COOKIE_SESSION_NAME)

			if cookie.Value == "" || err != nil {
				fmt.Fprintf(w, "authorization is failed: %s!", err)
				w.WriteHeader(http.StatusUnauthorized)
			} else {

				var email = infrastructures.InMemorySession.Get(cookie.Value)
				var user, _ = webservice.ProfileInteractor.GetUser(email)

				fmt.Fprint(w, user)

				w.WriteHeader(http.StatusOK)
			}
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
