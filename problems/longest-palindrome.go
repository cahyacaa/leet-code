package problems

import "fmt"

type LongestPalindrome interface {
	FindLongestPalindrome() int
}

type LongestPalindromeProblem struct {
	Input string
}

func (lp *LongestPalindromeProblem) isPalindrome(word string) bool {
	reverseString := ""
	for i := len(word) - 1; i >= 0; i-- {
		reverseString += fmt.Sprintf("%c", word[i])
	}
	return reverseString == word
}

func (lp *LongestPalindromeProblem) FindLongestPalindrome() (result int) {
	maksimumExpand := 0
	for i, _ := range lp.Input {
		if i == 0 {
			result = 1
		} else {
			maksimumExpand = i - (i - 1)
			for l := 0; l < maksimumExpand; l++ {

			}
		}
	}
	return
}
