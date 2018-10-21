package infrastructures

import (
	"net/http"
	"time"
)

func AddCookie(w http.ResponseWriter, sessionId string) {
	cookie := &http.Cookie{
		Name:    COOKIE_SESSION_NAME,
		Value:   sessionId,
		Expires: time.Now().Add(5 * time.Minute),
	}

	http.SetCookie(w, cookie)
}

func DeleteCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:    COOKIE_SESSION_NAME,
		Value:   "",
		Expires: time.Unix(0, 0),
	}

	http.SetCookie(w, cookie)
}
