package c6_test

import (
	"fmt"

	"github.com/julieqiu/cryptopals/set1/c6"
)

func ExampleHammingDistance() {
	input1 := "this is a test"
	input2 := "wokka wokka!!!"
	ans := c6.HammingDistance(input1, input2)
	fmt.Println(ans)
	// output: 37
}
