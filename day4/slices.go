package day4

import "fmt"

func CheckSlice() {
	slc := []int{1, 2, 3, 4, 6}
	fmt.Println("lenght: ", len(slc))
	slcCopy := slc
	slcCopy[0] = 99
	fmt.Println("Org = ", slc)
	fmt.Println("New = ", slcCopy)

	slc1 := make([]int, len(slc), 8)

	copy(slc1, slc)

	slc1[2] = 200

	fmt.Println("org: ", slc)
	fmt.Println("org: ", slc1)
}
