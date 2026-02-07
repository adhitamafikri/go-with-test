package web_racer

import (
	"errors"
	"net/http"
	"time"
)

var ServerTimeoutError = errors.New("Timeout, Request takes too long")
var tenSecondTimeout = 10 * time.Second

func Racer(URL1 string, URL2 string) (winner string) {
	aDuration := measureResponseTime(URL1)
	bDuration := measureResponseTime(URL2)

	if aDuration < bDuration {
		return URL1
	}

	return URL2
}

func measureResponseTime(URL string) time.Duration {
	start := time.Now()
	http.Get(URL)
	duration := time.Since(start)

	return duration
}

func RacerWithSelect(url1 string, url2 string) (winner string, err error) {
	return ConfigurableRacer(url1, url2, tenSecondTimeout)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", ServerTimeoutError
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
