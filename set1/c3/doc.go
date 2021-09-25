// This file is generated using cmd/makedoc. DO NOT EDIT.
// To update, edit the doc.txt file in this directory.
// Then run
//     go run ./cmd/makedoc 1 3
//

// Set 1 Challenge 3
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
// There are no new concepts for this challenge.
// 
// 
// Solution
// 
// Like in challenge 2, start by decoding the string to bytes.
//
//
package c3