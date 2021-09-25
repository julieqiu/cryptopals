package c3_test

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

const input = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

func Example() {
	inputBytes, err := hexToBytes(input)
	if err != nil {
		log.Fatal(err)
	}

	var (
		out     string
		highest int
	)
	for k := 0; k < 256; k++ {
		var (
			b []byte
			s int
		)
		for i := 0; i < len(inputBytes); i++ {
			c := inputBytes[i] ^ byte(k)
			b = append(b, c)
			l := strings.ToUpper(string(c))
			v, ok := letterToValue[l]
			if ok {
				s += v
			}
		}

		if s > highest {
			out = string(b)
			highest = s
		}
	}
	fmt.Println(out)
	// output: Cooking MC's like a pound of bacon
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

var letterToValue = map[string]int{
	"E": 13,
	"T": 12,
	"A": 11,
	"O": 10,
	"I": 9,
	"N": 8,
	" ": 7,
	"S": 6,
	"H": 5,
	"R": 4,
	"D": 3,
	"L": 2,
	"U": 1,
}
