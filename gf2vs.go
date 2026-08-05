// Copyright 2024 Ralf Poeppel. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gf2vs is implementing the type vector space of GF(2) the Galois Field of order 2.
// It is sometimes called bit array (also known as bit map, bit set, bit string, or bit vector).
// The vectors are defined as special type.
// In addition to math/bits it implements functions of the vector space of a given size.
// Each vector is constraint to the vector space given at creation time.
// The operation complement has only meaning for a vector space of a dedicated size.
// The unit vectors are considered the base of the vector space.
// There are functions for verifying is a vector a base vector.
// The boolean operations and the vector operations are implemented.
// The count of ones is considered the norm of the vectors. It is the l_1 norm, or hamming weight
// of the vector. Some times this function is named popcount. It is the result of the scalar product.
// The scalar product deliveres the count of common bits of 2 vectors.
// Sub vector spaces are supported. Each set of vectors spans a sub vector space.
// If the size of the set equals the dimension of the subspace. The set is
// confined to a subspace. If a set is partinioned into 2 subsets and one of
// the subsets is confined to a subspace, the second subset can be restricted to
// the complement subspace of the vector space. Then the direct sum of the vector spaces exists.
// The function RestrictToComplement is implementing this algorithmus.
package gf2vs

import (
	"fmt"
	"math/bits"
)

/////////////////////////////////////// Types /////////////////////////////////

// GF2VectorSpace represents a vector space of size n over GF(2).
type GF2VectorSpace struct {
	dim  uint // dimension of the vector space
	ones uint // bitvector where all bit are set, no zeros as allways 0
}

// GF2VectorSubspace represents a vector subspace
type GF2VectorSubspace struct {
	*GF2VectorSpace                 // embbed vector space, a subspace is a space
	container       *GF2VectorSpace // reference to the vector space containing this subspace
}

// GF2vector represents a vector in GF(2^n) a bitvector of len n,
// in a vector space of dim n.
type GF2Vector struct {
	sp  *GF2VectorSpace // the space of this vector
	val uint            // value of the vector
}

// GF2VectorSet a set of bit vectors
type GF2VectorSet []*GF2Vector

/////////////////////////////////////// GF2VectorSpace ////////////////////////

// NewGF2VectorSpace create a vector space of dimension n.
// Return a pointer, as only a pointer has a null value, but a struct not.
// Panic if n is out of range.
func NewGF2VectorSpace(n uint) *GF2VectorSpace {
	if n == 0 {
		return &GF2VectorSpace{n, n}
	}
	if n > bits.UintSize {
		panic(fmt.Sprintf("NewGF2VectorSpace(dim): dim = %v > %v = bits.UintSize", n, bits.UintSize))
	}

	ones := uint(1)
	for i := ones; i < n; i++ {
		ones <<= 1
		ones += 1
	}

	sp := GF2VectorSpace{n, ones}
	return &sp
}

// String returns a string representing
func (sp *GF2VectorSpace) String() string {
	return fmt.Sprintf("GF(2)sp{%v: %v}", sp.dim, sp.ones)
}

// NewGF2VectorSubspace create a sub vector space of sp with base bits b as sub space
// Panic if b is not element of sp, b needs more bits in the space as dim.
func (sp *GF2VectorSpace) NewGF2VectorSubspace(b uint) *GF2VectorSubspace {
	zero := uint(0)
	if b == zero {
		sp = &GF2VectorSpace{zero, zero}
		return &GF2VectorSubspace{sp, sp}
	}
	l := uint(bits.Len(b))
	if l > sp.dim {
		panic(fmt.Sprintf("NewGF2VectorSubspace(dim): vector %v not in space with dim = %v", b, sp.dim))
	}
	dim := uint(bits.OnesCount(b))
	subsp := &GF2VectorSpace{dim, b} // GF2VectorSpace is embedded
	return &GF2VectorSubspace{subsp, sp}
}

// NewGF2Vector create a vector with value in vector space,
// value must be greater equal 0.
func (s *GF2VectorSpace) NewGF2Vector(value uint) *GF2Vector {
	zero := uint(0)
	if s.dim == zero {
		return &GF2Vector{s, zero}
	}
	vmx := (uint(1) << s.dim) - 1
	if value > vmx {
		panic(fmt.Sprintf("NewGF2Vector(value): value = %v > %v", value, vmx))
	}

	return &GF2Vector{s, value}
}

