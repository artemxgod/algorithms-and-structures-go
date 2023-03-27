package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// Go allows us to run code concurrently using goroutines.
// However, when concurrent processes access the same piece of data, it can lead to race conditions.

func isEven(n int) bool {
	return n & 1 == 0 
}

func MutexLock() {
	n := 0
	var mut sync.Mutex

	go func() {
		// locks the data used in this goroutine
		mut.Lock()
		defer mut.Unlock()
		resEven := isEven(n)
		time.Sleep(time.Millisecond * 5)

		if resEven {
			fmt.Println(n, "is even number")
			return
		}
		fmt.Println(n, "is odd number")
	}()

	go func() {
		mut.Lock()
		n++
		mut.Unlock()
	}()
	
	time.Sleep(time.Second)
}

// Testing mutex lock, should not let number increment before defining even/odd is the number
func TestMutexLock() {
	MutexLock()
}

// In this case, we can use a sync.RWMutex type, which has different locks for reading and writing to data:
func RWMutexLock() {
	n := 0
	var mut sync.RWMutex

	// Since we are only reading data here, we can call the `RLock` 
	// method, which obtains a read-only lock
	go func() {
		mut.RLock()
		defer mut.RUnlock()

		resEven := isEven(n)
		time.Sleep(time.Millisecond * 5)

		if resEven {
			fmt.Println(n, "is even number")
			return
		}
		fmt.Println(n, "is odd number")
	}()

	// here we wont wait for the first goroutine to end because n is locked for writing, so n is readonly
	go func() {
		mut.RLock()
		defer mut.RUnlock()
		nIsPositive := n > 0
		time.Sleep(5 * time.Millisecond)
		if nIsPositive {
			fmt.Println(n, " is positive")
			return
		}
		fmt.Println(n, "is not positive")
	}()

	// Since we are writing into data here, we use the
	// `Lock` method, like before
	go func() {
		mut.Lock()
		n++
		mut.Unlock()
	}()

	time.Sleep(time.Second)
}

// test RWmutex lock
func TestRWMutexLock() {
	RWMutexLock()
}

