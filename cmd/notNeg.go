// Simple program to check behavior of Not and Neg of math/big Int
// Ralf Poeppel, 2026
package main

import (
	"fmt"
	"math/big"
)

func main() {
	seven := big.NewInt(7)
	lenSeven := seven.BitLen()
	negSeven := new(big.Int).Neg(seven)
	lenNegSeven := negSeven.BitLen()
	notSeven := new(big.Int).Not(seven)
	lenNotSeven := notSeven.BitLen()
	fmt.Println("Val", seven, *seven, "has", lenSeven, "Bits")
	fmt.Println("Neg", negSeven, *negSeven, "has", lenNegSeven, "Bits")
	fmt.Println("Not", notSeven, *notSeven, "has", lenNotSeven, "Bits")

	// Initialize twoWord as 10^21, an big/Int integer with 2 words.
	var twWrd big.Int // allocate variable
	twoWord := &twWrd // retrieve pointer
	twoWord.Exp(big.NewInt(10), big.NewInt(21), nil)
	fmt.Printf("Int decimal value %v, struct %v\n", twoWord, twWrd)
	negTwoWord := new(big.Int).Neg(twoWord)
	fmt.Printf("Neg decimal value %v, struct %v\n", negTwoWord, *negTwoWord)
	notTwoWord := new(big.Int).Not(twoWord)
	fmt.Printf("Not decimal value %v, struct %v\n", notTwoWord, *notTwoWord)
	fmt.Printf("Not hex value %x, struct %v\n", notTwoWord, *notTwoWord)
}
