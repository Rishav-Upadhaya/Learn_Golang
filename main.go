package main

import (
	"errors"
	"fmt"
	"golangproject/day4"
	"golangproject/day5"
	"log"
	"math"
	"strings"
)

func add(numbers ...int) (int, int, string) {
	fmt.Printf("array: %d", numbers)
	sum := 0
	for _, num := range numbers {
		sum = sum + num
	}
	return sum, len(numbers), string("successful addition")
}

func CalculateSquareRoot(x float64) (*float64, error) {
	if x < 0 {
		return nil, errors.New("negative value appear")
	} else {
		sqrt := math.Sqrt(x)
		return &sqrt, nil
	}
}

func main() {
	// res,length,mess:=add(10,5,23,23,231223,213)
	// fmt.Printf("\n length = %d \n result = %d \n message :%s",length,res,mess)
	//anonymous function
	func() {
		fmt.Printf(("this is anonymous \n"))
	}()
	display := func(name string) {
		fmt.Printf("name = %s", name)
	}
	display("harry")
	next := sequencegenrator()
	a := next() //closure function
	fmt.Printf("%d", a)
	b := next()
	fmt.Printf("%d", b)
	for i := 0; i < 10; i++ {
		res := next()
		fmt.Printf("%d\n", res)

	}
	abc := day4.Calculate("add", 23, 123, 213, 2)
	fmt.Printf("addition: %d\n", abc)
	day4.CheckArray()
	day4.CheckSlice()
	day4.CheckMap()

	samriddhi := day5.College{
		Name:             "Samriddhi",
		Address:          "Kathmandu",
		Students:         []string{"ram", "shyam", "sita"},
		TotalNoOfStudent: 3,
		// totalNoOfStaff: 5,
	}

	day5.DisplayCollegeInfo(samriddhi)
	samriddhi.EmbeddedDisplayCollegeInfo()

	circle := day5.Circle{
		Radius: 30,
	}

	rect := day5.Rectangle{
		Length:  500,
		Breadth: 34,
	}

	fmt.Println("Circle Info....")
	day5.DisplayShapeinfo(circle)
	fmt.Println("Rectangle Info....")
	day5.DisplayShapeinfo(rect)

	resSqrt, err := CalculateSquareRoot(36)
	defer func() {
		if r := recover(); r != nil {
			log.Println("recover from panic: ", r)
		}
	}()
	if err != nil {
		//error appear
		log.Println(err.Error())
	} else {
		//result case
		log.Println("result = ", *resSqrt)
	}

	capilatize, err := DisplayName("rishav")

	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}
	log.Println("Name:", capilatize)
	defer log.Println("deferred function called")
	log.Println("Normal execution")
	log.Println("Second Normal Execution")

}

func sequencegenrator() func() int {
	i := 0 //capture
	return func() int {
		i++ //retain
		return i
	}
}

func DisplayName(name string) (string, error) {
	if len(name) < 5 {
		println("Invalid Name")
		return "", errors.New("Length must be greater")
	} else {
		return strings.ToUpper(name), nil
	}
}
