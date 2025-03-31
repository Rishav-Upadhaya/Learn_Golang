package day4

import "fmt"

func CheckArray() {
	arr := [5]int{0, 2, 10, 19}
	fmt.Println(arr, "length=", len(arr))
	fmt.Println(arr[2])
	arrCopy := arr
	arrCopy[2] = 100

	fmt.Println("Original = ", arr)
	fmt.Println("New = ", arrCopy)

}
