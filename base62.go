// Package base62 provides functions to convert between uint64 numbers and base62 encoded strings
package base62

import (
	"fmt"
	"math"
)

// Base62 alphabet, mapped from number to string, and from rune (char) to number
var numberToString = [62]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var runeToNumber = map[rune]uint64{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'a': 10, 'b': 11, 'c': 12, 'd': 13, 'e': 14, 'f': 15, 'g': 16, 'h': 17, 'i': 18, 'j': 19, 'k': 20, 'l': 21, 'm': 22, 'n': 23, 'o': 24, 'p': 25, 'q': 26, 'r': 27, 's': 28, 't': 29, 'u': 30, 'v': 31, 'w': 32, 'x': 33, 'y': 34, 'z': 35, 'A': 36, 'B': 37, 'C': 38, 'D': 39, 'E': 40, 'F': 41, 'G': 42, 'H': 43, 'I': 44, 'J': 45, 'K': 46, 'L': 47, 'M': 48, 'N': 49, 'O': 50, 'P': 51, 'Q': 52, 'R': 53, 'S': 54, 'T': 55, 'U': 56, 'V': 57, 'W': 58, 'X': 59, 'Y': 60, 'Z': 61}

// ToB62 accepts a uint64 number as its sole argument, and returns a base62 encoded string
func ToB62(i uint64) (encodedString string) {
	// Hard coded value to stop floating point errors
	if i == 0 {
		return "0"
	}

	// Start with empty string
	s := ""

	// Using logarithms, find the exponent of 61 that gives us i. Floor it to find most significant bit
	mostSignificantBit := math.Floor(math.Log(float64(i)) / math.Log(62.0))
	var value uint64
	var multiple uint64
	// Begin loop, run until we have calculated least significant bit
	for bit := mostSignificantBit; bit >= 0; bit-- {
		// The multiple of this bit, calculated by taking 61 to the power of x
		multiple = uint64(math.Pow(62, bit))

		// Determine how many times the value of the bit position can wholly go into our remaining number, this quotient is the value of this bit's position
		value = i / multiple

		// Remove the value represented by this bit position from the number using modulus
		i = i % multiple

		// Add the character of the base62 alphabet associated with this bit's multiple
		s += numberToString[value]
	}
	// After all of the bits have been run through, return the generated string
	return s
}

// FromB62 accepts a base62 encoded string, and returns the represented uint64 number and an optional error if invalid characters were found in the provided string
func FromB62(b62 string) (decodedNumber uint64, invalidError error) {
	// Get the length of the id
	length := len(b62)

	// Create array of runes from the string
	runes := []rune(b62)

	// Store the total value
	var total uint64

	// Loop through runes starting from the last
	for index := 0; index < length; index++ {
		// Check if the rune exists in the alphabet, if it doesn't return an error
		if valueOfCharacter, ok := runeToNumber[runes[index]]; ok {
			total += valueOfCharacter * uint64(math.Pow(62.0, float64(length-index-1)))
		} else {
			return 0, fmt.Errorf("Invalid base62 character at index %d", index)
		}
	}
	return total, nil
}
