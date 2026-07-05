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
package gf2vs

import (
	"fmt"
	"math/bits"
)

// GF2VectorSpace represents a vector space of size n over GF(2).
type GF2VectorSpace struct {
	dim  uint // dimension of the vector space
	ones uint // bitvector where all bit are set, no zeros as allways 0
}

// NewGF2VectorSpace create a vector space of dimension n.
// Return a pointer, as only a pointer has a null value, but a struct not.
// Panic if n is out of range.
func NewGF2VectorSpace(n uint) *GF2VectorSpace {
	if n < 1 {
		panic(fmt.Sprintf("NewGF2VectorSpace(dim): dim = %v < 1", n))
	}
	if n > bits.UintSize {
		panic(fmt.Sprintf("NewGF2VectorSpace(dim): dim = %v > %v = bits.UintSize", n, bits.UintSize))
	}

	ones := uint(1)
	for i := uint(1); i < n; i++ {
		ones <<= 1
		ones += 1
	}

	sp := GF2VectorSpace{n, ones}
	return &sp
}

func (sp *GF2VectorSpace) String() string {
	return fmt.Sprintf("GF(2)sp{%v: %v}", sp.dim, sp.ones)
}

// GF2VectorSubspace represents a vector subspace
type GF2VectorSubspace struct {
	GF2VectorSpace
	subOnes uint // bitvector where the bits of the base are set
}

// NewGF2VectorSubspace create a sub vector space with base bits b as sub space
// of a vector space with dimension n.
// Panic if b > n.
// We call internally NewGF2VectorSpace which may panic for n out of range.
func NewGF2VectorSubspace(n, b uint) *GF2VectorSubspace {
	if b > n {
		panic(fmt.Sprintf("NewGF2VectorSubspace(dim): base %v not in space with dim = %v", b, n))
	}
	vs := NewGF2VectorSpace(n)
	svs := GF2VectorSubspace{*vs, b}
	return &svs
}

func (sp *GF2VectorSubspace) String() string {
	return fmt.Sprintf("GF(2)ssp{%v: %v, %v}", sp.dim, sp.ones, sp.subOnes)
}

// GF2vector represents a vector in GF(2^n) a bitvector of len n,
// in a vector space of dim n.
type GF2Vector struct {
	sp  *GF2VectorSpace // the space of this vector
	val uint            // value of the vector
}

// Val return the value as int.
// If no valid v is given we return 0.
func (v *GF2Vector) Val() uint {
	if v == nil {
		return 0
	}
	return v.val
}

// String returns a string representing
func (v *GF2Vector) String() string {
	return fmt.Sprintf("%0[1]*[2]b", v.sp.dim, v.val)
}

// NewGF2Vector create a vector with value in vector space,
// value must be greater equal 0.
func (s *GF2VectorSpace) NewGF2Vector(value uint) *GF2Vector {
	vmx := (uint(1) << s.dim) - 1
	if value > vmx {
		panic(fmt.Sprintf("NewGF2Vector(value): value = %v > %v", value, vmx))
	}

	v := GF2Vector{s, value}
	return &v
}

// NewGF2VectorSet return a set of vectors with the values given.
func (s *GF2VectorSpace) NewGF2VectorSet(u []uint) []*GF2Vector {
	set := make([]*GF2Vector, len(u))
	for i, x := range u {
		set[i] = s.NewGF2Vector(x)
	}
	return set
}

// GF2BaseVector return a GF2Vector representing the base with index i.
// Panic if i is out of range.
func (s *GF2VectorSpace) GF2BaseVector(i uint) *GF2Vector {
	if i == 0 || s.dim < i {
		panic(fmt.Sprintf("GF2BaseVector(i): i = %v out of range [1, %v]", i, s.dim))
	}
	v := uint(1) << (i - 1)
	b := GF2Vector{s, v}
	return &b
}

// BaseVector return the base vector representing the base with index i in the same vector space.
func (v *GF2Vector) GF2BaseVector(i uint) *GF2Vector {
	return v.sp.GF2BaseVector(i)
}

