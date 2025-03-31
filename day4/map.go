package day4

import (
	"fmt"
	"maps"
)

func CheckMap() {
	person := map[string]interface{}{
		"name":    "ram",
		"age":     45,
		"address": "kathmandu",
	}
	fmt.Println("ram: ", person, "Length: ", len(person))
	for key, value := range person {
		fmt.Println("key: ", key, "value: ", value)
	}
	student := make(map[string]any)
	student["name"] = "Ramesh"
	student["age"] = 56
	student["class"] = "10th"
	fmt.Println("Student: ", student)

	book := map[string]any{
		"book1": "book1",
		"book2": "book2",
		"book3": "book3",
	}
	mergedMap := make(map[string]any)
	maps.Copy(mergedMap, book)
	maps.Copy(mergedMap, student)
	fmt.Println("merged map:", mergedMap)
}
