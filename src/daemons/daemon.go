package daemons

import "net/http"

func Run() error {

	initScheduleHandling()

	http.ListenAndServe(":8081", nil)
}

func initScheduleHandling() {

}
