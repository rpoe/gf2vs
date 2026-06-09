// Copyright 2024 Ralf Poeppel. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gf2vs is implementing the type vector space of GF(2) the Galois Field of order 2.
// It is sometimes called bit array (also known as bit map, bit set, bit string, or bit vector).
// The vectors are defined as special type.
// In addition to math/bits it implements functions of the vector space of a given size.
// Each vector is constraint to the vector space given at creation time. The unit vectors
// are considered the base of the vector space. There are functions for verifying is a vector
// a base vector. The boolean operations and the vector operations are implemented.
// The count of ones is considered the norm of the vectors. It is the l_1 norm, or hamming weight
// of the vector. Some times this function is named popcount. It is the result of the scalar product.
// Sub vector spaces are supported too. A set of vectors may be a span of a sub vector space. This
// property is verified.
// This file holds functions limited to 16 bit, uint16.
package gf2vs

import (
	"fmt"
	"math/bits"
)

// The maximum size supported.
const uintSize16 = uint16(16)

// GF2VectorSpace16 represents a vector space of size n over GF(2).
type GF2VectorSpace16 struct {
	dim  uint16   // dimension of the vector space
	ones uint16   // bitvector where all bit are set, max value, zero not needed
	base []uint16 // array of the base vectors
}

// NewGF2VectorSpace16 create a vector space of dimension n.
// Return a pointer, as only a pointer has a null value, but a struct not.
// Panic if n is out of range.
func NewGF2VectorSpace16(n uint16) *GF2VectorSpace16 {
	if n < 1 {
		panic(fmt.Sprintf("NewGF2VectorSpace16(dim): dim = %v < 1", n))
	}
	if n > uintSize16 {
		panic(fmt.Sprintf("NewGF2VectorSpace16(dim): dim = %v > %v = uintSize16", n, uintSize16))
	}

	ones := uint16(1)
	ones = uint16(1<<n - 1)
	base := make([]uint16, n)
	base[0] = uint16(1)
	for i := uint16(2); i < n; i++ {
		base[i] = base[i-1] << 1
	}

	sp := GF2VectorSpace16{n, ones, base}
	return &sp
}

func (sp *GF2VectorSpace16) String() string {
	return fmt.Sprintf("GF(2)sp{%v: %v}", sp.dim, sp.ones)
}

// GF2SubVectorSpace16 represents a sub vector space
type GF2SubVectorSpace16 struct {
	GF2VectorSpace16
	subOnes uint16 // bitvector where the bits of the base are set
}

// NewGF2SubVectorSpace16 create a sub vector space with base bits b as sub space
// of a vector space with dimension n.
// Panic if b > n.
// We call internally NewGF2VectorSpace16 which may panic for n out of range.
func NewGF2SubVectorSpace16(n, b uint16) *GF2SubVectorSpace16 {
	if b > n {
		panic(fmt.Sprintf("NewGF2SubVectorSpace16(dim): base %v not in space with dim = %v", b, n))
	}
	vs := NewGF2VectorSpace16(n)
	svs := GF2SubVectorSpace16{*vs, b}
	return &svs
}

func (sp *GF2SubVectorSpace16) String() string {
	return fmt.Sprintf("GF(2)ssp{%v: %v, %v}", sp.dim, sp.ones, sp.subOnes)
}

// GF2Vector16 represents a vector in GF(2^n) a bitvector of len n,
// in a vector space of dim n.
type GF2Vector16 struct {
	sp  *GF2VectorSpace16 // the space of this vector
	val uint16            // value of the vector
}

// Val return the value as int.
// If no valid v is given we return 0.
func (v *GF2Vector16) Val() uint16 {
	if v == nil {
		return 0
	}
	return v.val
}

// String returns a string representing
func (v *GF2Vector16) String() string {
	return fmt.Sprintf("%0[1]*[2]b", v.sp.dim, v.val)
}

// NewGF2Vector16 create a vector with value in vector space,
// value must be greater equal 0.
func (s *GF2VectorSpace16) NewGF2Vector16(value uint16) *GF2Vector16 {
	if value > s.ones {
		panic(fmt.Sprintf("NewGF2Vector16(value): value = %v > %v", value, s.ones))
	}

	v := GF2Vector16{s, value}
	return &v
}

// GF2BaseVector return a GF2Vector16 representing the base with index i.
// Panic if i is out of range.
func (s *GF2VectorSpace16) GF2BaseVector(i uint16) *GF2Vector16 {
	if i == 0 || s.dim < i {
		panic(fmt.Sprintf("GF2BaseVector(i): i = %v out of range [1, %v]", i, s.dim))
	}
	v := uint16(1) << (i - 1)
	b := GF2Vector16{s, v}
	return &b
}

// GF2BaseVector return the base vector representing the base with index i in the same vector space.
func (v *GF2Vector16) GF2BaseVector(i uint16) *GF2Vector16 {
	return v.sp.GF2BaseVector(i)
}

// GF2Zeros return a GF2Vector16 where all bits are unset.
func (s *GF2VectorSpace16) GF2Zeros() *GF2Vector16 {
	b := GF2Vector16{s, 0}
	return &b
}

// GF2Ones return a GF2Vector16 where dim bits are set.
func (s *GF2VectorSpace16) GF2Ones() *GF2Vector16 {
	b := GF2Vector16{s, s.ones}
	return &b
}

// IsZeros return true if all bits are unset.
func (v *GF2Vector16) IsZeros() bool {
	return v.val == 0
}

// IsOnes return true if all bits are set.
func (v *GF2Vector16) IsOnes() bool {
	return v.val == v.sp.ones
}

