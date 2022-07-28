package main

func subtractProductAndSum(n int) int {
	multiply := 1
	sum := 0
	for n > 0 {
		// fmt.Println(n%10)
		sum += n % 10
		multiply *= n % 10
		n = n / 10
	}
	return multiply - sum
}
