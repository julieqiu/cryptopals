package c6

func HammingDistance(input1, input2 string) int {
	var d int
	for i := 0; i < len(input1); i++ {
		for j := 0; j < 8; j++ {
			if input1[i]&(1<<uint(j)) != input2[i]&(1<<uint(j)) {
				d++
			}
		}
	}
	return d
}
