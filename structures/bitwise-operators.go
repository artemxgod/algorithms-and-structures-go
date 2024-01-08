package structures

import "fmt"

func BitOperators() {
	Shift()
	Logic()
}

// shift сдвигает биты влево(увеличение) и вправо(уменьшение). Shift может заменить умнажение числа на 2 в n-ой степени
func Shift() {
	a, b := 2, 16
	fmt.Println(a >> 1) // 1  
	fmt.Println(b << 3) // 128
}

func Logic() {
	//  AND
	fmt.Println("2 AND 1:", 2 & 1) // can check odd/even number, 0 for even and 1 for odd (10 & 01 = 00)
	//  OR
	fmt.Println("2 OR 1", 2 | 1) // + (10 | 01 = 11) 
	// XOR 
	fmt.Println("2 XOR 1", 2 ^ 1) // see xor usage func (10 ^ 01) == 11

}


// XOR USAGE
func FindUnique() {
	var result int
	arr := []int{1, 2, 3, 2, 1}
	for _, elem := range arr {
		result = result ^ elem
	}

	fmt.Println(result)
}