// GF2Zeros return a GF2Vector where all bits are unset.
func (s *GF2VectorSpace) GF2Zeros() *GF2Vector {
	b := GF2Vector{s, 0}
	return &b
}

// GF2Ones return a GF2Vector where dim bits are set.
func (s *GF2VectorSpace) GF2Ones() *GF2Vector {
	b := GF2Vector{s, s.ones}
	return &b
}

// IsZeros return true if all bits are unset.
func (v *GF2Vector) IsZeros() bool {
	return v.val == 0
}

// IsOnes return true if all bits are set.
func (v *GF2Vector) IsOnes() bool {
	return v.val == v.sp.ones
}

// Index return the index of the coordinate of a base vector.
// Use https://graphics.stanford.edu/~seander/bithacks.html#DetermineIfPowerOf2
// Index is zero and isBase is false if v is no base vector.
func (v *GF2Vector) Index() (index uint, isBase bool) {
	c := v.val
	if (c > 0) && (c&(c-1)) == 0 {
		return uint(bits.Len(c)), true
	}
	return
}

// IsBaseVector return true if v is a base vector.
// Use https://graphics.stanford.edu/~seander/bithacks.html#DetermineIfPowerOf2
func (v *GF2Vector) IsBaseVector() bool {
	c := v.val
	return (c > 0) && (c&(c-1)) == 0
}

// Copy return a copy of x, sharing the same vector space.
func (x *GF2Vector) Copy() *GF2Vector {
	// we allow zero value
	if x == nil {
		return x
	}
	c := GF2Vector{x.sp, x.val}
	return &c
}

// Not returns ^x, the negation of x.
func Not(x *GF2Vector) *GF2Vector {
	b := GF2Vector{x.sp, x.sp.ones ^ x.val}
	return &b
}

