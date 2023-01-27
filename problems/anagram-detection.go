package problems

import "fmt"

type AnagramDetection interface {
	DetectAnagram() []string
}

type AnagramDetectionProblem struct {
	Input []string
}

func (ad *AnagramDetectionProblem) DetectAnagram() (result []string) {
	var charTotal int32
	wordMap := make(map[int32][]string)
	for _, s := range ad.Input {
		for _, charUnicode := range s {
			charTotal += charUnicode
		}
		if val, _ := wordMap[charTotal]; len(val) >= 1 {
			result = append(result, val...)
			result = append(result, s)

		}
		wordMap[charTotal] = append(wordMap[charTotal], s)
		charTotal = 0
	}
	fmt.Println(result)
	return
}