// NewGF2VectorSet return a set of vectors with the values given.
// TODO add test // zero subspace
func (s *GF2VectorSpace) NewGF2VectorSet(u []uint) GF2VectorSet {
	if len(u) == 0 {
		// return empty set
		return GF2VectorSet{}
	}
	zero := uint(0)
	if s.dim == zero {
		if len(u) == 1 && u[0] == zero {
			// we have the zero subspace
			return GF2VectorSet{&GF2Vector{s, zero}}
		} else {
			// Only the zero space has dim zero
			panic(fmt.Sprintf("%v is not element of the vector space %v.", u, s))
		}
	}
	set := make([]*GF2Vector, len(u))
	for i, x := range u {
		set[i] = s.NewGF2Vector(x)
	}
	return set
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

// GF2VectorSpaceBase return the set of base vectors of s.
func (s *GF2VectorSpace) GF2VectorSpaceBase() GF2VectorSet {
	set := make(GF2VectorSet, 0, s.dim)
	o := s.ones
	for o != 0 {
		b := bits.TrailingZeros(o)
		v := uint(1) << b
		set = append(set, s.NewGF2Vector(v))
		o -= v
	}
	return set
}

/////////////////////////////////////// GF2VectorSubspace /////////////////////

// String returns a string representing
func (sp *GF2VectorSubspace) String() string {
	return fmt.Sprintf("GF(2)ssp{%v: %v, %v}", sp.dim, sp.ones, sp.container)
}

// NewGF2Vector not implemented, panic if accidently called
func (sp *GF2VectorSubspace) NewGF2Vector(value uint) *GF2Vector {
	panic("not implemented")
	return nil
}

// GF2BaseVector not implemented, panic if accidently called
func (sp *GF2VectorSubspace) GF2BaseVector(i uint) *GF2Vector {
	panic("not implemented")
	return nil
}

// GF2VectroSpaceBase return the set of base vectors of s.
func (sb *GF2VectorSubspace) GF2VectorSpaceBase() GF2VectorSet {
	set := make(GF2VectorSet, 0, sb.dim)
	o := sb.ones
	for o != 0 {
		b := bits.TrailingZeros(o)
		v := uint(1) << b
		sp := sb.container
		vv := sp.NewGF2Vector(v)
		set = append(set, vv)
		o -= v
	}
	return set
}

/////////////////////////////////////// GF2Vector /////////////////////////////

// String returns a string representing
func (v *GF2Vector) String() string {
	return fmt.Sprintf("%0[1]*[2]b", v.sp.dim, v.val)
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

// Val return the value as int.
// If no valid v is given we return 0.
func (v *GF2Vector) Val() uint {
	if v == nil {
		return 0
	}
	return v.val
}

// IsZeros return true if all bits are unset.
func (v *GF2Vector) IsZeros() bool {
	return v.val == 0
}

// IsOnes return true if all bits are set.
func (v *GF2Vector) IsOnes() bool {
	return v.val == v.sp.ones
}

/* TODO
Cmp compares x and y and returns:

-1 if x < y;
0 if x == y;
+1 if x > y.
*/

// BaseVector return the base vector representing the base with index i in the same vector space.
func (v *GF2Vector) GF2BaseVector(i uint) *GF2Vector {
	return v.sp.GF2BaseVector(i)
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

// UnitVectors return the set of unit vectors of the given bit vector
func (v *GF2Vector) UnitVectors() GF2VectorSet {
	vec := v.Val()
	res := make([]*GF2Vector, 0, bits.Len(vec))
	for vec != 0 {
		b := bits.TrailingZeros(vec)
		vb := uint(1) << b
		res = append(res, v.sp.NewGF2Vector(vb))
		vec -= vb
	}
	return res
}

/////////////////////////////////////// GF2VectorSet //////////////////////////

// Subspace return the the subspace spanned by set.
func (set GF2VectorSet) Subspace() (ssp *GF2VectorSubspace) {
	if len(set) == 0 {
		sp := NewGF2VectorSpace(uint(0))
		return sp.NewGF2VectorSubspace(uint(0))
	}
	ones := Or(set...)
	return set[0].sp.NewGF2VectorSubspace(ones.val)
}

// Subspace return first subspace of set, with dim.
// Check all combinations of subsets of a set of bit vectors
// to see wether a subset is confined to a subspace.
// If found equals false sp is nil.
// dim is int and must be greater 0, to support combinations.
func (set GF2VectorSet) HasSubspaceWithDim(dim int) (sp *GF2VectorSubspace, found bool) {
	if dim < 0 {
		panic(fmt.Sprintf("dim = %v < 0, no Subspace defined.", dim))
	}

	// compute all combinations of subset of size dim
	subsets := combinations(set, dim)

	for _, s := range subsets {
		set := GF2VectorSet(s)
		subSp := set.Subspace()
		if subSp.dim == uint(dim) {
			return subSp, true
		}
	}

	return
}

// ContainsBitOfVector return true if any element of the set has a bit of value set
func (set GF2VectorSet) ContainsBitOfVector(v *GF2Vector) bool {
	if v.IsZeros() {
		return false
	}
	ub := v.UnitVectors()
	for _, d := range ub {
		for _, s := range set {
			cmpVal := d.val & s.val
			if cmpVal != 0 && cmpVal == d.val {
				return true
			}
		}
	}
	return false
}

// ClearBits remove the bits of v from the elements of set
func (set GF2VectorSet) ClearBits(v *GF2Vector) {
	notv := Not(v)
	for i, s := range set {
		set[i] = And(s, notv)
	}
}

/////////////////////////////////////// Functions /////////////////////////////

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

// ComplementAnd return z = Not(And(x) = ^(x_1 | x_2 | ...),
// This can be used to as complement mask
// TODO add test
func ComplementAnd(x ...*GF2Vector) *GF2Vector {
	z := And(x...)
	return Not(z)
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
