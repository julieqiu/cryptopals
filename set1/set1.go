package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
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

func Base64ToBytes(src string) (_ []byte, err error) {
	defer derrors.Wrap(&err, "base64ToHex(%q)", src)
	return base64.StdEncoding.DecodeString(src)
}

func EtaoinShrdluScore(input []byte) int {
	var score int
	for _, c := range input {
		l := strings.ToUpper(string(c))
		v, ok := letterToValue[l]
		if ok {
			score += v
		}
	}
	return score
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

func DecryptMultiByteXOR(h, key string) []byte {
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

func DecryptSingleByteXOR(h string) (output string, highestScore int, err error) {
	defer derrors.Wrap(&err, "DecryptHexSingleByteXOR(%q)", h)
	b, err := HexToBytes(h)
	if err != nil {
		return "", 0, err
	}
	out, score := decryptBytesSingleByteXOR(b)
	return string(out), score, nil
}

func decryptBytesSingleByteXOR(b []byte) (output []byte, highestScore int) {
	for k := 0; k < 256; k++ {
		out, s := decryptWithKey(b, k)
		if s > highestScore {
			highestScore = s
			output = out
		}
	}
	return output, highestScore
}

func decryptWithKey(input []byte, key int) (out []byte, score int) {
	for i := 0; i < len(input); i++ {
		c := input[i] ^ byte(key)
		out = append(out, c)
	}
	score = EtaoinShrdluScore(out)
	return out, score
}

// HammingDistance reports the minimum number of substitutions required to
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

// hammingDistance reports the minimum number of substitutions required to
// change one string into the other, normalized by input length.
// It is assumed that input1 and input2 are the same length.
func hammingDistanceNormalized(input1, input2 []byte) int {
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
	return distance * 100 / len(input1)
}

func GuessRepeatingXORKeySize(h []byte) []int {
	keysizeToScore := make(map[int]int, 3)
	maxHighest := math.MaxInt
	for size := keysizeMin; size <= keysizeMax; size++ {
		// Break up f into 2 keysize blocks.
		block1 := h[0:size] // fix panic here
		block2 := h[size : size*2]
		// Get hamming distance between these blocks, normalized.
		val := hammingDistanceNormalized(block1, block2)
		if val < maxHighest {
			maxHighest = val
			keysizeToScore[size] = val
			if len(keysizeToScore) > maxGuesses {
				for ks, score := range keysizeToScore {
					if score > maxHighest {
						delete(keysizeToScore, ks)
					}
				}
			}
		}
	}
	var results []int
	for ks := range keysizeToScore {
		results = append(results, ks)
	}
	return results
}

const (
	keysizeMin = 2
	keysizeMax = 40
	maxGuesses = 3
)

func DecryptRepeatingXOR(b []byte, ks int) ([]byte, int) {
	chunks := groupChunks(b, ks)

	var outputChunks [][]byte
	for _, c := range chunks {
		// 3. For each chunck, run single-byte XOR decrpytion.
		out, _ := decryptBytesSingleByteXOR(c)
		outputChunks = append(outputChunks, out)
	}

	// 4. Rerrange the text back.
	out := rearrangeChunks(outputChunks)
	score := EtaoinShrdluScore(out)
	return out, score
}

func groupChunks(b []byte, ks int) [][]byte {
	chunks := make([][]byte, ks)
	for i, c := range b {
		j := i % ks
		chunks[j] = append(chunks[j], c)
	}
	return chunks
}

func rearrangeChunks(chunks [][]byte) []byte {
	var r []byte

	longestChunkLength := len(chunks[0])
	// Look through each i-th position of the chunk (look at the first
	// character of each, then the second, then the third...)
	//
	// For each chunk, if i > length of the chunk, we've run out of letters in
	// that chunk, so continue to the next.
	// Otherwise, append to the result.
	for i := 0; i < longestChunkLength; i++ {
		for _, chunk := range chunks {
			if i >= len(chunk) {
				continue
			}
			r = append(r, chunk[i])
		}
	}
	return r
}
