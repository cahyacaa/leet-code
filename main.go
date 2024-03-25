package main

import (
	"fmt"
	data_structure "leet-code/data-structure"
	"leet-code/problems"
)

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	for pos, _ := range s {
		lenSubstr := getStrLong(pos, s)
		if lenSubstr >= maxLen {
			maxLen = lenSubstr
		}
	}

	return maxLen
}

func getStrLong(idx int, s string) int {
	strMap := make(map[byte]bool)

	res := 0
	for i := idx; i < len(s); i++ {
		if _, ok := strMap[s[i]]; ok {
			break
		}
		strMap[s[i]] = true
		res++
	}
	return res
}
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
	adProblem.Input = []string{"makan", "nakam", "uka", "kau", "aku", "kasur", "kusut", "tusuk", "raba"}
	fmt.Printf("Anagrams words : %v \n", adProblem.DetectAnagram())

	/*
		Longest Palindromic Substring
	*/

	lpProblem := problems.LongestPalindromeProblem{}
	lpProblem.Input = "aba"
	fmt.Println(lpProblem.FindLongestPalindrome())

	fmt.Println(lengthOfLongestSubstring("abababab"))

	/*
		Binary Tree Data Structure
	*/

	binaryTree := data_structure.NewBinaryTrees(1)
	for i := 2; i <= 10; i++ {
		binaryTree.Add(i)
	}

	binaryTree.Print()
}
