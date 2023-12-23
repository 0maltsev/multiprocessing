package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ConcurrentPartition - ConcurrentQuicksort function for partitioning the array (randomized choice of a pivot)
// func ConcurrentPartition(A []int, p int, r int) int {
// 	index := rand.Intn(r-p) + p
// 	pivot := A[index]
// 	A[index] = A[r]
// 	A[r] = pivot
// 	x := A[r]
// 	j := p - 1
// 	i := p
// 	for i < r {
// 		if A[i] <= x {
// 			j++
// 			tmp := A[j]
// 			A[j] = A[i]
// 			A[i] = tmp
// 		}
// 		i++
// 	}
// 	temp := A[j+1]
// 	A[j+1] = A[r]
// 	A[r] = temp
// 	return j + 1
// }

func ConcurrentQuicksort(A []int, p int, r int, wg *sync.WaitGroup) {
	q, k := partition(A, p, r)
	if r -p > 1000 {
		wg.Add(2)
		go ConcurrentQuicksort(q, p, k-1, wg)
		go ConcurrentQuicksort(q, k+1, r, wg)
	} else {
		quickSort(q, p, k-1)
		quickSort(q, k+1, r)
	}
	time.Sleep(1 * time.Second)
	defer wg.Done()

}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func main() {

	sortSize := 100000000
	unsorted := make([]int, 0, sortSize)
	unsorted = rand.Perm(sortSize)

	start := time.Now()
	
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go ConcurrentQuicksort(unsorted, 0, len(unsorted)-1, &wg)
	// wg.Wait()
	quickSortStart(unsorted)
	
	elapsed := time.Since(start)

	fmt.Print(elapsed)

}

