package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"time"
)

func myTask() {
	fmt.Println("This task will run periodically")
}
func executeCronJob() {
	gocron.Every(1).Second().Do(myTask)
	<- gocron.Start()
}
func main() {
	go executeCronJob()
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("dafsd")
}
