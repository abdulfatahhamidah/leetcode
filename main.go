package main

import (
	"fmt"
	"github.com/abdulfatahhamidah/leetcode/twosum"

)

func main() {
	// fmt.Println("welcome to daily LeetCode")
	// fmt.Println("Open lib for more solution")

	list1 := []int{2, 7, 11, 15}
	list2 := []int{3, 2, 4}
	list3 := []int{3, 3}
	fmt.Println(twosum.TwoSumOptimze(list1, 9))
	fmt.Println(twosum.TwoSumOptimze(list2, 6))
	fmt.Println(twosum.TwoSumOptimze(list3, 6))
}
