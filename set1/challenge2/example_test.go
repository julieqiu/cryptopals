package challenge2_test

import (
	"encoding/hex"
	"fmt"
	"log"
)

const (
	input1 = "1c0111001f010100061a024b53535009181c"
	input2 = "686974207468652062756c6c277320657965"
)

func Example() {
	b1, err := hexToBytes(input1)
	if err != nil {
		log.Fatal(err)
	}
	b2, err := hexToBytes(input2)
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
	fmt.Println(bytesToHex(out))
	// output: 746865206b696420646f6e277420706c6179
}

// hexToBytes converts a hex string to the equivalent in bytes (decimal).
func hexToBytes(input string) ([]byte, error) {
	src := []byte(input)
	dst := make([]byte, hex.DecodedLen(len(src)))
	if _, err := hex.Decode(dst, src); err != nil {
		return nil, err
	}
	return dst, nil
}

// bytesToHex converts a byte slice to the equivalent in hex.
func bytesToHex(src []byte) string {
	return hex.EncodeToString(src)
}
