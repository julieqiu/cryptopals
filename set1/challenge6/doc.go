// This file is generated using cmd/makedoc. DO NOT EDIT.
// To update, edit the doc.txt file in this directory.
// Then run
//     go run ./cmd/makedoc 1 6
//

// This package provides a solution to https://cryptopals.com/sets/1/challenges/6.
// 
// Break repeating key XOR
// 
//   It is officially on, now.
// 
//   This challenge isn't conceptually hard, but it involves actual error-prone
//   coding. The other challenges in this set are there to bring you up to speed.
//   This one is there to qualify you. If you can do this one, you're probably just
//   fine up to Set 6.
// 
// There's a file here.
// https://cryptopals.com/static/challenge-data/6.txt
// 
// It's been base64'd after being encrypted with repeating-key XOR.
// 
// Decrypt it.
// 
// Here's how:
// 
// 1. Let KEYSIZE be the guessed length of the key; try values from 2 to (say) 40.
// 
// 2. Write a function to compute the edit distance/Hamming distance between two
// strings. The Hamming distance is just the number of differing bits. The
// distance between:
// 
//   this is a test
// and
// 
//   wokka wokka!!!
// 
// is 37. Make sure your code agrees before you proceed.
// 
// 3. For each KEYSIZE, take the first KEYSIZE worth of bytes, and the second
// KEYSIZE worth of bytes, and find the edit distance between them. Normalize this
// result by dividing by KEYSIZE.
// 
// 4. The KEYSIZE with the smallest normalized edit distance is probably the key.
// You could proceed perhaps with the smallest 2-3 KEYSIZE values. Or take 4
// KEYSIZE blocks instead of 2 and average the distances.
// 
// 5. Now that you probably know the KEYSIZE: break the ciphertext into blocks of KEYSIZE length.
// 
// 6. Now transpose the blocks: make a block that is the first byte of every
// block, and a block that is the second byte of every block, and so on.
// 
// 7. Solve each block as if it was single-character XOR. You already have code to
// do this.
// 
// 8. For each block, the single-byte XOR key that produces the best looking
// histogram is the repeating-key XOR key byte for that block. Put them together
// and you have the key.
// 
// This code is going to turn out to be surprisingly useful later on. Breaking
// repeating-key XOR ("Vigenere") statistically is obviously an academic exercise,
// a "Crypto 101" thing. But more people "know how" to break it than can actually
// break it, and a similar technique breaks something much more important.
// 
//   No, that's not a mistake.
// 
//   We get more tech support questions for this challenge than any of the other
//   ones. We promise, there aren't any blatant errors in this text. In particular:
//   the "wokka wokka!!!" edit distance really is 37.
// 
// 
// The Basics
// 
// This challenge requires you to understand two new operators to compare two bits
// (i.e. 0 and 1).
// 
// The AND operator (represented as & in Go) returns 1 only if both bits are 1.
// 
// Here’s all possibilities of AND-ing two bits:
// 
//   | bit1 | 0  | 0  | 1  | 1 |
//   | ---  | -- | -- | -- | --|
//   | bit2 | 0  | 1  | 0  | 1 |
//   | ---  | -- | -- | -- | --|
//   | AND  | 0  | 0  | 0  | 1 |
// 
// The leftshift operator takes all the bits in a byte and shifts them 1 to the
// left.
// 
// The bit on the very right (in th 2^6 position) is dropped. The bit on the very
// right (in the 2^0 position) is filled in with a 0.
// 
//   | input  | 1 | 1  | 0  | 0 | 0 | 1 | 1 |
// 
//   LEFTSHIFT (<<)
// 
//   | output | 1 | 0  | 0 | 0 | 1 | 1 | 0 |
// 
// 
// The Hamming Distance
// 
// As the challenge tells us:
// 
//   The Hamming distance is just the number of differing bits. The distance
//   between:
// 
//     this is a test
// 
//   and
// 
//     wokka wokka!!!
// 
//   is 37. Make sure your code agrees before you proceed.
// 
// Or from https://en.wikipedia.org/wiki/Hamming_distance:
// 
//   In information theory, the Hamming distance between two strings of equal length
//   is the number of positions at which the corresponding symbols are different. In
//   other words, it measures the minimum number of substitutions required to change
//   one string into the other, or the minimum number of errors that could have
//   transformed one string into the other.
// 
// 
// In order to compute the hamming distance for those two strings, we need to:
// 
// 1. Convert each letter from ASCII to an 8-bit binary number.
// 
//    ASCII | Byte / Decimal | Binary
//   ----------------------------------
//    t     |      84        | 01110100
//    w     |     119        | 01110111
// 
// 
// 2. Line up each bit of the binary numbers against each other, and compare them.
// 
//    t    | 0 | 1 | 1 | 1 | 0 | 1 | 0 | 0
//    w    | 0 | 1 | 1 | 1 | 0 | 1 | 1 | 1
//   ----------------------------------------
//         | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ☓ | ☓
// 
// 3. Count up each bit that is different. That is the hamming distance. If we
//    could the ☓'s, we will see that the hamming distance is 2.
// 
// https://en.wikipedia.org/wiki/Hamming_distance#Algorithm_example provides an
// algorithm for the hamming distance in Python.
// 
// I wrote this in Go at
// https://pkg.go.dev/github.com/julieqiu/cryptopals/set1/c6#HammingDistance.
//
//
package challenge6