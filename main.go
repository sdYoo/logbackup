package main

import (
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
	"logbackup/services"
)

func secondTask() {
	fmt.Println("Task:", time.Now())
	services.GetS3Sync()
}

func hourTask() {
	fmt.Println("Hour Task:", time.Now())
}

func main() {
	fmt.Println("start main!!")
	gocron.Every(1).Second().Do(secondTask)

	// Begin job immediately upon start
	gocron.Every(1).Hour().From(gocron.NextTick()).Do(hourTask)

	//gocron.Every(1).Hour().Do(hourTask)

	// Start all the pending jobs
	<-gocron.Start()

	// also, you can create a new scheduler
	// to run two schedulers concurrently
	//s := gocron.NewScheduler()
	//<-s.Start()

	fmt.Println("end main!!")
}
