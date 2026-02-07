package web_racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"web_racer/test_utils"
)

func TestWebRacer(t *testing.T) {
	// Doing actual call to the external service is slow and produces unpredictable result
	// slowURL := "https://dev.to/"
	// fastURL := "https://www.detik.com/"

	// expected := fastURL
	// result := Racer(slowURL, fastURL)

	// More predictable result and faster test execution with mocked server approach
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(15 * time.Millisecond)

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	expected := fastURL
	result := Racer(slowURL, fastURL)

	test_utils.AssertValue(t, result, expected)

	slowServer.Close()
	fastServer.Close()
}

func TestRacerWithSelect(t *testing.T) {
	t.Run("compare speeds of the server, returning the URL of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(5 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		expected := fastURL
		result, err := RacerWithSelect(slowURL, fastURL)

		test_utils.AssertValue(t, result, expected)
		test_utils.AssertNoError(t, err)
	})

	t.Run("returns an error if a server does not respond within 10s", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(15 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		_, err := ConfigurableRacer(slowURL, fastURL, 10*time.Millisecond)

		test_utils.AssertError(t, err, ServerTimeoutError)
	})
}

func makeDelayedServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}
