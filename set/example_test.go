package s1c1_test

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

const input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

func Example() {
	b, err := hexToBytes()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", bytesToBase64(b))
	// output: SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
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

// bytesToBase64 encodes a byte string to base64.
func bytesToBase64(b []byte) []byte {
	eb := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(eb, b)
	return eb
}
