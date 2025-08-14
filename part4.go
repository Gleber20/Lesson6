package main

import "fmt"

func DoubleSlice(s []int) []int {
	double := make([]int, len(s))
	for i, v := range s {
		double[i] = v * 2
	}
	return double
}

func main() {
	nums := [5]int{10, 20, 30, 40, 50}
	fmt.Println(nums)
	slice := nums[0:3]
	fmt.Println(slice)
	DoubleSl := DoubleSlice(slice)
	fmt.Println(DoubleSl)
	merge := append(slice, DoubleSl...)
	fmt.Println(merge)
}
