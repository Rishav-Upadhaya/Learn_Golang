package day5

import "fmt"

// defining struct
type College struct {
	Name             string
	Address          string
	Students         []string
	TotalNoOfStudent int
	totalNoOfStaff   int
}

func DisplayCollegeInfo(c College) {
	fmt.Println("Name:", c.Name)
	fmt.Println("Address:", c.Address)
	fmt.Println("Students:", c.Students)
	fmt.Println("TotalNoOfStudents:", c.TotalNoOfStudent)
	fmt.Println("TotalNoOfStudents:", c.totalNoOfStaff)
}

func (c College) EmbeddedDisplayCollegeInfo() {
	fmt.Println("E Name: ", c.Name)
	fmt.Println("E Name: ", c.Address)
	fmt.Println("E Name: ", c.Students)
	fmt.Println("E Name: ", c.TotalNoOfStudent)
}
