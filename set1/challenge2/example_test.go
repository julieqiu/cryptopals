package challenge2_test

import (
	"fmt"
	"log"

	"github.com/julieqiu/cryptopals/set1"
)

const (
	input1 = "1c0111001f010100061a024b53535009181c"
	input2 = "686974207468652062756c6c277320657965"
)

func Example() {
	b1, err := set1.HexToBytes(input1)
	if err != nil {
		log.Fatal(err)
	}
	b2, err := set1.HexToBytes(input2)
	if err != nil {
		log.Fatal(err)
	}
	if len(b1) != len(b2) {
		log.Fatalf("len(b1) = %d and len(b2) = %d", len(b1), len(b2))
	}

	var out []byte
	for i := 0; i < len(b1); i++ {
		out = append(out, b1[i]^b2[i])
	}
	// fmt.Printf("%s", out) = the kid don't play
	fmt.Println(set1.BytesToHex(out))
	// output: 746865206b696420646f6e277420706c6179
}
