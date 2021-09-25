package c3_test

import (
	"encoding/hex"
    "fmt"
	"log"
)

const input = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

func Example() {
	b, err := hexToBytes(input)
	if err != nil {
		log.Fatal(err)
	}

    var out []byte
    for key := 0; key < 256; key ++ {
        for i := 0; i < len(b); i++ {
            out = append(out, b[i]^byte(key))
        }
        fmt.Println(out)
	}
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
func bytesToHex(src []byte) (string) {
	return hex.EncodeToString(src)
}
