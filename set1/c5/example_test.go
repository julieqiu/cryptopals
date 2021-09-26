package c5_test

import (
	"encoding/hex"
	"fmt"
)

const (
	key = "ICE"

	input = `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal
Encrypt it, under the key "ICE", using repeating-key XOR.`
)

func Example() {
	out := decrypt()
	fmt.Println(bytesToHex(out))
	// output:
	// 0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272
	// a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f

}

func decrypt() []byte {
	out := make([]byte, len(input))
	for _, k := range []byte(key) {
		for i := 0; i < len(input); i++ {
			c := input[i]
			out[i] = k ^ c
		}
	}
	return out
}

// bytesToHex converts a byte slice to the equivalent in hex.
func bytesToHex(src []byte) string {
	return hex.EncodeToString(src)
}
