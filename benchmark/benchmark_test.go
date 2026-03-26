// Ralf Poeppel, 2026

// Bnechmark tests to select library functions for implementation
package gf2vs

import (
	"math/bits"
	"testing"
)

func BenchmarkLen(b *testing.B) {
	var j uint
	for b.Loop() {
		for j = 1; j < 64; j <<= 1 {
			// source of Len is smaller
			// as source of LeadingZeros
			bits.Len(j)
		}
	}
}

func BenchmarkLeadingZeros(b *testing.B) {
	var j uint
	for b.Loop() {
		for j = 1; j < 64; j <<= 1 {
			// source of Len is smaller
			// as source of LeadingZeros
			bits.LeadingZeros(j)
		}
	}
}
