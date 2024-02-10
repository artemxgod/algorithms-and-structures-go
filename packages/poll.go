package packages

import (
	"fmt"
	"sync"
)

func Poll() {
	b := make([]byte, 0)

	b = append(b, 'a')
	fmt.Println(b)

	putBytes(b)
	fmt.Println(b)

	c := getBytes()
	fmt.Println(c)
}

func getBytes() (byteRes []byte) {
	poolBytes := bytesPoll.Get()
	if poolBytes != nil {
		byteRes = poolBytes.([]byte)
	}
	return
}

func putBytes(p_bytes []byte) {
	if cap(p_bytes) <= 1024 {
		p_bytes = p_bytes[:0]
		bytesPoll.Put(p_bytes)
	}
}

var bytesPoll = sync.Pool{
	New: func() any {
		return []byte{}
	},
}

// sync.Pool can be used to save some piece of memory.
// it helps to save memory because it takes some time before garbage collector clear unused memory
