// Simple program to retrieve the processor related constants for reference
// Ralf Poeppel, 2026
package main

import (
	"fmt"
	"math/big"
	"math/bits"
)

func main() {
	fmt.Printf("uint has size %v\n\n", bits.UintSize)
	fmt.Println("Max value supported")
	largest := int(^uint(0) >> 1)
	maxValue := uint(1)<<(bits.UintSize-1) - 1
	largestInt := big.NewInt(int64(largest))
	largestBits := largestInt.BitLen()
	maxFromInt := int(largestInt.Int64())
	fmt.Println(largest)
	fmt.Println(maxValue)
	fmt.Println(largestInt)
	fmt.Println(maxFromInt)
	fmt.Println(largestBits, "# bits")

	// Initialize twoWord as 10^21, an big/Int integer with 2 words.
	var twoWord big.Int
	twoWord.Exp(big.NewInt(10), big.NewInt(21), nil)
	fmt.Printf("\nInt has decimal value %v\n", &twoWord)

	txt := twoWord.Text(2)
	twoWordBits := twoWord.Bits()

	// The Scan function is rarely used directly;
	// the fmt package recognizes it as an implementation of fmt.Scanner.
	v := new(big.Int)
	_, err := fmt.Sscanf(txt, "%b", v)
	if err != nil {
		fmt.Println("error scanning value:", err)
		return
	}

	fmt.Printf("scanned value%31d\n\n", v)

	fmt.Printf("binary value %v\n", txt)
	fmt.Printf("scanned     %71s\n", v.Text(2))

	fmt.Printf("word{0]            %064b\n", twoWordBits[0])
	fmt.Printf("word[1]     %7b\n", twoWordBits[1])

	// see https://de.wikipedia.org/wiki/Byte-Reihenfolge
	fmt.Printf("Speicherung ist Little Endian in %v words, length is %v Bits\n\n", len(twoWordBits), twoWord.BitLen())

	var ntTwWrd big.Int
	notTwoWord := &ntTwWrd
	notTwoWord.Not(&twoWord)
	fmt.Printf("org  %v %v\n", twoWord, &twoWord)
	fmt.Printf("not  %v %v\n", ntTwWrd, notTwoWord)
	var andNt big.Int
	andNot := &andNt
	andNot.AndNot(&twoWord, &twoWord)
	fmt.Printf("&^ 1 %v %v\n", andNt, andNot)
	andNot.And(&twoWord, notTwoWord)
	fmt.Printf("&^ 2 %v %v\n", andNt, andNot)

	var ngTwWrd big.Int
	negTwoWord := &ngTwWrd
	negTwoWord.Not(&twoWord)
	fmt.Printf("org  %v %v\n", twoWord, &twoWord)
	fmt.Printf("neg  %v %v\n", ngTwWrd, negTwoWord)

	ntTwWrdBits := (&ntTwWrd).Bits()
	fmt.Printf("word{0]            %064b\n", ntTwWrdBits[0])
	fmt.Printf("word[1]     %7b\n", ntTwWrdBits[1])

	var xrTwWrd big.Int
	xorTwoWord := &xrTwWrd
	xorTwoWord.Xor(&twoWord, notTwoWord)
	fmt.Println((&xrTwWrd).Sign(), xrTwWrd)
	(&xrTwWrd).Not(&xrTwWrd)
	fmt.Println((&xrTwWrd).Sign(), xrTwWrd)
	//(&xrTwWrd).Not(&xrTwWrd)
	fmt.Println((&xrTwWrd).Cmp(big.NewInt(0)))

	fmt.Println("SetBit")
	var OnBt big.Int
	OneBit := &OnBt
	OneBit.SetBit(OneBit, 0, 1)
	fmt.Println(0, OneBit.Sign(), OnBt)
	OneBit.SetBit(OneBit, 1, 1)
	fmt.Println(1, OneBit.Sign(), OnBt)
	for i := 63; i <= 65; i++ {
		OneBit.SetBit(OneBit, i, 1)
		fmt.Println(i, OneBit.Sign(), OnBt)
		//OneBit.SetBit(OneBit, i, 0)
	}
}
