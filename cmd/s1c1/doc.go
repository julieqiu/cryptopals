// This file is generated using cmd/makedoc. DO NOT EDIT.
// To update, edit the doc.txt file in this directory.
// Then run
//     go run ./cmd/makedoc s1c1
//

// Set 1 Challenge 1
// 
// Convert hex to base64
// 
// The first challenge is to convert hex to 64.
// 
// The string:
// 
//     49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
// 
// Should produce:
// 
//     SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
// 
// So go ahead and make that happen. You'll need to use this code for the rest of the exercises.
// 
// Cryptopals Rule: Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.
// 
// 
// The Basics
// 
// Concepts you'll need to know for this challenge:
// 
// - Hexidecimal
// - Binary
// - Decimal
// - Base64
// - ASCII
// 
// 
// Decimal (Base 10)
// 
// Decimal is the notation we are used to seeing numbers represented in in every
// day life — it uses 10 as the base notation. For example, when you see the
// number 81, it means:
// 
//     8 * 10^1 (8)
//   + 1 * 10^0 (1)
//   ----------
//   = 81
// 
// 
// Binary (Base 2)
// 
// Rather than using base 10, in binary, we will use base 2. So 81 looks like
// this:
// 
//     ? * 2^6 (64)
//   + ? * 2^5 (32)
//   + ? * 2^4 (16)
//   + ? * 2^3 (8)
//   + ? * 2^2 (4)
//   + ? * 2^1 (2)
//   + ? * 2^0 (1)
//   ----------
//   = 81
// 
// Computing the `?` gives us:
// 
//     1 * 2^6 (64)
//   + 0 * 2^5 (32)
//   + 1 * 2^4 (16)
//   + 0 * 2^3 (8)
//   + 0 * 2^2 (4)
//   + 0 * 2^1 (2)
//   + 1 * 2^0 (1)
//   ----------
//   = 81 (decimal)
//   = 1010001 (binary)
//
//
package main