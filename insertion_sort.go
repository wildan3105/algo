// insertion_sort.go, an efficient algorithm for sorting a small number of elements

package main

import "fmt"

var numbers = []int{5,3,4,2,6,7}

func main() {
	fmt.Printf("sequence: %d\n", insertion(numbers, false))
}

func insertion(arr []int, smallerToBig bool) []int {

	for j := 0; j < len(arr); j++ {
		key := arr[j]
		i := j-1

		if(smallerToBig == true) {
			for(i >=0 && arr[i] > key){
				arr[i+1] = arr[i]
				i--
			}
			arr[i+1] = key
		} else {
			for(i >=0 && arr[i] < key){
				arr[i+1] = arr[i]
				i--
			}
			arr[i+1] = key
		}
	}
	return arr
}