// Index return the index of the coordinate of a base vector.
// Use https://graphics.stanford.edu/~seander/bithacks.html#DetermineIfPowerOf2
// Index is zero and isBase is false if v is no base vector.
func (v *GF2Vector16) Index() (index uint16, isBase bool) {
	c := v.val
	if (c > 0) && (c&(c-1)) == 0 {
		return uint16(bits.Len16(c)), true
	}
	return
}

// IsBaseVector return true if v is a base vector.
// Use https://graphics.stanford.edu/~seander/bithacks.html#DetermineIfPowerOf2
func (v *GF2Vector16) IsBaseVector() bool {
	c := v.val
	return (c > 0) && (c&(c-1)) == 0
}

// Copy return a copy of x, sharing the same vector space.
func (x *GF2Vector16) Copy() *GF2Vector16 {
	// we allow zero value
	if x == nil {
		return x
	}
	c := GF2Vector16{x.sp, x.val}
	return &c
}

// Not16 returns ^x, the negation of x.
func Not16(x *GF2Vector16) *GF2Vector16 {
	b := GF2Vector16{x.sp, x.sp.ones ^ x.val}
	return &b
}

// And16 return x_1 & x_2 & ...,
// panic if x_i and x_j are of vector spaces of different dimension.
func And16(x ...*GF2Vector16) *GF2Vector16 {
	n := len(x)
	z := *x[0]
	for i := 1; i < n; i++ {
		y := x[i]
		if z.sp.dim != y.sp.dim {
			panic(fmt.Sprintf("And: incompatible vector spaces: "+
				"z.dim = %v != %v = y.dim",
				z.sp.dim, y.sp.dim))
		}
		z.val &= y.val
	}
	return &z
}

// Or16 return x_1 | x_2 | ...,
// panic, if x_i and x_j are of vector spaces of different dimension.
func Or16(x ...*GF2Vector16) *GF2Vector16 {
	n := len(x)
	z := *x[0]
	for i := 1; i < n; i++ {
		y := x[i]
		if z.sp.dim != y.sp.dim {
			panic(fmt.Sprintf("Or16: incompatible vector spaces: "+
				"z.dim = %v != %v = y.dim",
				z.sp.dim, y.sp.dim))
		}
		z.val |= y.val
	}
	return &z
}

// Xor16 return x_1 | x_2 | ...,
// panic, if x_i and x_j are of vector spaces of different dimension.
func Xor16(x ...*GF2Vector16) *GF2Vector16 {
	n := len(x)
	z := *x[0]
	for i := 1; i < n; i++ {
		y := x[i]
		if z.sp.dim != y.sp.dim {
			panic(fmt.Sprintf("Xor16: incompatible vector spaces: "+
				"z.dim = %v != %v = y.dim",
				z.sp.dim, y.sp.dim))
		}
		z.val ^= y.val
	}
	return &z
}

// ComplementOr16 return z = Not16(Or16(x) = ^(x_1 | x_2 | ...),
// This can be used to "subtract" the Or16(x) from Ones.
func ComplementOr16(x ...*GF2Vector16) *GF2Vector16 {
	z := Or16(x...)
	return Not16(z)
}

// ComplementXor16 return z = Not16(Xor16(x)) = ^(x_1 ^ x_2 ^ ...),
// This can be used to subtract the sum of x = Xor16(x) from Ones.
func ComplementXor16(x ...*GF2Vector16) *GF2Vector16 {
	z := Xor16(x...)
	return Not16(z)
}

// MaskBits16 return z = And16(x, m) = x & m.
func MaskBits16(x, m *GF2Vector16) *GF2Vector16 {
	return And16(x, m)
}

// ClearBits16 return z = And16(x, Not16(m)) = x & ^m.
func ClearBits16(x, m *GF2Vector16) *GF2Vector16 {
	return And16(x, Not16(m))
}

// SetBits16 return z = Or16(x, m) = x | m.
func SetBits16(x, m *GF2Vector16) *GF2Vector16 {
	return Or16(x, m)
}

// ToggleBits16 return z = Xor16(x, m) = x ^ m.
func ToggleBits16(x, m *GF2Vector16) *GF2Vector16 {
	return Xor16(x, m)
}

// OnesCount16 returns the number of one bits ("population count") in x.
func OnesCount16(x *GF2Vector16) int {
	return bits.OnesCount16(x.val)
}

// ScalarProduct16 returns the scalar product of 2 vectors, which is the norm, the OnesCount of the product vector.
func ScalarProduct16(a, b *GF2Vector16) int {
	prod := And16(a, b)
	return OnesCount16(prod)
}

// SpanOfSubspace16 returns true and the subspace of a set of vectors s if the
// vectors span a subspace. For true the dimension of the set s must be
// greater or equal the Norm of the union of the set.
// Steinitz exchange lemma:
// https://en.wikipedia.org/w/index.php?title=Steinitz_exchange_lemma&oldid=1336271854
// The dimension of each span of a vector space is greater or equal to the
// dimension of the base of the vector space.
func SpanOfSubspace16(s []*GF2Vector16) (Ok bool, sp *GF2SubVectorSpace16) {
	span := Or16(s...) // demands all s are in same vectorspace
	dim := OnesCount16(span)
	if dim <= len(s) {
		// we have a subspace
		svs := GF2SubVectorSpace16{(*s[0].sp), span.val}
		return true, &svs
	}
	return
}

// SubSpace16OfSet check all combinations of subsets of a set of bit vectors
// to see wether a subset is confined to a subspace.
// Check only bit vectors with more than 1 bit set.
// All bit vectors with one bit for a subset.
// Return the first subsets and subspaces found.
func SubSpaceOfSets16() {
}
