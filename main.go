package main

import (
	"fmt"
	"leet-code/problems"
)

func main() {
	//	twoSumProblem
	tsProblem := problems.TwoSum{}
	tsProblem.Input = []int{1, 2, 3, 3, 4, 4}
	tsProblem.TargetSum = 6

	resultTWProblem, err := tsProblem.FindIndexForTwoSum()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Index For Targeted Sum Is : %v", resultTWProblem)
}
