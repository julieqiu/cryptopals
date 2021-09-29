package challenge5_test

import (
	"fmt"

	"github.com/julieqiu/cryptopals/set1"
)

const (
	key = "ICE"

	input = `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
)

func Example() {
	b := set1.EncryptRepeatingXOR(input, key)
	out := set1.BytesToHex(b)
	fmt.Println(out)
	// output:
	// 0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f

}
