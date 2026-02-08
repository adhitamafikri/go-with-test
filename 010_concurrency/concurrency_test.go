package concurrency

import (
	"github.com/adhitamafikri/go-with-test/concurrency/test_utils"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://bababazinga.bro"
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"https://facebook.com",
		"https://x.com",
		"https://google.com",
		"waat://bababazinga.bro",
	}

	expected := map[string]bool{
		"https://facebook.com":   true,
		"https://x.com":          true,
		"https://google.com":     true,
		"waat://bababazinga.bro": false,
	}

	result := CheckWebsites(mockWebsiteChecker, websites)

	test_utils.AssertValueDeepEqual(t, expected, result)
}
