package problems

import (
	"errors"
)

type TwoSumInterface interface {
	FindIndexForTwoSum() ([]int, error)
}

type TwoSum struct {
	Input     []int
	TargetSum int
}

func (tw *TwoSum) FindIndexForTwoSum() ([]int, error) {
	numberMap := make(map[int]int, 0)
	indexOfNumberSum := make([]int, 0, 2)
	for i, number := range tw.Input {
		_, exists := numberMap[tw.TargetSum-number]
		if exists {
			if len(indexOfNumberSum) == 2 {
				break
			}
			indexOfNumberSum = append(indexOfNumberSum, numberMap[tw.TargetSum-number])
			indexOfNumberSum = append(indexOfNumberSum, i)
		}
		numberMap[number] = i
	}

	if len(tw.Input) == 0 {
		return []int{}, errors.New("input value cannot be empty")
	}
	return indexOfNumberSum, nil
}
