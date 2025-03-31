package day4

func Calculate(condition string, numbers ...int) int {
	res := 0
	if condition == "multiply" {
		res = 1
		for _, num := range numbers {
			res *= num
		}

	} else if condition == "add" {
		for _, num := range numbers {
			res += num
		}
	} else {
		for _, num := range numbers {
			res -= num
		}
	}
	return res
}