// And return x_1 & x_2 & ...,
// panic if x_i and x_j are of vector spaces of different dimension.
func And(x ...*GF2Vector) *GF2Vector {
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

// Or return x_1 | x_2 | ...,
// panic, if x_i and x_j are of vector spaces of different dimension.
func Or(x ...*GF2Vector) *GF2Vector {
	n := len(x)
	z := *x[0]
	for i := 1; i < n; i++ {
		y := x[i]
		if z.sp.dim != y.sp.dim {
			panic(fmt.Sprintf("Or: incompatible vector spaces: "+
				"z.dim = %v != %v = y.dim",
				z.sp.dim, y.sp.dim))
		}
		z.val |= y.val
	}
	return &z
}

// Xor return x_1 | x_2 | ...,
// panic, if x_i and x_j are of vector spaces of different dimension.
func Xor(x ...*GF2Vector) *GF2Vector {
	n := len(x)
	z := *x[0]
	for i := 1; i < n; i++ {
		y := x[i]
		if z.sp.dim != y.sp.dim {
			panic(fmt.Sprintf("Xor: incompatible vector spaces: "+
				"z.dim = %v != %v = y.dim",
				z.sp.dim, y.sp.dim))
		}
		z.val ^= y.val
	}
	return &z
}

// ComplementOr return z = Not(Or(x) = ^(x_1 | x_2 | ...),
// This can be used to "subtract" the Or(x) from Ones.
func ComplementOr(x ...*GF2Vector) *GF2Vector {
	z := Or(x...)
	return Not(z)
}

// ComplementXor return z = Not(Xor(x)) = ^(x_1 ^ x_2 ^ ...),
// This can be used to subtract the sum of x = Xor(x) from Ones.
func ComplementXor(x ...*GF2Vector) *GF2Vector {
	z := Xor(x...)
	return Not(z)
}

// MaskBits return z = And(x, m) = x & m.
func MaskBits(x, m *GF2Vector) *GF2Vector {
	return And(x, m)
}

// ClearBits return z = And(x, Not(m)) = x & ^m.
func ClearBits(x, m *GF2Vector) *GF2Vector {
	return And(x, Not(m))
}

// SetBits return z = Or(x, m) = x | m.
func SetBits(x, m *GF2Vector) *GF2Vector {
	return Or(x, m)
}

// ToggleBits return z = Xor(x, m) = x ^ m.
func ToggleBits(x, m *GF2Vector) *GF2Vector {
	return Xor(x, m)
}

// OnesCount returns the number of one bits ("population count") in x.
func OnesCount(x *GF2Vector) int {
	return bits.OnesCount(x.val)
}

// ScalarProduct returns the scalar product of 2 vectors, which is the norm, the OnesCount of the product vector.
func ScalarProduct(a, b *GF2Vector) int {
	prod := And(a, b)
	return OnesCount(prod)
}

// BaseOfOnesVector return the base, the set of unit vectors of the given bit vector
func BaseOfOnesVector(vc *GF2Vector) []*GF2Vector {
	vec := vc.Val()
	res := make([]*GF2Vector, 0, bits.Len(vec))
	for vec != 0 {
		b := bits.TrailingZeros(vec)
		v := uint(1) << b
		res = append(res, vc.sp.NewGF2Vector(v))
		vec -= v
	}
	return res
}

// OnesOfSet return the ones of the span and the dimension of the subspace
// of the vectors in the given set.
func OnesOfSet(set []*GF2Vector) (*GF2Vector, int) {
	span := Or(set...) // demands all s are in same vectorspace
	return span, OnesCount(span)
}

// SubspaceOfSet return first subspace of set, with dim.
// check all combinations of subsets of a set of bit vectors
// to see wether a subset is confined to a subspace.
// if found equals false sp is nil.
func SubspaceOfSet(set []*GF2Vector, dim int) (sp *GF2VectorSubspace, found bool) {
	// compute all combinations of subset of size dim
	subsets := combinations(set, dim)

	for _, s := range subsets {
		spOnes, spDim := OnesOfSet(s)
		if spDim == dim {
			spa := GF2VectorSubspace{*set[0].sp, spOnes.Val()}
			return &spa, true
		}
	}

	return
}

/* TODO
// isBaseElementOfValueInSet return true if any element of the set has the bits of value set
func isValueInSet(value uint, set []*GF2Vector) bool {
	if value == 0 {
		return false
	}
	ub := unitBits(uint(value))
	for _, d := range ub {
		for _, s := range set {
			cmpVal := grid[s] & d
			if cmpVal != 0 && cmpVal == d {
				return true
			}
		}
	}
	return false
}

// removeValueFromSet remove the digit(value) from the elements of grid in set
func removeValueFromSet(d, allDigits int, set, grid []int) {
	notd := allDigits ^ d
	for _, s := range set {
		grid[s] &= notd
	}
}
*/

/////////////////////////////////////// utility functions /////////////////////

// binominalCoefficient compute the binomial coefficient (n over k)
// = \prod_{j=1}^k (n+1-j)/j
func binominalCoefficient(n, k uint) uint {
	uOne := uint(1)

	nmk := n - k
	if k > nmk {
		k = nmk
		nmk = n - k
	}

	b := uOne
	nom := n + uOne
	for j := uOne; j <= k; j++ {
		b = b * (nom - j) / j
	}
	return b
}

// combinations return all k-element subsets from elements
func combinations[T any](elements []T, k int) [][]T {
	// prevent reallocations of result and comb, by using final length
	est := binominalCoefficient(uint(len(elements)), uint(k))
	result := make([][]T, 0, est)
	comb := make([]T, 0, k)

	// recursive closure using result and comb
	var recursiveCombination func(start int)
	recursiveCombination = func(start int) {
		if len(comb) == k {
			// append copy of comb to result, comb is modified for each run
			temp := make([]T, len(comb))
			copy(temp, comb)
			result = append(result, temp)
			return
		}

		for i := start; i < len(elements); i++ { // iterate remaining elements
			comb = append(comb, elements[i]) // append current element
			recursiveCombination(i + 1)      // recurse with remaining elements
			comb = comb[:len(comb)-1]        // remove last element added at loop start for next run
		}
	}

	recursiveCombination(0)
	return result
}
