package packages

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// -- CONTEXT IN REQUESTS --

// A Context carries deadlines, cancellation signals,
// and other request-scoped values across API boundaries and goroutines.
func hello(w http.ResponseWriter, r *http.Request) {
	// A context.Context is created for each request by the net/http machinery, and is available with the Context() method.
	ctx := r.Context()
	fmt.Println("hello handler started")
	defer fmt.Println("hello handler ended")

	select {
	case <-time.After(time.Second * 10):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():

		// will explain why done was closed
		err := ctx.Err()
		fmt.Println("error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// cancel before getting a responce
func reqHello() {
	time.Sleep(time.Second * 1)
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/hello", nil)
	if err != nil {
		fmt.Println(err)
	}

	ctxDL, cancel := context.WithCancel(req.Context())

	go func() {
		time.Sleep(time.Second * 5)
		cancel()
	}()
	req = req.WithContext(ctxDL)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))

}

// if we send request and cancel it before we get message, error will occur "error: context canceled"
func TestPkgHello() {
	http.HandleFunc("/hello", hello)
	fmt.Println("Server is live on 127.0.0.1:8080")
	go reqHello()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return
	}
}

// -- CONTEXT VALUES --

type nameKey string

const namekey nameKey = "name"

// we can send ctx with value
// Can be used in middlewares, regular request to provide additional info
func dataCtx(ctx context.Context) {
	fmt.Println("datactx: name value is", ctx.Value(namekey).(string))
}

func TestDataCtx() {
	ctx := context.Background()

	// ctx values are immutable
	ctx = context.WithValue(ctx, namekey, "Alex")
	ctx2 := context.WithValue(ctx, namekey, "Juli")

	dataCtx(ctx)
	dataCtx(ctx2)
	// wont be modified
	dataCtx(ctx)
}

// -- CONTEXT CANCELATION --

func cancelCtx(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	printchan := make(chan int)
	go cancelAnother(ctx, printchan)

	for i := 0; i < 3; i++ {
		printchan <- i
	}

	cancel()
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Cancel ctx finished")

}

func cancelAnother(ctx context.Context, printchan <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Println("cancelAnother err:", err)
			}
			fmt.Println("CancelAnother finished")
			return
		case num := <-printchan:
			fmt.Println("CancelAnother num = ", num)
		}
	}
}

func TestCancelCtx() {
	ctx := context.Background()
	cancelCtx(ctx)
}

// -- CONTEXT CANCELATION WITH DEADLINE --

// context with deadline cancels after deadline exceeded
func deadlineCtx(ctx context.Context) {
	// setting a deadline
	deadline := time.Now().Add(time.Millisecond * 1500)

	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()
	
	printCh := make(chan int)
	go cancelAnother(ctx, printCh)

	br := false
	for num := 1; num <= 3; num++ {
		if br { break }
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			br = true
		}
	}

	cancel()

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("DeadlineCtx: finished\n")
}

// Test context.WithDeadline func
func TestDeadlineCtx() {
	deadlineCtx(context.Background())
}
