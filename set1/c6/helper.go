package c6

import "fmt"

// HammingDistance reports he minimum number of substitutions required to
// change one string into the other. It is assumed that input1 and input2 are
// the same length.
func HammingDistance(input1, input2 string) int {
	var distance int
	for i := 0; i < len(input1); i++ {
		c1 := fmt.Sprintf("%08b", input1[i])
		c2 := fmt.Sprintf("%08b", input2[i])
		for j := 0; j < 8; j++ {
			if c1[j] != c2[j] {
				distance++
			}
		}
	}
	return distance
}
