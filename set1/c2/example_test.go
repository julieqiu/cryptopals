package s1c2_test

import (
	"encoding/hex"
	"log"
)

const (
	input1 = "1c0111001f010100061a024b53535009181c"
	input2 = "686974207468652062756c6c277320657965"
)

func Example() {
	b, err := hexToBytes()
	if err != nil {
		log.Fatal(err)
	}
	// output: 746865206b696420646f6e277420706c6179
}

// hexToBytes converts a hex string to the equivalent in bytes (decimal).
func hexToBytes() ([]byte, error) {
	src := []byte(input)
	dst := make([]byte, hex.DecodedLen(len(src)))
	if _, err := hex.Decode(dst, src); err != nil {
		return nil, err
	}
	return dst, nil
}
