package c5_test

import (
	"encoding/hex"
	"fmt"
)

const (
	key = "ICE"

	input = `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
)

func Example() {
	out := bytesToHex(decrypt(input))
	fmt.Println(out)
	// output:
	// 0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f

}

func decrypt(line string) []byte {
	j := 0
	out := make([]byte, len(line))
	for i := 0; i < len(line); i++ {
		out[i] = byte(key[j]) ^ byte(line[i])
		j += 1
		if j == 3 {
			j = 0
		}
	}
	return out
}

// bytesToHex converts a byte slice to the equivalent in hex.
func bytesToHex(src []byte) string {
	return hex.EncodeToString(src)
}
