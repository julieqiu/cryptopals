package challenge6_test

import (
	"fmt"

	"github.com/julieqiu/cryptopals/set1/challenge6"
)

const (
	input1 = "this is a test"
	input2 = "wokka wokka!!!"
)

func ExampleHammingDistance() {
	ans := challenge6.HammingDistance(input1, input2)
	fmt.Println(ans)
	// output: 37
}
