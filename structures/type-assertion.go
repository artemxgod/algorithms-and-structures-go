package structures

import "fmt"

// A type assertion provides access to an interface value's underlying concrete value.
func Assertion() {
	var i interface{} = "hello"
	var num interface{} = 3.141

	// will print
	if s, ok := i.(string); ok {
		fmt.Println(s)
	}

	// won't print
	if n, ok := i.(int); ok {
		fmt.Println(n)
	}

	// won't print
	if n, ok := num.(int); ok {
		fmt.Println(n)
	}
	
	// will print
	if n, ok := num.(float64); ok {
		fmt.Println(n)
	}

	// if we try to print value with wrong type and forget to check it => panic will occur
}