package interfaces

import (
	"fmt"
	"net/http"
	"yurko/infrastructures"
	"yurko/usecases"
)

func Login(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	pass := r.FormValue("password") //todo: encode pass
	fmt.Println(email, pass)

	if err := usecases.ValidateLogin(email, pass); err != nil {
		return err
	} else {
		sessionId := inMemorySession.Init(email)

		infrastructures.AddCookie(w, sessionId)
	}
	return nil
}

func Logout(w http.ResponseWriter, r *http.Request) error {
	cookie, _ := r.Cookie("sessionId")

	inMemorySession.Delete(cookie.Value)

	infrastructures.DeleteCookie(w)

	return nil
}
