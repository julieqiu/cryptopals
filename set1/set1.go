package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/julieqiu/derrors"
)

// HexToBytes converts a hex string to the equivalent in bytes (decimal).
func HexToBytes(input string) ([]byte, error) {
	src := []byte(input)
	dst := make([]byte, hex.DecodedLen(len(src)))
	if _, err := hex.Decode(dst, src); err != nil {
		return nil, err
	}
	return dst, nil
}

// BytesToBase64 encodes a byte string to base64.
func BytesToBase64(b []byte) []byte {
	eb := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(eb, b)
	return eb
}

// BytesToHex converts a byte slice to the equivalent in hex.
func BytesToHex(src []byte) string {
	return hex.EncodeToString(src)
}

var LetterToValue = map[string]int{
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

func DecryptHexMultiByteXOR(h, key string) []byte {
	j := 0
	out := make([]byte, len(h))
	for i := 0; i < len(h); i++ {
		out[i] = byte(key[j]) ^ byte(h[i])
		j += 1
		if j == 3 {
			j = 0
		}
	}
	return out
}

func DecryptHexSingleByteXOR(h string) (output string, highestScore int, err error) {
	defer derrors.Wrap(&err, "DecryptHexSingleByteXOR(%q)", h)
	b, err := HexToBytes(h)
	if err != nil {
		return "", 0, err
	}

	for k := 0; k < 256; k++ {
		out, s := decryptWithKey(b, k)
		if s > highestScore {
			highestScore = s
			output = string(out)
		}
	}
	return output, highestScore, nil
}

func decryptWithKey(input []byte, key int) (out []byte, score int) {
	for i := 0; i < len(input); i++ {
		c := input[i] ^ byte(key)
		out = append(out, c)
		l := strings.ToUpper(string(c))
		v, ok := LetterToValue[l]
		if ok {
			score += v
		}
	}
	return out, score
}

// HammingDistance reports he minimum number of substitutions required to
// change one string into the other. It is assumed that input1 and input2 are
// the same length.
func HammingDistance(input1, input2 string) int {
	var distance int
	for i := 0; i < len(input1); i++ {
		c1 := fmt.Sprintf("%08b", input1[i])
		c2 := fmt.Sprintf("%08b", input2[i])
		for j := 0; j < 8; j++ {
			if c1[j] != c2[j] {
				distance++
			}
		}
	}
	return distance
}

func Base64ToHex(src string) (_ []byte, err error) {
	defer derrors.Wrap(&err, "base64ToHex(%q)", src)
	return base64.StdEncoding.DecodeString(src)
}
