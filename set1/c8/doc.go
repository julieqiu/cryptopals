// This file is generated using cmd/makedoc. DO NOT EDIT.
// To update, edit the doc.txt file in this directory.
// Then run
//     go run ./cmd/makedoc 1 8
//

// Set 1 Challenge 8
// https://cryptopals.com/sets/1/challenges/8
// 
// Detect AES in ECB mode
// 
// In this file (https://cryptopals.com/static/challenge-data/8.txt) are a bunch
// of hex-encoded ciphertexts.
// 
// One of them has been encrypted with ECB.
// 
// Detect it.
// 
// Remember that the problem with ECB is that it is stateless and deterministic;
// the same 16 byte plaintext block will always produce the same 16 byte
// ciphertext.
//
//
package c8