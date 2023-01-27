package problems

import (
	"errors"
)

type Interface interface {
	FindIndexForTwoSum() ([]int, error)
}

type TwoSum struct {
	Input     []int
	TargetSum int
}

func (tw *TwoSum) FindIndexForTwoSum() ([]int, error) {
	for i, _ := range tw.Input {
		currentNumber := tw.Input[i]
		if i == 0 && currentNumber == tw.TargetSum {
			return []int{i}, nil
		}
		if i > 0 && tw.Input[i]+tw.Input[i-1] == tw.TargetSum {
			return []int{i - 1, i}, nil
		}
	}

	if len(tw.Input) == 0 {
		return []int{}, errors.New("input value cannot be empty")
	}
	return []int{}, nil
}
