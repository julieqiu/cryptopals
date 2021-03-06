// This file is generated using cmd/makedoc. DO NOT EDIT.
// To update, edit the doc.txt file in this directory.
// Then run
//     go run ./cmd/makedoc 1 3
//

// This package provides a solution to https://cryptopals.com/sets/1/challenges/2.
// 
// Single-byte XOR cipher
// 
// The hex encoded string:
// 
//   1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
// 
// ... has been XOR'd against a single character. Find the key, decrypt the
// message.
// 
// You can do this by hand. But don't: write code to do it for you.
// 
// How? Devise some method for "scoring" a piece of English plaintext. Character
// frequency is a good metric. Evaluate each output and choose the one with the
// best score.
// 
// Achievement Unlocked: You now have our permission to make "ETAOIN SHRDLU" jokes on Twitter.
// 
// 
// Basics
// 
// 
// Decryption: XOR has a property - if
// 
//   a = b ^ c
//   // then
//   b = a ^ c
// 
// As a result, the decryption process is exactly the same as the encryption.e.
// 
// We iterate through the encrypted message bytewise and XOR each byte with the
// encryption key - the resultant will be the original message.
// 
// ETAOIN SHRDLU is the approximate order of frequency of the 12 most commonly used letters
// in the English language. https://en.wikipedia.org/wiki/Etaoin_shrdlu
// 
// 
// Solution
// 
// Like in challenge 2, start by decoding the string to bytes.
//
//
package challenge3