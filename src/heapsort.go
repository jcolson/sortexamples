package main

import "fmt"

func main() {
	arr := []int{121, 10, 130, 57, 36, 17}
	n := len(arr)

	heapSort(arr, n)

	fmt.Printf("Sorted array is\n%v\n", arr)
}

func heapify(arr []int, n int, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	// if left child is larger than root
	if l < n && arr[l] > arr[largest] {
		largest = l
	}

	// if right child is larger than largest so far
	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	// if largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]

		// recursively heapify the affected sub-tree
		heapify(arr, n, largest)
	}
}

func heapSort(arr []int, n int) {
	// build heap (rearrange array)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// one by one extract an element from heap
	for i := n - 1; i >= 0; i-- {
		// move current root to end
		arr[0], arr[i] = arr[i], arr[0]

		// call max heapify on the reduced heap
		heapify(arr, i, 0)
	}
}
