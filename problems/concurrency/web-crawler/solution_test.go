package webcrawler

import (
	"errors"
	"testing"
)

// fakeFetcher is a mock Fetcher for testing
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, errors.New("not found")
}

var testFetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/cmd/": &fakeResult{
		"Commands",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func TestCrawlMutex(t *testing.T) {
	visited := CrawlMutex("https://golang.org/", 2, testFetcher)

	expected := []string{
		"https://golang.org/",
		"https://golang.org/pkg/",
		"https://golang.org/cmd/",
	}

	for _, url := range expected {
		if !visited[url] {
			t.Errorf("Expected %s to be visited", url)
		}
	}

	if len(visited) < len(expected) {
		t.Errorf("Expected at least %d URLs to be visited, got %d", len(expected), len(visited))
	}
}

func TestCrawlSequential(t *testing.T) {
	visited := make(map[string]bool)
	CrawlSequential("https://golang.org/", 2, testFetcher, visited)

	expected := []string{
		"https://golang.org/",
		"https://golang.org/pkg/",
		"https://golang.org/cmd/",
	}

	for _, url := range expected {
		if !visited[url] {
			t.Errorf("Expected %s to be visited", url)
		}
	}
}

func TestCrawlDepthZero(t *testing.T) {
	visited := CrawlMutex("https://golang.org/", 0, testFetcher)

	if len(visited) != 0 {
		t.Errorf("Expected no URLs to be visited with depth 0, got %d", len(visited))
	}
}

func TestCrawlInvalidURL(t *testing.T) {
	visited := CrawlMutex("https://invalid.url/", 1, testFetcher)

	// The invalid URL is marked as visited even though fetch fails
	// This is expected behavior - we try to visit it, then fail
	if len(visited) > 1 {
		t.Errorf("Expected at most 1 URL to be visited for invalid URL, got %d", len(visited))
	}
}

func TestCrawlCircularLinks(t *testing.T) {
	// Test that circular links don't cause infinite loops
	fetcher := fakeFetcher{
		"https://a.com/": &fakeResult{
			"Page A",
			[]string{"https://b.com/"},
		},
		"https://b.com/": &fakeResult{
			"Page B",
			[]string{"https://a.com/"},
		},
	}

	visited := CrawlMutex("https://a.com/", 5, fetcher)

	if len(visited) != 2 {
		t.Errorf("Expected 2 URLs to be visited, got %d", len(visited))
	}
}

// Benchmark tests
func BenchmarkCrawlMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CrawlMutex("https://golang.org/", 2, testFetcher)
	}
}

func BenchmarkCrawlSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		visited := make(map[string]bool)
		CrawlSequential("https://golang.org/", 2, testFetcher, visited)
	}
}

func BenchmarkCrawlDeep(b *testing.B) {
	// Create a deep link structure
	deepFetcher := make(fakeFetcher)
	for i := 0; i < 20; i++ {
		url := "https://example.com/page" + string(rune('0'+i))
		nextURL := "https://example.com/page" + string(rune('0'+i+1))
		deepFetcher[url] = &fakeResult{
			"Page " + string(rune('0'+i)),
			[]string{nextURL},
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CrawlMutex("https://example.com/page0", 10, deepFetcher)
	}
}

func BenchmarkCrawlWide(b *testing.B) {
	// Create a wide link structure
	wideFetcher := make(fakeFetcher)
	root := "https://example.com/"
	urls := make([]string, 50)
	for i := 0; i < 50; i++ {
		urls[i] = "https://example.com/page" + string(rune('0'+i))
		wideFetcher[urls[i]] = &fakeResult{"Page " + string(rune('0'+i)), nil}
	}
	wideFetcher[root] = &fakeResult{"Root", urls}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CrawlMutex(root, 2, wideFetcher)
	}
}
