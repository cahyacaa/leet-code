package problems

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
			wordMap[charTotal] = append(wordMap[charTotal], s)

		} else {
			wordMap[charTotal] = append(wordMap[charTotal], s)
		}
		charTotal = 0
	}

	for _, val := range wordMap {
		if len(val) > 1 {
			result = append(result, val...)
		}
	}
	return
}
