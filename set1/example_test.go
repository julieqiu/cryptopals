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
	s := "this is a test"
	key := "hi"
	b := set1.EncryptRepeatingXOR(s, key)
	keysizes := set1.GuessRepeatingXORKeySize(b)
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
