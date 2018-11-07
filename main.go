package main

import (
	"fmt"
	"github.com/newfolder31/yurko/daemons"
	"time"
)

func main() {
	fmt.Println("Init application")
	fmt.Println(time.Now())

	//TODO start case #1
	//testUserId := uint64(1)
	//
	//intervalRepository := initTestIntervalRepository()
	//schedulerRepository := initTestSchedulerRepository()
	//a := usecases.SchedulingInteractor{IntervalRepository: intervalRepository, SchedulerRepository: schedulerRepository}
	//
	//days := make([]usecases.Day, 0, 5)
	//for i := 1; i < 6; i++ {
	//	start, _ := usecases.InitTime(uint16(i), 0)
	//	end, _ := usecases.InitTime(uint16(i), 30)
	//	timeRange, _ := usecases.InitTimeRange(start, end)
	//	day, _ := usecases.InitDay(uint8(i), []usecases.TimeRange{timeRange})
	//	days = append(days, day)
	//}
	//
	//fmt.Println(days)
	//a.CreateScheduler(testUserId, "Jurist", days)
	//fmt.Println("--------------------------------------")
	//
	//for _, i := range intervalRepository.storage {
	//	fmt.Println(i)
	//}
	//for _, i := range schedulerRepository.storage {
	//	fmt.Println(i)
	//}
	//TODO end case #1

	//TODO test code
	//var a []*Test
	//a1 := Test{a: 2}
	//a2 := Test{a: 1}
	//a3 := Test{a: 4}
	//a4 := Test{a: 3}
	//a5 := Test{a: 5}
	//a = append(a, &a1)
	//a = append(a, &a2)
	//a = append(a, &a3)
	//a = append(a, &a4)
	//a = append(a, &a5)
	//p := &a
	//sort.Slice(*p, func(i, j int) bool {
	//	return (*p)[i].a < (*p)[j].a
	//})
	//
	//fmt.Println(len(*p))
	//for _, i := range a {
	//	fmt.Println(i)
	//}\

	//b := make([]scheduling.TimeRange, 2)
	//for _, c := range b {
	//	fmt.Println(c)
	//}
	daemons.Run()
}
