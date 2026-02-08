package main

import (
	"github.com/adhitamafikri/go-with-test/mocking"
	"os"
	"time"
)

func main() {
	sleeper := mocking.NewConfigurableSleeper(1*time.Second, time.Sleep)
	mocking.Countdown(os.Stdout, sleeper)
}
