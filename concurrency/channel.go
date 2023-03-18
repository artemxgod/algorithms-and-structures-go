package concurrency

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// non-buffered channels blocks goroutine after write/read
// buffered channels blocks goroutine if buffer is full

func writeNumbers() {
	message := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			message <- fmt.Sprintf("%d", i+1)
			time.Sleep(time.Millisecond * 250)
		}
		// close channel to avoid deadlock
		close(message)
	}()

	for msg := range message {
		fmt.Println(msg)
	}
}

func doHttp(url string) {
	t := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to get", url, ":", err.Error())
	} else {
		fmt.Println(url, "- Status code:", resp.StatusCode, "Latency - ", time.Since(t).Milliseconds(), "ms.")
	}
	defer resp.Body.Close()

}

func Selection() {
	message1, message2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			message1 <- "Half a second pass"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Millisecond * 1000)
			message1 <- "A second pass"
		}
	}()

	// select allow us to read from two channels without blocking a goroutine
	for {
		select {
		case msg := <-message1:
			fmt.Println(msg)
		case msg := <-message2:
			fmt.Println(msg)
		}
	}
}

func TestWriteNumbers() {
	writeNumbers()
}

func TestDoHttp() {
	urls := []string{
		"https://google.com/",
		"https://youtube.com/",
		"https://medium.com/",
		"https://github.com/",
		"https://t.me/",
	}

	wg := sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go func(url string){
			doHttp(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func TestSelection() {
	Selection()
}

