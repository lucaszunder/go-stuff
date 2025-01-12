package main

type WebsiteCheck func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteCheck, urls []string) map[string]bool {
	results := make(map[string]bool)
	outputChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			outputChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-outputChannel
		results[result.string] = result.bool
	}

	return results

}
