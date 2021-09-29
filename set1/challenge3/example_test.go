package challenge3_test

import (
	"fmt"
	"log"

	"github.com/julieqiu/cryptopals/set1"
)

const input = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

func Example() {
	out, _, err := set1.DecryptSingleByteXOR(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
	// output: Cooking MC's like a pound of bacon
}
