// This file is generated using cmd/makedoc. DO NOT EDIT.
// To update, edit the doc.txt file in this directory.
// Then run
//     go run ./cmd/makedoc 1 2
//

// This package provides a solution to https://cryptopals.com/sets/1/challenges/2.
// 
// Fixed XOR Decryption in Go
// 
// The second challenge says:
// 
// Write a function that takes two equal-length buffers and produces their XOR combination.
// If your function works properly, then when you feed it the string:
// 
//   1c0111001f010100061a024b53535009181c
// 
// ... after hex decoding, and when XOR'd against:
// 
//   686974207468652062756c6c277320657965
// 
// ... should produce:
// 
//   746865206b696420646f6e277420706c6179
// 
// 
// The Basics
// 
// This challenge requires you to understand XOR, which is an operator used to
// compare two bits (i.e. 0 and 1).
// 
// It returns 1 if the two bits are different, and 0 otherwise. Here’s all the
// possible of XOR-ing two bits:
// 
//   | bit1 | 0  | 0  | 1  | 1 |
//   | ---  | -- | -- | -- | --|
//   | bit2 | 0  | 1  | 0  | 1 |
//   | ---  | -- | -- | -- | --|
//   | XOR  | 0  | 1  | 1  | 0 |
// 
// In order to XOR two numbers, we need to:
// 
// 1. Check that they are the same length
// 2. Represent each value in binary
// 3. Line them up, and XOR each digit
// 
// 
// Solution
// 
// We have these two hex values:
// 
//   1c0111001f010100061a024b53535009181c
//   686974207468652062756c6c277320657965
// 
// We can use our hexToBytes function from challenge 1, and convert the hex to
// bytes/decimal.
// 
// This will give us:
// 
//   // 1c0111001f010100061a024b53535009181c
//   [ 28   1  17  0  31   1   1  0  6  26   2  75 83  83 80   9  24  28]
// 
//   // 686974207468652062756c6c277320657965
//   [104 105 116 32 116 104 101 32 98 117 108 108 39 115 32 101 121 101]
// 
// Then, we convert each number to binary, and XOR the bits.
// 
// For example, the first ones are 28 and 104, which map to:
// 
// 
//   | 28  | 0 | 0 | 0 | 1 | 1 | 1 | 0 | 0 |
//   | --- | - | - | - | - | - | - | - | - |
//   | 104 | 0 | 1 | 1 | 0 | 1 | 0 | 0 | 0 |
//   | --- | - | - | - | - | - | - | - | - |
//   | XOR | 0 | 1 | 1 | 1 | 0 | 1 | 0 | 0 |
// 
// This gives us 116 for the first byte.
// 
// Lastly, we want to hex encode the output.
//
//
package challenge2