package main

import (
	"fmt"
	"time"
)

func main() {
	sleepDuration := 2 * time.Second
	now := time.Now()

	fmt.Printf("Sleeping for %v...\n", sleepDuration)
	time.Sleep(sleepDuration)
	fmt.Printf("Awake!\n")

	duration := time.Since(now)
	fmt.Printf("Time elapsed since start: %v\n", duration)

	oneHourLater := now.Add(time.Hour)
	fmt.Printf("One hour later: %v\n", oneHourLater)

	oneDayAgo := now.AddDate(0, 0, -1)
	fmt.Printf("One day ago: %v\n", oneDayAgo)

	timer := time.NewTimer(3 * time.Second)
	<-timer.C
	fmt.Printf("Timer expired!\n")

	ticker := time.NewTicker(1 * time.Second)
	for i := 0; i < 3; i++ {
		<-ticker.C
		fmt.Printf("Tick %d\n", i+1)
	}
	ticker.Stop()

	militaryTimeStr := "15:04:05"

	parsedTime, err := time.Parse("15:04:05", militaryTimeStr)
	if err != nil {
		fmt.Printf("Error parsing military time: %v\n", err)
		return
	}

	fmt.Printf("Parsed military time: %s\n", parsedTime.Format("15:04:05"))
}
