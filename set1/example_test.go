package set1_test

import (
	"fmt"

	"github.com/julieqiu/cryptopals/set1"
)

const (
	input1 = "this is a test"
	input2 = "wokka wokka!!!"
)

func ExampleHammingDistance() {
	ans := set1.HammingDistance(input1, input2)
	fmt.Println(ans)
	// output: 37
}

func ExampleGuessRepeatingXORKeySize() {
	var h []byte
	keysizes := set1.GuessRepeatingXORKeySize(h)
	fmt.Println(keysizes)
	// output: keysizes
}

func ExampleDecryptRepeatingXOR() {
	var (
		h  []byte
		ks int
	)
	keysizes, score := set1.DecryptRepeatingXOR(h, ks)
	fmt.Println(keysizes, score)
	// output: TODO
}
