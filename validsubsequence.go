// main is an algoexpert
package algoexpert

// IsValidSubsequence using range
func IsValidSubsequence(array []int, sequence []int) bool {
	current := 0
	for _, value := range array {
		if sequence[current] == value {
			current++
		}
		if current == len(sequence) {
			return true
		}
	}

	return false
}

// IsValidSubsequenceTwo without range
func IsValidSubsequenceTwo(array []int, sequence []int) bool {
	i, j := 0, 0
	for i < len(array) && j < len(sequence) {
		if array[i] == sequence[j] {
			j++
		}
		i++
	}
	return j == len(sequence)
}
