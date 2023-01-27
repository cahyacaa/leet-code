package main

import (
	"fmt"
	"leet-code/problems"
)

func main() {
	//	twoSumProblem
	tsProblem := problems.TwoSum{}
	tsProblem.Input = []int{3, 2, 3}
	tsProblem.TargetSum = 6

	resultTWProblem, err := tsProblem.FindIndexForTwoSum()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Index For Targeted Sum Is : %v \n", resultTWProblem)

	/*
		Anagram Detection problem
	*/
	adProblem := problems.AnagramDetectionProblem{}

	adProblem.Input = []string{"makan", "nakam", "kau", "aku", "kasur", "kusut", "tusuk"}
	adProblem.DetectAnagram()
}
