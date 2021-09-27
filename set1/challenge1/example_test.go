package challenge1_test

import (
	"fmt"
	"log"

	"github.com/julieqiu/cryptopals/set1"
)

const input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

func Example() {
	b, err := set1.HexToBytes(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", set1.BytesToBase64(b))
	// output: SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
}
