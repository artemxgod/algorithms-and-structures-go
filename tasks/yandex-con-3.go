package tasks

import (
	"fmt"
	"math"
	"sort"
)

// Рассмотрим целочисленный массив a длины n
// . Назовём расстоянием от индекса i до множества индексов S величину dist(i,S) = ∑ j ∈ S ∣∣ a[i] − a[j] ∣∣.
//Зафиксируем целое число k. Рассмотрим функцию
// f(i) = min dist (i,S), где минимум берётся по множествам S размера k, не содержащим индекс i.
// Определите значение f(i) для всех i от 1 до n.

// Формат ввода
// В первой строке заданы два целых числа n и k (2≤n≤300000, 1≤k<n), описанные в условии.
// Во второй строке содержится n целых чисел a[i] (1≤a[i]≤109) — элементы массива a.

// Формат вывода
// Выведите n целых чисел: значения f(i) для i=1,i=2,…,i=n.

func MinDistance() {
	var n, k int 
	fmt.Scanf("%d %d\n", &n, &k)
	arr := make([]int, n)

	for idx := range arr {
		fmt.Scan(&arr[idx])
	}
	
	for idx := range arr {
		fmt.Print(sum(arr, idx, k), " ")
	}
}

func sum(arr []int, i, k int) float64 {
	var res float64
	sub := make([]int, 0)
	sub = append(sub, arr[:i]...)
	sub = append(sub, arr[i+1:]...)
	sort.Ints(sub)
	
	for idx := 0; idx < k; idx++ {
		num, rmIdx := findClosest(sub, arr[i])
		res += float64(num)
		sub = append(sub[:rmIdx], sub[rmIdx+1:]...)
	}
	return res
}

func findClosest(sub []int, num int) (int, int) {
	var idxRes int
	mindist := int(math.Abs(float64(num - sub[0])))
	for  i := 1; i < len(sub); i++ {
		if int(math.Abs(float64(num - sub[i]))) < mindist {
			mindist = int(math.Abs(float64(num - sub[i])))
			idxRes = i
		} else {
			return mindist, idxRes
		}
	}

	return mindist, idxRes
}
