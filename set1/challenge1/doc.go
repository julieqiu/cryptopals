// This file is generated using cmd/makedoc. DO NOT EDIT.
// To update, edit the doc.txt file in this directory.
// Then run
//     go run ./cmd/makedoc 1 1
//

// This package provides a solution to https://cryptopals.com/sets/1/challenges/1.
// 
// Problem
// 
// Convert hex to base64
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
// This challenge is all about learning different ways to represent a given number.
// 
// We'll the number 81 to go through the basics. Before starting the problem,
// you'll need to understand why:
// 
//     81        in decimal
//   = 1010001   in binary
//   = 51        in hexidecimal
//   = Q         in ASCII
//   = UQ==      in base64
// 
// 
// Decimal
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
// Binary
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
// Hexidecimal
// 
// Hexidecimal uses base 16 rather than base 10 or base 2. It represents numbers
// 10-16 using letters.
// 
// https://kb.iu.edu/d/afdl shows a table of what a number looks like in decimal,
// hexidecimal, and binary.
// 
// https://www.wikihow.com/Convert-Binary-to-Hexadecimal provides an explanation
// of how to convert binary to hexidecimal.
// 
// 
// To get 81 in hexidecimal:
// 
//     ? * 16^1 (16)
//   + ? * 16^0 (1)
//   ----------
//   = 81
// 
// Computing the `?` gives us:
// 
//     5 * 16^1 (16)
//   + 1 * 16^0 (1)
//   ----------
//   = 81 (decimal)
//   = 1010001 (binary)
//   = 51 (hex)
// 
// 
// ASCII
// 
// From https://en.wikipedia.org/wiki/ASCII:
// 
//   ASCII is the American Standard Code for Information Interchange, which is a
//   character encoding standard for electronic communication. ASCII codes represent
//   text in computers, telecommunications equipment, and other devices. Most modern
//   character-encoding schemes are based on ASCII, although they support many
//   additional characters.
// 
// You can think of it as another representation of numbers, where characters map
// to the table at https://www.lookuptables.com/text/ascii-table.
// 
// 81 in ASCII is `Q`.
// 
// 
// Base64
// 
// From https://en.wikipedia.org/wiki/Base64:
// 
//   Base64 is a group of binary-to-text encoding schemes that represent binary
//   data (more specifically, a sequence of 8-bit bytes) in an ASCII string format
//   by translating the data into a radix-64 representation. The term Base64
//   originates from a specific MIME content transfer encoding. Each non-final
//   Base64 digit represents exactly 6 bits of data. Three 8-bit bytes (i.e., a
//   total of 24 bits) can therefore be represented by four 6-bit Base64 digits.
// 
// As a reminder:
// 
//   81 in decimal
//   = 01010001 in binary
//   = 51 in hex
//   = Q in ASCII
// 
// Using the binary representation of 81, we can convert it to base64 by:
// 
// 1. Divide it into sextets:
// 
//   010100 01----  // 01010001 in sextets
// 
// 
// 2. Replace missing digits with 0s.
// 
//   010100 010000 // 010100 01---- from above, filled with 0s
// 
// 
// 3. Use the table at https://en.wikipedia.org/wiki/Base64#Base64_table to map a
// number to its base64 representation.
// 
//   010100 010000 = UQ
// 
// 4. Lastly, base64 must ALWAYS consist of 4 characters. In our case, we have
// `UQ`, so we pad the remaining characters with `=`, giving us
// 
//   UQ==
// 
// 
// Manual Solution
// 
// Going back to the original challenge, our goal is to get from Hex to base64.
// Now that we know how to convert hex to binary to base64, we can even do this
// manually!
// 
//   49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
// 
// to
// 
//   SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
// 
// 
// The Manual Solution
// 
// We can use the process above to get:
// 
// Hex
// 
//   49276d
// 
// Binary
// 
//   0100 1001 0010 0111 0110 1101
// 
// Base64
// 
//   // Binary in sextets: | 010010 | 010010 | 011101 | 101101 |
//   // Base64             | S      | S      | d      | t      |
// 
// There we go! The first 4 characters of the output string that we want!
// 
// 
// Code Solution
// 
// Now that we know what all these things mean, we can use these Go functions from
// the standard library to help us:
// 
// - https://pkg.go.dev/encoding/hex#Decode
// - https://pkg.go.dev/encoding/base64#Encoding.Encode
//
//
package challenge1