package main

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteCheck(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	expectation := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	output := CheckWebsites(mockWebsiteCheck, websites)

	if !reflect.DeepEqual(expectation, output) {
		t.Fatalf("expected %v, output %v", expectation, output)
	}
}

func slowStubCheckWebsites(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubCheckWebsites, urls)
	}
}
