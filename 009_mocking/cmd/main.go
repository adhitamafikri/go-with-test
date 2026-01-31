package main

import (
	"mocking"
	"os"
	"time"
)

func main() {
	sleeper := mocking.NewConfigurableSleeper(1*time.Second, time.Sleep)
	mocking.Countdown(os.Stdout, sleeper)
}
