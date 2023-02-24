package tasks

import (
	"fmt"
	"log"
)

// У Андрея есть n резервуаров, расположенных в один ряд. Изначально в каждом резервуаре находится некоторое количество кислоты.
// Начальство Андрея требует, чтобы во всех резервуарах содержался одинаковый объем кислоты.
// За одну операцию он способен разлить по одному литру кислоты в каждый из первых
// k( 1 ≤ k ≤ n ) резервуаров. Обратите внимание, что для разных операций k
//  могут быть разными. Поскольку кислота очень дорогая, Андрею не разрешается выливать кислоту из резервуаров.
// Андрей просит вас узнать, можно ли уравнять объемы кислоты в резервуарах, и, если это возможно,
// то посчитать минимальное количество операций.

// Формат ввода
// Первая строка содержит число  n ( 1 ≤ n ≤ 100000) — количество резервуаров.
// Во второй строке содержатся  n целых чисел a[i] (1 ≤ a[i] ≤ 10^9) где a[i] означает исходный объём кислоты
// в i-м резервуаре в литрах.

// Формат вывода
// Если объемы кислоты в резервуарах можно уравнять, выведите минимальное количество операций, необходимых для этого.
// Если это невозможно, выведите «-1».

func Acid() {
	var n, a int
	fmt.Scan(&n)

	storages := make([]int, n)

	for idx := 0; idx < n; idx++ {
		fmt.Scan(&a)
		storages[idx] = a
	}

	if !checkRow(storages) {
		fmt.Print(-1)
		return
	}

	fmt.Print(storages[n-1] - storages[0])
}

func checkRow(arr []int) bool {
	min := arr[0]

	for idx := 1; idx < len(arr); idx++ {
		if arr[idx] > min {
			min = arr[idx]
		} else if arr[idx] < min {
			return false
		}
	} 
	return true
}

func FatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// My comments: At first i tried to check all array for element equality and that was too slow.
// Then i realized that i do not need to check array or even change it because for one time element can be
// incremented by 1, so the min amount of action needed is 'maxval' - 'minval'. In this task max val will always be the last one
// and min val will always be the first otherwise we cannot make elements equal (because we add acid starting from the first storage)
