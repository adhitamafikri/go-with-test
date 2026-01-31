package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// Spawns goroutine for each `url`
	// This is a non-blocking process, because this part only spawns the goroutines
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// Collecting results from goroutines
	// Blocking process, waits for each goroutine results to arrive
	// Results are collected in order of *goroutine completion*
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}
