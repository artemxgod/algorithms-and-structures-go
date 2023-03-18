package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

// Concurrency is the way to architect a program.
// It allows program to do multiple tasks at the 'same' time

// Goroutines are controlled by the runtime,
// so it changes the size of the goroutine dynamically

func PseudoParseUrl(url string) {
	for i := 0; i < 5; i++ {
		latency := rand.Intn(500) + 500
		time.Sleep(time.Millisecond * time.Duration(latency))
		fmt.Printf("Parsing <%s> ## Step %d ## Latency %d\n", url, i+1, latency)
	}
}

func TestPseudoParseUrl() {
	t := time.Now()
	rand.Seed(t.UnixNano())
	PseudoParseUrl("google.com")
	PseudoParseUrl("youtube.com")
	fmt.Printf("Parse completed. Time %.2f sec.\n", time.Since(t).Seconds())
}

// Less time
func TestConcurrencyParse() {
	t := time.Now()
	rand.Seed(t.UnixNano())
	go PseudoParseUrl("google.com")
	PseudoParseUrl("youtube.com")
	fmt.Printf("Parse completed. Time %.2f sec.\n", time.Since(t).Seconds())
}