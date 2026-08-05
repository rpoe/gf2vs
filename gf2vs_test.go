// Copyright 2024 Ralf Poeppel. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gf2vs

import (
	"fmt"
	"math/bits"
	"reflect"
	"testing"
)

func TestGF2VectorSpace(t *testing.T) {
	v := GF2VectorSpace{3, 7}
	want := "{3 7}"
	got := fmt.Sprintf("%v", v)
	if got != want {
		t.Errorf("GF2VectorSpace{3, 7, [1, 2, 4]} =\n%v, want\n%v", got, want)
	}
}

// TestNewGF2VectorSpaceString test NewGF2VectorSpace and String
func TestNewGF2VectorSpaceString(t *testing.T) {
	cases := []struct {
		in   uint
		want string
	}{
		{bits.UintSize + 1, "NewGF2VectorSpace(dim): dim = 65 > 64 = bits.UintSize"},
	}
	for _, c := range cases {
		func(in uint, want string) {
			defer func(in uint, want string) {
				r := recover()
				if r == nil {
					t.Errorf("NewGF2VectorSpace(%v) did not panic with %v",
						in, want)
				} else {
					if r != want {
						t.Errorf("NewGF2VectorSpace(%v) == Panic(%v),"+
							" want Panic(%v)", in, r, want)
					}
				}
			}(in, want)
			NewGF2VectorSpace(c.in)
		}(c.in, c.want)
	}

	cases = []struct {
		in   uint
		want string
	}{
		{0, "GF(2)sp{0: 0}"},
		{1, "GF(2)sp{1: 1}"},
		{2, "GF(2)sp{2: 3}"},
		{3, "GF(2)sp{3: 7}"},
		{4, "GF(2)sp{4: 15}"},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.in)
		got := fmt.Sprint(sp)
		if got != c.want {
			t.Errorf("NewGF2VectorSpace(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}

// TestNewGF2VectorSubspaceString test NewGF2VectorSubspace and String
func TestNewGF2VectorSubspaceString(t *testing.T) {
	cases := []struct {
		n    uint
		b    uint
		want string
	}{
		// panic from NewGF2VectorSpace
		{bits.UintSize + 1, 2, "NewGF2VectorSpace(dim): dim = 65 > 64 = bits.UintSize"},
		// panic from NewGF2VectorSubspace
		{1, 2, "NewGF2VectorSubspace(dim): vector 2 not in space with dim = 1"},
		{3, 8, "NewGF2VectorSubspace(dim): vector 8 not in space with dim = 3"},
	}
	for _, c := range cases {
		func(in, b uint, want string) {
			defer func(in uint, want string) {
				r := recover()
				if r == nil {
					t.Errorf("NewGF2VectorSubspace(%v) did not panic with %v",
						in, want)
				} else {
					if r != want {
						t.Errorf("NewGF2VectorSubspace(%v) == Panic(%v),"+
							" want Panic(%v)", in, r, want)
					}
				}
			}(in, want)
			sp := NewGF2VectorSpace(c.n)
			sp.NewGF2VectorSubspace(c.b)
		}(c.n, c.b, c.want)
	}

	cases = []struct {
		n    uint
		b    uint
		want string
	}{
		{0, 0, "GF(2)ssp{0: 0, GF(2)sp{0: 0}}"},
		{1, 0, "GF(2)ssp{0: 0, GF(2)sp{0: 0}}"}, // zero vector space
		{2, 1, "GF(2)ssp{1: 1, GF(2)sp{2: 3}}"},
		{3, 2, "GF(2)ssp{1: 2, GF(2)sp{3: 7}}"},
		{4, 3, "GF(2)ssp{2: 3, GF(2)sp{4: 15}}"},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.n)
		ssp := sp.NewGF2VectorSubspace(c.b)
		got := fmt.Sprint(ssp)
		if got != c.want {
			t.Errorf("NewGF2VectorSubspace(%v) = %v, want %v", c.n, got, c.want)
		}
	}
}

// TestNewGF2VectorString test NewGF2Vector and String
func TestNewGF2VectorString(t *testing.T) {
	cases := []struct {
		spin uint
		vin  uint
		want string
	}{
		//{1, -1, "NewGF2Vector(value): value = -1 < 0"}, invalid, test, vin is not uint
		{1, 2, "NewGF2Vector(value): value = 2 > 1"},
		{2, 4, "NewGF2Vector(value): value = 4 > 3"},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.spin)
		func(spin, vin uint, want string) {
			defer func(spin, vin uint, want string) {
				r := recover()
				if r == nil {
					t.Errorf("%v.NewGF2Vector(%v) did not panic with %v",
						spin, vin, want)
				} else {
					if r != want {
						t.Errorf("%v.NewGF2Vector(%v) == Panic(%v),"+
							" want Panic(%v)", spin, vin, r, want)
					}
				}
			}(spin, vin, want)
			sp.NewGF2Vector(vin)
		}(c.spin, c.vin, c.want)
	}

	cases = []struct {
		spin uint
		vin  uint
		want string
	}{
		{0, 0, "0"},
		{0, 1, "0"},
		{1, 0, "0"},
		{1, 1, "1"},
		{2, 0, "00"},
		{2, 1, "01"},
		{2, 2, "10"},
		{2, 3, "11"},
		{3, 0, "000"},
		{3, 1, "001"},
		{3, 2, "010"},
		{3, 3, "011"},
		{3, 4, "100"},
		{3, 5, "101"},
		{3, 6, "110"},
		{3, 7, "111"},
		{4, 0, "0000"},
		{4, 1, "0001"},
		{4, 2, "0010"},
		{4, 3, "0011"},
		{4, 4, "0100"},
		{4, 5, "0101"},
		{4, 6, "0110"},
		{4, 7, "0111"},
		{4, 8, "1000"},
		{4, 9, "1001"},
		{4, 10, "1010"},
		{4, 11, "1011"},
		{4, 12, "1100"},
		{4, 13, "1101"},
		{4, 14, "1110"},
		{4, 15, "1111"},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.spin)
		v := sp.NewGF2Vector(c.vin)
		got := fmt.Sprint(v)
		if got != c.want {
			t.Errorf("%v.NewGF2Vector(%v) = %v, want %v", c.spin, c.vin, got, c.want)
		}
	}
}

func TestNewGF2VectorSet(t *testing.T) {
	cases := []struct {
		n    uint
		u    []uint
		want string
	}{
		// Panic from inside NewGF2Vector, for vector not in vector space
		{1, []uint{2}, "NewGF2Vector(value): value = 2 > 1"},
	}
	for _, c := range cases {
		func(n uint, u []uint, want string) {
			sp := NewGF2VectorSpace(n)
			defer func(n uint, u []uint, want string) {
				r := recover()
				if r == nil {
					t.Errorf("%v.NewGF2VectorSet(%v) did not panic with %v",
						n, u, want)
				} else {
					if r != want {
						t.Errorf("%v.NewGF2VectorSet(%v) == Panic(%v),"+
							" want Panic(%v)", n, u, r, want)
					}
				}
			}(n, u, want)
			sp.NewGF2VectorSet(u)
		}(c.n, c.u, c.want)
	}

	cases = []struct {
		n    uint
		u    []uint
		want string
	}{
		{1, []uint{}, "[]"},
		{1, []uint{0}, "[0]"},
		{1, []uint{1}, "[1]"},
		{1, []uint{0, 1}, "[0 1]"},
		{3, []uint{0, 1, 2, 3}, "[000 001 010 011]"},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.n)
		set := sp.NewGF2VectorSet(c.u)
		str := fmt.Sprint(set)
		if str != c.want {
			t.Errorf("%v.NewGf2VectorSet(%v) = \n%v, want\n%v",
				sp, c.u, str, c.want)
		}
	}
}

func TestGF2Zeros(t *testing.T) {
	cases := []struct {
		dim  uint
		want string
	}{
		{1, "0"},
		{2, "00"},
		{3, "000"},
		{4, "0000"},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		v := sp.GF2Zeros()
		got := fmt.Sprint(v)
		if got != c.want {
			t.Errorf("%v.GF2Zeros() = %v, want %v", c.dim, got, c.want)
		}
	}
}

func TestGF2Ones(t *testing.T) {
	cases := []struct {
		dim  uint
		want string
	}{
		{1, "1"},
		{2, "11"},
		{3, "111"},
		{4, "1111"},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		v := sp.GF2Ones()
		got := fmt.Sprint(v)
		if got != c.want {
			t.Errorf("%v.GF2Ones() = %v, want %v", c.dim, got, c.want)
		}
	}
}

// One combined test for the function of GF2VectorSpace and GF2Vector
func TestGF2BaseVector(t *testing.T) {
	cases := []struct {
		dim   uint
		index uint
		want  string
	}{
		{1, 0, "GF2BaseVector(i): i = 0 out of range [1, 1]"},
		{1, 2, "GF2BaseVector(i): i = 2 out of range [1, 1]"},
		{2, 3, "GF2BaseVector(i): i = 3 out of range [1, 2]"},
	}
	for _, c := range cases {
		func(dim, index uint, want string) {
			defer func(dim, index uint, want string) {
				r := recover()
				if r == nil {
					t.Errorf("%v.GF2BaseVector(%v) did not panic with %v",
						dim, index, want)
				} else {
					if r != want {
						t.Errorf("%v.GF2BaseVector(%v) == Panic(%v),"+
							" want Panic(%v)", dim, index, r, want)
					}
				}
			}(dim, index, want)
			sp := NewGF2VectorSpace(dim)
			sp.GF2BaseVector(c.index)
		}(c.dim, c.index, c.want)
	}

	cases = []struct {
		dim   uint
		index uint
		want  string
	}{
		{1, 1, "1"},
		{2, 1, "01"},
		{2, 2, "10"},
		{3, 1, "001"},
		{3, 2, "010"},
		{3, 3, "100"},
		{4, 1, "0001"},
		{4, 2, "0010"},
		{4, 3, "0100"},
		{4, 4, "1000"},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		v := sp.GF2BaseVector(c.index)
		got := fmt.Sprint(v)
		if got != c.want {
			t.Errorf("%v.GF2BaseVector(%v) = %v, want %v", c.dim, c.index, got, c.want)
		}
		z := sp.GF2Zeros()
		v = z.GF2BaseVector(c.index)
		got = fmt.Sprint(v)
		if got != c.want {
			t.Errorf("%v.GF2BaseVector(%v) = %v, want %v", c.dim, c.index, got, c.want)
		}
	}
}

func TestGF2VectorSpaceBase(t *testing.T) {
	cases := []struct {
		dim     uint
		ones    uint
		baseStr string
	}{
		{0, 0, "[]"},
		{1, 1, "[1]"},
		{2, 3, "[01 10]"},
		{3, 7, "[001 010 100]"},
	}
	for _, c := range cases {
		sp := GF2VectorSpace{c.dim, c.ones}
		base := sp.GF2VectorSpaceBase()
		baseStr := fmt.Sprint(base)
		if baseStr != c.baseStr {
			t.Errorf("%v.GF2VectorSpaceBase() = \n%v, want\n%v", sp, baseStr, c.baseStr)
		}
	}
}

func TestGF2VectorSubspaceBase(t *testing.T) {
	cases := []struct {
		dim     uint
		ones    uint
		subones uint
		baseStr string
	}{
		{0, 0, 0, "[]"},
		{1, 1, 0, "[]"},
		{1, 1, 1, "[1]"},
		{2, 3, 0, "[]"},
		{2, 3, 1, "[01]"},
		{2, 3, 2, "[10]"},
		{3, 7, 1, "[001]"},
		{3, 7, 2, "[010]"},
		{3, 7, 4, "[100]"},
		{3, 7, 3, "[001 010]"},
		{3, 7, 5, "[001 100]"},
		{3, 7, 6, "[010 100]"},
		{3, 7, 7, "[001 010 100]"},
	}
	for _, c := range cases {
		sp := GF2VectorSpace{c.dim, c.ones}
		subsp := sp.NewGF2VectorSubspace(c.subones)
		base := subsp.GF2VectorSpaceBase()
		baseStr := fmt.Sprint(base)
		if baseStr != c.baseStr {
			t.Errorf("%v.GF2VectorSubspaceBase() = \n%v, want\n%v", sp, baseStr, c.baseStr)
		}
	}
}

func TestCopy(t *testing.T) {
	cases := []struct {
		dim  uint
		val  uint
		want uint
	}{
		{1, 0, 0},
		{1, 1, 1},
		{2, 0, 0},
		{2, 1, 1},
		{2, 2, 2},
		{2, 3, 3},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		v := sp.NewGF2Vector(c.val)
		vs1 := fmt.Sprint(v)
		w := sp.NewGF2Vector(c.want)
		ws := fmt.Sprint(w)
		cp := v.Copy()
		cs1 := fmt.Sprint(cp)
		if cs1 != ws {
			t.Errorf("%v.Copy() = %v, want %v", vs1, cp, ws)
		}
		// make sure change of source is not affecting copy
		v.val++
		cs2 := fmt.Sprint(cp)
		if cs2 != cs1 {
			t.Errorf("%v.Copy() source changed to %v", cs1, cs2)
		}
		// make sure change of copy is not affecting source
		v = sp.NewGF2Vector(c.val)
		vs1 = fmt.Sprint(v)
		cp = v.Copy()
		cp.val++
		vs2 := fmt.Sprint(v)
		if vs2 != vs1 {
			t.Errorf("%v.Copy() copy changed to %v", vs1, vs2)
		}
	}
}

func TestGF2VectorVal(t *testing.T) {
	s := NewGF2VectorSpace(3)
	v := GF2Vector{s, 2}
	want := "010"
	wVal2 := uint(2)
	wVal0 := uint(0)
	got := fmt.Sprintf("%v", &v)
	if got != want {
		t.Errorf("GF2Vector{3, 2} = %v, want %v", got, want)
	}
	gVal := v.Val()
	if gVal != wVal2 {
		t.Errorf("GF2Vector{3, 2}.Val() = %v, want %v", gVal, wVal2)
	}
	// special handling of uninitialized vectors
	var vp *GF2Vector
	gVal = vp.Val()
	if gVal != wVal0 {
		t.Errorf("nil.Val() = %v, want %v", gVal, wVal0)
	}
}

func TestIsZeros(t *testing.T) {
	cases := []struct {
		dim  uint
		val  uint
		want bool
	}{
		{3, 0, true},
		{3, 1, false},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		v := sp.NewGF2Vector(c.val)
		g := v.IsZeros()
		if g != c.want {
			t.Errorf("%v.IsZeros() = %v, want %v", v, g, c.want)
		}
	}
}

func TestIsOnes(t *testing.T) {
	cases := []struct {
		dim  uint
		val  uint
		want bool
	}{
		{3, 7, true},
		{3, 1, false},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		v := sp.NewGF2Vector(c.val)
		g := v.IsOnes()
		if g != c.want {
			t.Errorf("%v.IsOnes() = %v, want %v", v, g, c.want)
		}
	}
}

func TestIndexIsBaseVector(t *testing.T) {
	cases := []struct {
		dim     uint
		val     uint
		wIndex  uint
		wIsBase bool
	}{
		{1, 0, 0, false},
		{1, 1, 1, true},
		{2, 0, 0, false},
		{2, 1, 1, true},
		{2, 2, 2, true},
		{2, 3, 0, false},
		{3, 0, 0, false},
		{3, 1, 1, true},
		{3, 2, 2, true},
		{3, 3, 0, false},
		{3, 4, 3, true},
		{3, 5, 0, false},
		{3, 6, 0, false},
		{3, 7, 0, false},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		v := sp.NewGF2Vector(c.val)
		vs := fmt.Sprint(v)

		gIndex, gIsBase := v.Index()
		if gIndex != c.wIndex || gIsBase != c.wIsBase {
			t.Errorf("%v.Index() = %v, %v, want %v, %v",
				vs, gIndex, gIsBase, c.wIndex, c.wIsBase)
		}

		got := v.IsBaseVector()
		if got != c.wIsBase {
			t.Errorf("%v.IsBaseVector() = %v, want %v", vs, got, c.wIsBase)
		}
	}
}

func TestUnitVectors(t *testing.T) {
	cases := []struct {
		n      uint
		val    uint
		setStr string
	}{
		{3, 0, "[]"},
		{3, 1, "[001]"},
		{3, 2, "[010]"},
		{3, 4, "[100]"},
		{3, 3, "[001 010]"},
		{3, 5, "[001 100]"},
		{3, 6, "[010 100]"},
		{3, 7, "[001 010 100]"},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.n)
		v := sp.NewGF2Vector(c.val)
		set := v.UnitVectors()
		setStr := fmt.Sprint(set)
		if setStr != c.setStr {
			t.Errorf("%v.UnitVectors() = \n%v, want \n%v",
				v, setStr, c.setStr)
		}
	}
}

/////////////////////////////////////// GF2VectorSet //////////////////////////

func TestSubspace(t *testing.T) {
	cases := []struct {
		n    uint
		set  []uint
		ones uint
	}{
		{3, []uint{}, 0b0},
		{3, []uint{0b001, 0b101, 0b100}, 0b101},
		{3, []uint{0b001, 0b011, 0b010}, 0b011},
		{3, []uint{0b001, 0b011, 0b110}, 0b111},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.n)
		set := sp.NewGF2VectorSet(c.set)
		subSp := set.Subspace()
		ones := subSp.ones
		if ones != c.ones {
			t.Errorf("%v.Subspace(%v) = \n%0[3]*[4]b, want \n%0[3]*[5]b",
				sp, set, c.n, ones, c.ones)
		}
	}
}

func TestHasSubspaceWithDim(t *testing.T) {
	cases := []struct {
		n     uint
		set   []uint
		dim   int
		ones  uint
		found bool
	}{
		{3, []uint{0b001, 0b101, 0b100}, 0, 0b000, true},
		{3, []uint{0b001, 0b101, 0b100}, 2, 0b101, true},
		{3, []uint{0b001, 0b011, 0b010}, 1, 0b001, true},
		{3, []uint{0b001, 0b011, 0b010}, 2, 0b011, true},
		{3, []uint{0b001, 0b011, 0b010}, 3, 0, false},
		{3, []uint{0b001, 0b011, 0b110}, 3, 0b111, true},
		{4, []uint{0b0001, 0b0011, 0b0110}, 1, 0b0001, true},
		{4, []uint{0b1001, 0b0011, 0b0110}, 1, 0, false},
		{4, []uint{0b0001, 0b0011, 0b0110}, 3, 0b0111, true},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.n)
		csubSp := sp.NewGF2VectorSubspace(c.ones)
		set := sp.NewGF2VectorSet(c.set)
		subSp, found := set.HasSubspaceWithDim(c.dim)
		if c.found {
			if !found || !reflect.DeepEqual(subSp, csubSp) {
				t.Errorf("%v.HasSubspaceWithDim(%v) = \n%v, %v, want \n%v, %v",
					set, c.dim, subSp, found, csubSp, c.found)
			}
		} else {
			if found || subSp != nil {
				t.Errorf("%v.HasSubspaceWithDim(%v) = \n%v, %v, want \n%v, %v",
					set, c.dim, subSp, found, csubSp, c.found)
			}
		}
	}
}

func TestContainsBitOfVector(t *testing.T) {
	cases := []struct {
		n        uint
		set      []uint
		val      uint
		contains bool
	}{
		{0, []uint{}, 0, false},
		{1, []uint{1}, 0, false},
		{1, []uint{1}, 1, true},
		{1, []uint{1, 0}, 1, true},
		{2, []uint{1, 0}, 1, true},
		{2, []uint{1, 2, 3}, 1, true},
		{2, []uint{1, 2, 3}, 2, true},
		{2, []uint{1, 2, 3}, 3, true},
		{2, []uint{1, 0, 3}, 2, true},
		{3, []uint{1, 2, 3}, 4, false},
		{3, []uint{1, 2, 3}, 2, true},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.n)
		set := sp.NewGF2VectorSet(c.set)
		v := sp.NewGF2Vector(c.val)
		contains := set.ContainsBitOfVector(v)
		if contains != c.contains {
			t.Errorf("%v.ContainsBitOfVector(%v) = %v, wamt %v",
				set, v, contains, c.contains)
		}
	}
}

// RemoveBitsOfVector remove the bits of v from the elements of set
func TestSetClearBits(t *testing.T) {
	cases := []struct {
		n   uint
		set []uint
		val uint
		str string
	}{
		{0, []uint{}, 0, "[]"},
		{0, []uint{}, 1, "[]"},
		// illegal test case, will panic {0, []uint{1}, 0, "[1]"},
		{1, []uint{0}, 0, "[0]"},
		{1, []uint{0}, 1, "[0]"},
		{1, []uint{1}, 1, "[0]"},
		{1, []uint{1}, 0, "[1]"},
		{3, []uint{1, 2, 3}, 0, "[001 010 011]"},
		{3, []uint{1, 2, 3}, 1, "[000 010 010]"},
		{3, []uint{1, 2, 3}, 2, "[001 000 001]"},
		{3, []uint{1, 2, 3}, 3, "[000 000 000]"},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.n)
		set := sp.NewGF2VectorSet(c.set)
		str := fmt.Sprint(set)
		v := sp.NewGF2Vector(c.val)
		set.ClearBits(v)
		rmvstr := fmt.Sprint(set)
		if rmvstr != c.str {
			t.Errorf("%v.ClearBits(%v) = \n%v, want\n%v", str, v, rmvstr, c.str)
		}
	}
}

/////////////////////////////////////// Functions /////////////////////////////

func TestNot(t *testing.T) {
	cases := []struct {
		dim  uint
		val  uint
		want string
	}{
		{1, 0, "1"},
		{1, 1, "0"},
		{2, 0, "11"},
		{2, 1, "10"},
		{2, 2, "01"},
		{2, 3, "00"},
		{3, 0, "111"},
		{3, 1, "110"},
		{3, 2, "101"},
		{3, 3, "100"},
		{3, 4, "011"},
		{3, 5, "010"},
		{3, 6, "001"},
		{3, 7, "000"},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		v := sp.NewGF2Vector(c.val)
		vs := fmt.Sprint(v)
		z := Not(v)
		got := fmt.Sprint(z)
		if got != c.want {
			t.Errorf("%v.Not() = %v, want %v", vs, got, c.want)
		}
	}
}

func TestAnd(t *testing.T) {
	casesp := []struct {
		dim  []uint
		val  uint
		want string
	}{
		{[]uint{1, 2}, 1, "And: incompatible vector spaces: z.dim = 1 != 2 = y.dim"},
		{[]uint{1, 1, 2}, 1, "And: incompatible vector spaces: z.dim = 1 != 2 = y.dim"},
		{[]uint{1, 2, 2}, 1, "And: incompatible vector spaces: z.dim = 1 != 2 = y.dim"},
	}
	for _, c := range casesp {
		x := make([]*GF2Vector, len(c.dim))
		for i, d := range c.dim {
			sp := NewGF2VectorSpace(d)
			x[i] = sp.NewGF2Vector(c.val)
		}
		func(x []*GF2Vector, want string) {
			defer func(x []*GF2Vector, want string) {
				r := recover()
				if r == nil {
					t.Errorf("And(%v) did not panic with %v",
						x, want)
				} else {
					if r != want {
						t.Errorf("And(%v) == Panic(%v),"+
							" want Panic(%v)",
							x, r, want)
					}
				}
			}(x, want)
			And(x...)
		}(x, c.want)
	}

	cases := []struct {
		dim  uint
		val  []uint
		want string
	}{
		{1, []uint{0}, "0"},
		{1, []uint{1}, "1"},
		{1, []uint{0, 0}, "0"},
		{1, []uint{0, 1}, "0"},
		{1, []uint{1, 0}, "0"},
		{1, []uint{1, 1}, "1"},
		{2, []uint{0, 0}, "00"},
		{2, []uint{0, 1}, "00"},
		{2, []uint{0, 2}, "00"},
		{2, []uint{0, 3}, "00"},
		{2, []uint{1, 0}, "00"},
		{2, []uint{1, 1}, "01"},
		{2, []uint{1, 2}, "00"},
		{2, []uint{1, 3}, "01"},
		{2, []uint{2, 0}, "00"},
		{2, []uint{2, 1}, "00"},
		{2, []uint{2, 2}, "10"},
		{2, []uint{2, 3}, "10"},
		{2, []uint{3, 0}, "00"},
		{2, []uint{3, 1}, "01"},
		{2, []uint{3, 2}, "10"},
		{2, []uint{3, 3}, "11"},
		{3, []uint{0, 0}, "000"},
		{3, []uint{0, 1}, "000"},
		{3, []uint{0, 2}, "000"},
		{3, []uint{0, 3}, "000"},
		{3, []uint{0, 4}, "000"},
		{3, []uint{0, 5}, "000"},
		{3, []uint{0, 6}, "000"},
		{3, []uint{0, 7}, "000"},
		{3, []uint{1, 0}, "000"},
		{3, []uint{1, 1}, "001"},
		{3, []uint{1, 2}, "000"},
		{3, []uint{1, 3}, "001"},
		{3, []uint{1, 4}, "000"},
		{3, []uint{1, 5}, "001"},
		{3, []uint{1, 7}, "001"},
		{3, []uint{2, 0}, "000"},
		{3, []uint{2, 1}, "000"},
		{3, []uint{2, 2}, "010"},
		{3, []uint{2, 3}, "010"},
		{3, []uint{2, 4}, "000"},
		{3, []uint{2, 5}, "000"},
		{3, []uint{2, 6}, "010"},
		{3, []uint{2, 7}, "010"},
		{3, []uint{3, 0}, "000"},
		{3, []uint{3, 1}, "001"},
		{3, []uint{3, 2}, "010"},
		{3, []uint{3, 3}, "011"},
		{3, []uint{3, 4}, "000"},
		{3, []uint{3, 5}, "001"},
		{3, []uint{3, 6}, "010"},
		{3, []uint{3, 7}, "011"},
		{3, []uint{4, 0}, "000"},
		{3, []uint{4, 1}, "000"},
		{3, []uint{4, 2}, "000"},
		{3, []uint{4, 3}, "000"},
		{3, []uint{4, 4}, "100"},
		{3, []uint{4, 5}, "100"},
		{3, []uint{4, 6}, "100"},
		{3, []uint{4, 7}, "100"},
		{3, []uint{5, 0}, "000"},
		{3, []uint{5, 1}, "001"},
		{3, []uint{5, 2}, "000"},
		{3, []uint{5, 3}, "001"},
		{3, []uint{5, 4}, "100"},
		{3, []uint{5, 5}, "101"},
		{3, []uint{5, 6}, "100"},
		{3, []uint{5, 7}, "101"},
		{3, []uint{6, 0}, "000"},
		{3, []uint{6, 1}, "000"},
		{3, []uint{6, 2}, "010"},
		{3, []uint{6, 3}, "010"},
		{3, []uint{6, 4}, "100"},
		{3, []uint{6, 5}, "100"},
		{3, []uint{6, 6}, "110"},
		{3, []uint{6, 7}, "110"},
		{3, []uint{7, 0}, "000"},
		{3, []uint{7, 1}, "001"},
		{3, []uint{7, 2}, "010"},
		{3, []uint{7, 3}, "011"},
		{3, []uint{7, 4}, "100"},
		{3, []uint{7, 5}, "101"},
		{3, []uint{7, 6}, "110"},
		{3, []uint{7, 7}, "111"},
		{1, []uint{0, 0, 1}, "0"},
		{1, []uint{1, 1, 1}, "1"},
		{2, []uint{2, 2, 3}, "10"},
		{3, []uint{2, 2, 6}, "010"},
		{3, []uint{7, 5, 4}, "100"},
	}

	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		x := make([]*GF2Vector, len(c.val))
		for i, v := range c.val {
			x[i] = sp.NewGF2Vector(v)
		}
		z := And(x...)
		got := fmt.Sprint(z)
		if got != c.want {
			t.Errorf("And(%v) = %v, \"%v\", want %v", x, z, got, c.want)
		}
	}
}

func TestOr(t *testing.T) {
	casesp := []struct {
		dim  []uint
		val  uint
		want string
	}{
		{[]uint{1, 2}, 1, "Or: incompatible vector spaces: z.dim = 1 != 2 = y.dim"},
		{[]uint{1, 1, 2}, 1, "Or: incompatible vector spaces: z.dim = 1 != 2 = y.dim"},
		{[]uint{1, 2, 2}, 1, "Or: incompatible vector spaces: z.dim = 1 != 2 = y.dim"},
	}
	for _, c := range casesp {
		x := make([]*GF2Vector, len(c.dim))
		for i, d := range c.dim {
			sp := NewGF2VectorSpace(d)
			x[i] = sp.NewGF2Vector(c.val)
		}
		func(x []*GF2Vector, want string) {
			defer func(x []*GF2Vector, want string) {
				r := recover()
				if r == nil {
					t.Errorf("Or(%v) did not panic with %v",
						x, want)
				} else {
					if r != want {
						t.Errorf("Or(%v) == Panic(%v),"+
							" want Panic(%v)",
							x, r, want)
					}
				}
			}(x, want)
			Or(x...)
		}(x, c.want)
	}

	cases := []struct {
		dim  uint
		val  []uint
		want string
	}{
		{1, []uint{0}, "0"},
		{1, []uint{1}, "1"},
		{1, []uint{0, 0}, "0"},
		{1, []uint{0, 1}, "1"},
		{1, []uint{1, 0}, "1"},
		{1, []uint{1, 1}, "1"},
		{2, []uint{0, 0}, "00"},
		{2, []uint{0, 1}, "01"},
		{2, []uint{0, 2}, "10"},
		{2, []uint{0, 3}, "11"},
		{2, []uint{1, 0}, "01"},
		{2, []uint{1, 1}, "01"},
		{2, []uint{1, 2}, "11"},
		{2, []uint{1, 3}, "11"},
		{2, []uint{2, 0}, "10"},
		{2, []uint{2, 1}, "11"},
		{2, []uint{2, 2}, "10"},
		{2, []uint{2, 3}, "11"},
		{2, []uint{3, 0}, "11"},
		{2, []uint{3, 1}, "11"},
		{2, []uint{3, 2}, "11"},
		{2, []uint{3, 3}, "11"},
		{3, []uint{0, 0}, "000"},
		{3, []uint{0, 1}, "001"},
		{3, []uint{0, 2}, "010"},
		{3, []uint{0, 3}, "011"},
		{3, []uint{0, 4}, "100"},
		{3, []uint{0, 5}, "101"},
		{3, []uint{0, 6}, "110"},
		{3, []uint{0, 7}, "111"},
		{3, []uint{1, 0}, "001"},
		{3, []uint{1, 1}, "001"},
		{3, []uint{1, 2}, "011"},
		{3, []uint{1, 3}, "011"},
		{3, []uint{1, 4}, "101"},
		{3, []uint{1, 5}, "101"},
		{3, []uint{1, 7}, "111"},
		{3, []uint{2, 0}, "010"},
		{3, []uint{2, 1}, "011"},
		{3, []uint{2, 2}, "010"},
		{3, []uint{2, 3}, "011"},
		{3, []uint{2, 4}, "110"},
		{3, []uint{2, 5}, "111"},
		{3, []uint{2, 6}, "110"},
		{3, []uint{2, 7}, "111"},
		{3, []uint{3, 0}, "011"},
		{3, []uint{3, 1}, "011"},
		{3, []uint{3, 2}, "011"},
		{3, []uint{3, 3}, "011"},
		{3, []uint{3, 4}, "111"},
		{3, []uint{3, 5}, "111"},
		{3, []uint{3, 6}, "111"},
		{3, []uint{3, 7}, "111"},
		{3, []uint{4, 0}, "100"},
		{3, []uint{4, 1}, "101"},
		{3, []uint{4, 2}, "110"},
		{3, []uint{4, 3}, "111"},
		{3, []uint{4, 4}, "100"},
		{3, []uint{4, 5}, "101"},
		{3, []uint{4, 6}, "110"},
		{3, []uint{4, 7}, "111"},
		{3, []uint{5, 0}, "101"},
		{3, []uint{5, 1}, "101"},
		{3, []uint{5, 2}, "111"},
		{3, []uint{5, 3}, "111"},
		{3, []uint{5, 4}, "101"},
		{3, []uint{5, 5}, "101"},
		{3, []uint{5, 6}, "111"},
		{3, []uint{5, 7}, "111"},
		{3, []uint{6, 0}, "110"},
		{3, []uint{6, 1}, "111"},
		{3, []uint{6, 2}, "110"},
		{3, []uint{6, 3}, "111"},
		{3, []uint{6, 4}, "110"},
		{3, []uint{6, 5}, "111"},
		{3, []uint{6, 6}, "110"},
		{3, []uint{6, 7}, "111"},
		{3, []uint{7, 0}, "111"},
		{3, []uint{7, 1}, "111"},
		{3, []uint{7, 2}, "111"},
		{3, []uint{7, 3}, "111"},
		{3, []uint{7, 4}, "111"},
		{3, []uint{7, 5}, "111"},
		{3, []uint{7, 6}, "111"},
		{3, []uint{7, 7}, "111"},
		{1, []uint{0, 0, 1}, "1"},
		{1, []uint{1, 1, 1}, "1"},
		{2, []uint{2, 2, 3}, "11"},
		{3, []uint{2, 2, 6}, "110"},
		{3, []uint{4, 5, 7}, "111"},
	}

	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		x := make([]*GF2Vector, len(c.val))
		for i, v := range c.val {
			x[i] = sp.NewGF2Vector(v)
		}
		z := Or(x...)
		got := fmt.Sprint(z)
		if got != c.want {
			t.Errorf("Or(%v) = %v, \"%v\", want %v", x, z, got, c.want)
		}
	}
}

func TestXor(t *testing.T) {
	casesp := []struct {
		dim  []uint
		val  uint
		want string
	}{
		{[]uint{1, 2}, 1, "Xor: incompatible vector spaces: z.dim = 1 != 2 = y.dim"},
		{[]uint{1, 1, 2}, 1, "Xor: incompatible vector spaces: z.dim = 1 != 2 = y.dim"},
		{[]uint{1, 2, 2}, 1, "Xor: incompatible vector spaces: z.dim = 1 != 2 = y.dim"},
	}
	for _, c := range casesp {
		x := make([]*GF2Vector, len(c.dim))
		for i, d := range c.dim {
			sp := NewGF2VectorSpace(d)
			x[i] = sp.NewGF2Vector(c.val)
		}
		func(x []*GF2Vector, want string) {
			defer func(x []*GF2Vector, want string) {
				r := recover()
				if r == nil {
					t.Errorf("Xor(%v) did not panic with %v",
						x, want)
				} else {
					if r != want {
						t.Errorf("Xor(%v) == Panic(%v),"+
							" want Panic(%v)",
							x, r, want)
					}
				}
			}(x, want)
			Xor(x...)
		}(x, c.want)
	}

	cases := []struct {
		dim  uint
		val  []uint
		want string
	}{
		{1, []uint{0}, "0"},
		{1, []uint{1}, "1"},
		{1, []uint{0, 0}, "0"},
		{1, []uint{0, 1}, "1"},
		{1, []uint{1, 0}, "1"},
		{1, []uint{1, 1}, "0"},
		{2, []uint{0, 0}, "00"},
		{2, []uint{0, 1}, "01"},
		{2, []uint{0, 2}, "10"},
		{2, []uint{0, 3}, "11"},
		{2, []uint{1, 0}, "01"},
		{2, []uint{1, 1}, "00"},
		{2, []uint{1, 2}, "11"},
		{2, []uint{1, 3}, "10"},
		{2, []uint{2, 0}, "10"},
		{2, []uint{2, 1}, "11"},
		{2, []uint{2, 2}, "00"},
		{2, []uint{2, 3}, "01"},
		{2, []uint{3, 0}, "11"},
		{2, []uint{3, 1}, "10"},
		{2, []uint{3, 2}, "01"},
		{2, []uint{3, 3}, "00"},
		{3, []uint{0, 0}, "000"},
		{3, []uint{0, 1}, "001"},
		{3, []uint{0, 2}, "010"},
		{3, []uint{0, 3}, "011"},
		{3, []uint{0, 4}, "100"},
		{3, []uint{0, 5}, "101"},
		{3, []uint{0, 6}, "110"},
		{3, []uint{0, 7}, "111"},
		{3, []uint{1, 0}, "001"},
		{3, []uint{1, 1}, "000"},
		{3, []uint{1, 2}, "011"},
		{3, []uint{1, 3}, "010"},
		{3, []uint{1, 4}, "101"},
		{3, []uint{1, 5}, "100"},
		{3, []uint{1, 7}, "110"},
		{3, []uint{2, 0}, "010"},
		{3, []uint{2, 1}, "011"},
		{3, []uint{2, 2}, "000"},
		{3, []uint{2, 3}, "001"},
		{3, []uint{2, 4}, "110"},
		{3, []uint{2, 5}, "111"},
		{3, []uint{2, 6}, "100"},
		{3, []uint{2, 7}, "101"},
		{3, []uint{3, 0}, "011"},
		{3, []uint{3, 1}, "010"},
		{3, []uint{3, 2}, "001"},
		{3, []uint{3, 3}, "000"},
		{3, []uint{3, 4}, "111"},
		{3, []uint{3, 5}, "110"},
		{3, []uint{3, 6}, "101"},
		{3, []uint{3, 7}, "100"},
		{3, []uint{4, 0}, "100"},
		{3, []uint{4, 1}, "101"},
		{3, []uint{4, 2}, "110"},
		{3, []uint{4, 3}, "111"},
		{3, []uint{4, 4}, "000"},
		{3, []uint{4, 5}, "001"},
		{3, []uint{4, 6}, "010"},
		{3, []uint{4, 7}, "011"},
		{3, []uint{5, 0}, "101"},
		{3, []uint{5, 1}, "100"},
		{3, []uint{5, 2}, "111"},
		{3, []uint{5, 3}, "110"},
		{3, []uint{5, 4}, "001"},
		{3, []uint{5, 5}, "000"},
		{3, []uint{5, 6}, "011"},
		{3, []uint{5, 7}, "010"},
		{3, []uint{6, 0}, "110"},
		{3, []uint{6, 1}, "111"},
		{3, []uint{6, 2}, "100"},
		{3, []uint{6, 3}, "101"},
		{3, []uint{6, 4}, "010"},
		{3, []uint{6, 5}, "011"},
		{3, []uint{6, 6}, "000"},
		{3, []uint{6, 7}, "001"},
		{3, []uint{7, 0}, "111"},
		{3, []uint{7, 1}, "110"},
		{3, []uint{7, 2}, "101"},
		{3, []uint{7, 3}, "100"},
		{3, []uint{7, 4}, "011"},
		{3, []uint{7, 5}, "010"},
		{3, []uint{7, 6}, "001"},
		{3, []uint{7, 7}, "000"},
		{1, []uint{0, 0, 1}, "1"},
		{1, []uint{1, 1, 1}, "1"},
		{2, []uint{2, 2, 3}, "11"},
		{3, []uint{2, 2, 6}, "110"},
		{3, []uint{4, 5, 7}, "110"},
	}

	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		x := make([]*GF2Vector, len(c.val))
		for i, v := range c.val {
			x[i] = sp.NewGF2Vector(v)
		}
		z := Xor(x...)
		got := fmt.Sprint(z)
		if got != c.want {
			t.Errorf("Xor(%v) = %v, \"%v\", want %v", x, z, got, c.want)
		}
	}
}

func TestComplementOr(t *testing.T) {
	casesp := []struct {
		dim  []uint
		val  uint
		want string
	}{
		{[]uint{1, 2}, 1, "Or: incompatible vector spaces: z.dim = 1 != 2 = y.dim"},
		{[]uint{1, 1, 2}, 1, "Or: incompatible vector spaces: z.dim = 1 != 2 = y.dim"},
		{[]uint{1, 2, 2}, 1, "Or: incompatible vector spaces: z.dim = 1 != 2 = y.dim"},
	}
	for _, c := range casesp {
		x := make([]*GF2Vector, len(c.dim))
		for i, d := range c.dim {
			sp := NewGF2VectorSpace(d)
			x[i] = sp.NewGF2Vector(c.val)
		}
		func(x []*GF2Vector, want string) {
			defer func(x []*GF2Vector, want string) {
				r := recover()
				if r == nil {
					t.Errorf("Or(%v) did not panic with %v",
						x, want)
				} else {
					if r != want {
						t.Errorf("Or(%v) == Panic(%v),"+
							" want Panic(%v)",
							x, r, want)
					}
				}
			}(x, want)
			ComplementOr(x...)
		}(x, c.want)
	}

	cases := []struct {
		dim  uint
		val  []uint
		want string
	}{
		{1, []uint{0}, "1"},
		{1, []uint{1}, "0"},
		{1, []uint{0, 0}, "1"},
		{1, []uint{0, 1}, "0"},
		{1, []uint{1, 0}, "0"},
		{1, []uint{1, 1}, "0"},
		{2, []uint{0, 0}, "11"},
		{2, []uint{0, 1}, "10"},
		{2, []uint{0, 2}, "01"},
		{2, []uint{0, 3}, "00"},
		{2, []uint{1, 0}, "10"},
		{2, []uint{1, 1}, "10"},
		{2, []uint{1, 2}, "00"},
		{2, []uint{1, 3}, "00"},
		{2, []uint{2, 0}, "01"},
		{2, []uint{2, 1}, "00"},
		{2, []uint{2, 2}, "01"},
		{2, []uint{2, 3}, "00"},
		{2, []uint{3, 0}, "00"},
		{2, []uint{3, 1}, "00"},
		{2, []uint{3, 2}, "00"},
		{2, []uint{3, 3}, "00"},
		{3, []uint{0, 0}, "111"},
		{3, []uint{0, 1}, "110"},
		{3, []uint{0, 2}, "101"},
		{3, []uint{0, 3}, "100"},
		{3, []uint{0, 4}, "011"},
		{3, []uint{0, 5}, "010"},
		{3, []uint{0, 6}, "001"},
		{3, []uint{0, 7}, "000"},
		{3, []uint{1, 0}, "110"},
		{3, []uint{1, 1}, "110"},
		{3, []uint{1, 2}, "100"},
		{3, []uint{1, 3}, "100"},
		{3, []uint{1, 4}, "010"},
		{3, []uint{1, 5}, "010"},
		{3, []uint{1, 7}, "000"},
		{3, []uint{2, 0}, "101"},
		{3, []uint{2, 1}, "100"},
		{3, []uint{2, 2}, "101"},
		{3, []uint{2, 3}, "100"},
		{3, []uint{2, 4}, "001"},
		{3, []uint{2, 5}, "000"},
		{3, []uint{2, 6}, "001"},
		{3, []uint{2, 7}, "000"},
		{3, []uint{3, 0}, "100"},
		{3, []uint{3, 1}, "100"},
		{3, []uint{3, 2}, "100"},
		{3, []uint{3, 3}, "100"},
		{3, []uint{3, 4}, "000"},
		{3, []uint{3, 5}, "000"},
		{3, []uint{3, 6}, "000"},
		{3, []uint{3, 7}, "000"},
		{3, []uint{4, 0}, "011"},
		{3, []uint{4, 1}, "010"},
		{3, []uint{4, 2}, "001"},
		{3, []uint{4, 3}, "000"},
		{3, []uint{4, 4}, "011"},
		{3, []uint{4, 5}, "010"},
		{3, []uint{4, 6}, "001"},
		{3, []uint{4, 7}, "000"},
		{3, []uint{5, 0}, "010"},
		{3, []uint{5, 1}, "010"},
		{3, []uint{5, 2}, "000"},
		{3, []uint{5, 3}, "000"},
		{3, []uint{5, 4}, "010"},
		{3, []uint{5, 5}, "010"},
		{3, []uint{5, 6}, "000"},
		{3, []uint{5, 7}, "000"},
		{3, []uint{6, 0}, "001"},
		{3, []uint{6, 1}, "000"},
		{3, []uint{6, 2}, "001"},
		{3, []uint{6, 3}, "000"},
		{3, []uint{6, 4}, "001"},
		{3, []uint{6, 5}, "000"},
		{3, []uint{6, 6}, "001"},
		{3, []uint{6, 7}, "000"},
		{3, []uint{7, 0}, "000"},
		{3, []uint{7, 1}, "000"},
		{3, []uint{7, 2}, "000"},
		{3, []uint{7, 3}, "000"},
		{3, []uint{7, 4}, "000"},
		{3, []uint{7, 5}, "000"},
		{3, []uint{7, 6}, "000"},
		{3, []uint{7, 7}, "000"},
		{1, []uint{0, 0, 1}, "0"},
		{1, []uint{1, 1, 1}, "0"},
		{2, []uint{2, 2, 3}, "00"},
		{3, []uint{2, 2, 6}, "001"},
		{3, []uint{4, 5, 7}, "000"},
		{5, []uint{8, 21, 1, 29}, "00010"},
	}

	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		x := make([]*GF2Vector, len(c.val))
		for i, v := range c.val {
			x[i] = sp.NewGF2Vector(v)
		}
		z := ComplementOr(x...)
		got := fmt.Sprint(z)
		if got != c.want {
			t.Errorf("ComplementOr(%v) = %v, \"%v\", want %v", x, z, got, c.want)
		}
	}
}

func TestComplementXor(t *testing.T) {
	casesp := []struct {
		dim  []uint
		val  uint
		want string
	}{
		{[]uint{1, 2}, 1, "Xor: incompatible vector spaces: z.dim = 1 != 2 = y.dim"},
		{[]uint{1, 1, 2}, 1, "Xor: incompatible vector spaces: z.dim = 1 != 2 = y.dim"},
		{[]uint{1, 2, 2}, 1, "Xor: incompatible vector spaces: z.dim = 1 != 2 = y.dim"},
	}
	for _, c := range casesp {
		x := make([]*GF2Vector, len(c.dim))
		for i, d := range c.dim {
			sp := NewGF2VectorSpace(d)
			x[i] = sp.NewGF2Vector(c.val)
		}
		func(x []*GF2Vector, want string) {
			defer func(x []*GF2Vector, want string) {
				r := recover()
				if r == nil {
					t.Errorf("Xor(%v) did not panic with %v",
						x, want)
				} else {
					if r != want {
						t.Errorf("Xor(%v) == Panic(%v),"+
							" want Panic(%v)",
							x, r, want)
					}
				}
			}(x, want)
			ComplementXor(x...)
		}(x, c.want)
	}

	cases := []struct {
		dim  uint
		val  []uint
		want string
	}{
		{1, []uint{0}, "1"},
		{1, []uint{1}, "0"},
		{1, []uint{0, 0}, "1"},
		{1, []uint{0, 1}, "0"},
		{1, []uint{1, 0}, "0"},
		{1, []uint{1, 1}, "1"},
		{2, []uint{0, 0}, "11"},
		{2, []uint{0, 1}, "10"},
		{2, []uint{0, 2}, "01"},
		{2, []uint{0, 3}, "00"},
		{2, []uint{1, 0}, "10"},
		{2, []uint{1, 1}, "11"},
		{2, []uint{1, 2}, "00"},
		{2, []uint{1, 3}, "01"},
		{2, []uint{2, 0}, "01"},
		{2, []uint{2, 1}, "00"},
		{2, []uint{2, 2}, "11"},
		{2, []uint{2, 3}, "10"},
		{2, []uint{3, 0}, "00"},
		{2, []uint{3, 1}, "01"},
		{2, []uint{3, 2}, "10"},
		{2, []uint{3, 3}, "11"},
		{3, []uint{0, 0}, "111"},
		{3, []uint{0, 1}, "110"},
		{3, []uint{0, 2}, "101"},
		{3, []uint{0, 3}, "100"},
		{3, []uint{0, 4}, "011"},
		{3, []uint{0, 5}, "010"},
		{3, []uint{0, 6}, "001"},
		{3, []uint{0, 7}, "000"},
		{3, []uint{1, 0}, "110"},
		{3, []uint{1, 1}, "111"},
		{3, []uint{1, 2}, "100"},
		{3, []uint{1, 3}, "101"},
		{3, []uint{1, 4}, "010"},
		{3, []uint{1, 5}, "011"},
		{3, []uint{1, 7}, "001"},
		{3, []uint{2, 0}, "101"},
		{3, []uint{2, 1}, "100"},
		{3, []uint{2, 2}, "111"},
		{3, []uint{2, 3}, "110"},
		{3, []uint{2, 4}, "001"},
		{3, []uint{2, 5}, "000"},
		{3, []uint{2, 6}, "011"},
		{3, []uint{2, 7}, "010"},
		{3, []uint{3, 0}, "100"},
		{3, []uint{3, 1}, "101"},
		{3, []uint{3, 2}, "110"},
		{3, []uint{3, 3}, "111"},
		{3, []uint{3, 4}, "000"},
		{3, []uint{3, 5}, "001"},
		{3, []uint{3, 6}, "010"},
		{3, []uint{3, 7}, "011"},
		{3, []uint{4, 0}, "011"},
		{3, []uint{4, 1}, "010"},
		{3, []uint{4, 2}, "001"},
		{3, []uint{4, 3}, "000"},
		{3, []uint{4, 4}, "111"},
		{3, []uint{4, 5}, "110"},
		{3, []uint{4, 6}, "101"},
		{3, []uint{4, 7}, "100"},
		{3, []uint{5, 0}, "010"},
		{3, []uint{5, 1}, "011"},
		{3, []uint{5, 2}, "000"},
		{3, []uint{5, 3}, "001"},
		{3, []uint{5, 4}, "110"},
		{3, []uint{5, 5}, "111"},
		{3, []uint{5, 6}, "100"},
		{3, []uint{5, 7}, "101"},
		{3, []uint{6, 0}, "001"},
		{3, []uint{6, 1}, "000"},
		{3, []uint{6, 2}, "011"},
		{3, []uint{6, 3}, "010"},
		{3, []uint{6, 4}, "101"},
		{3, []uint{6, 5}, "100"},
		{3, []uint{6, 6}, "111"},
		{3, []uint{6, 7}, "110"},
		{3, []uint{7, 0}, "000"},
		{3, []uint{7, 1}, "001"},
		{3, []uint{7, 2}, "010"},
		{3, []uint{7, 3}, "011"},
		{3, []uint{7, 4}, "100"},
		{3, []uint{7, 5}, "101"},
		{3, []uint{7, 6}, "110"},
		{3, []uint{7, 7}, "111"},
		{1, []uint{0, 0, 1}, "0"},
		{1, []uint{1, 1, 1}, "0"},
		{2, []uint{2, 2, 3}, "00"},
		{3, []uint{2, 2, 6}, "001"},
		{3, []uint{4, 5, 7}, "001"},
		{5, []uint{2, 8}, "10101"},
		{5, []uint{2}, "11101"},
	}

	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		x := make([]*GF2Vector, len(c.val))
		for i, v := range c.val {
			x[i] = sp.NewGF2Vector(v)
		}
		z := ComplementXor(x...)
		got := fmt.Sprint(z)
		if got != c.want {
			t.Errorf("ComplementXor(%v) = %v, \"%v\", want %v", x, z, got, c.want)
		}
	}
}

func TestMaskBits(t *testing.T) {
	cases := []struct {
		dim  uint
		x    uint
		m    uint
		want string
	}{
		{4, 0b1011, 0b1110, "1010"},
		{8, 0b01001011, 0b10001110, "00001010"},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		x := sp.NewGF2Vector(c.x)
		m := sp.NewGF2Vector(c.m)
		z := MaskBits(x, m)
		got := fmt.Sprint(z)
		if got != c.want {
			t.Errorf("MaskBits(%v, %v) = %v, \"%v\", want %v", x, m, z, got, c.want)
		}
	}
}

func TestClearBits(t *testing.T) {
	cases := []struct {
		dim  uint
		x    uint
		m    uint
		want string
	}{
		{4, 0b1011, 0b1110, "0001"},
		{8, 0b01001011, 0b10001110, "01000001"},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		x := sp.NewGF2Vector(c.x)
		m := sp.NewGF2Vector(c.m)
		z := ClearBits(x, m)
		got := fmt.Sprint(z)
		if got != c.want {
			t.Errorf("ClearBits(%v, %v) = %v, \"%v\", want %v", x, m, z, got, c.want)
		}
	}
}

func TestSetBits(t *testing.T) {
	cases := []struct {
		dim  uint
		x    uint
		m    uint
		want string
	}{
		{4, 0b1000, 0b1110, "1110"},
		{8, 0b01001011, 0b00000100, "01001111"},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		x := sp.NewGF2Vector(c.x)
		m := sp.NewGF2Vector(c.m)
		z := SetBits(x, m)
		got := fmt.Sprint(z)
		if got != c.want {
			t.Errorf("SetBits(%v, %v) = %v, \"%v\", want %v", x, m, z, got, c.want)
		}
	}
}

func TestToggleBits(t *testing.T) {
	cases := []struct {
		dim  uint
		x    uint
		m    uint
		want string
	}{
		{4, 0b1000, 0b1110, "0110"},
		{8, 0b01001011, 0b00001100, "01000111"},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		x := sp.NewGF2Vector(c.x)
		m := sp.NewGF2Vector(c.m)
		z := ToggleBits(x, m)
		got := fmt.Sprint(z)
		if got != c.want {
			t.Errorf("ToggleBits(%v, %v) = %v, \"%v\", want %v", x, m, z, got, c.want)
		}
	}
}

func TestOnesCount(t *testing.T) {
	cases := []struct {
		dim  uint
		val  uint
		want int
	}{
		{1, 0, 0},
		{1, 1, 1},
		{2, 0, 0},
		{2, 1, 1},
		{2, 2, 1},
		{2, 3, 2},
		{3, 0, 0},
		{3, 1, 1},
		{3, 2, 1},
		{3, 3, 2},
		{3, 4, 1},
		{3, 5, 2},
		{3, 6, 2},
		{3, 7, 3},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		v := sp.NewGF2Vector(c.val)
		vs := fmt.Sprint(v)
		got := OnesCount(v)
		if got != c.want {
			t.Errorf("OnesCount(%v) = %v, want %v", vs, got, c.want)
		}
	}
}

func TestScalarProduct(t *testing.T) {
	cases := []struct {
		dim  uint
		a    uint
		b    uint
		want int
	}{
		{1, 0, 0, 0},
		{1, 0, 1, 0},
		{1, 1, 0, 0},
		{1, 1, 1, 1},
		{2, 0, 0, 0},
		{2, 0, 1, 0},
		{2, 0, 2, 0},
		{2, 0, 3, 0},
		{2, 1, 0, 0},
		{2, 1, 1, 1},
		{2, 1, 2, 0},
		{2, 1, 3, 1},
		{2, 2, 0, 0},
		{2, 2, 1, 0},
		{2, 2, 2, 1},
		{2, 2, 3, 1},
		{2, 3, 0, 0},
		{2, 3, 1, 1},
		{2, 3, 2, 1},
		{2, 3, 3, 2},
		{3, 0, 0, 0},
		{3, 0, 1, 0},
		{3, 0, 2, 0},
		{3, 0, 3, 0},
		{3, 0, 4, 0},
		{3, 0, 5, 0},
		{3, 0, 6, 0},
		{3, 0, 7, 0},
		{3, 1, 0, 0},
		{3, 1, 1, 1},
		{3, 1, 2, 0},
		{3, 1, 3, 1},
		{3, 1, 4, 0},
		{3, 1, 5, 1},
		{3, 1, 7, 1},
		{3, 2, 0, 0},
		{3, 2, 1, 0},
		{3, 2, 2, 1},
		{3, 2, 3, 1},
		{3, 2, 4, 0},
		{3, 2, 5, 0},
		{3, 2, 6, 1},
		{3, 2, 7, 1},
		{3, 3, 0, 0},
		{3, 3, 1, 1},
		{3, 3, 2, 1},
		{3, 3, 3, 2},
		{3, 3, 4, 0},
		{3, 3, 5, 1},
		{3, 3, 6, 1},
		{3, 3, 7, 2},
		{3, 4, 0, 0},
		{3, 4, 1, 0},
		{3, 4, 2, 0},
		{3, 4, 3, 0},
		{3, 4, 4, 1},
		{3, 4, 5, 1},
		{3, 4, 6, 1},
		{3, 4, 7, 1},
		{3, 5, 0, 0},
		{3, 5, 1, 1},
		{3, 5, 2, 0},
		{3, 5, 3, 1},
		{3, 5, 4, 1},
		{3, 5, 5, 2},
		{3, 5, 6, 1},
		{3, 5, 7, 2},
		{3, 6, 0, 0},
		{3, 6, 1, 0},
		{3, 6, 2, 1},
		{3, 6, 3, 1},
		{3, 6, 4, 1},
		{3, 6, 5, 1},
		{3, 6, 6, 2},
		{3, 6, 7, 2},
		{3, 7, 0, 0},
		{3, 7, 1, 1},
		{3, 7, 2, 1},
		{3, 7, 3, 2},
		{3, 7, 4, 1},
		{3, 7, 5, 2},
		{3, 7, 6, 2},
		{3, 7, 7, 3},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		a := sp.NewGF2Vector(c.a)
		b := sp.NewGF2Vector(c.b)
		got := ScalarProduct(a, b)
		if got != c.want {
			t.Errorf("Scalarproduct(%v, %v) = %v, want %v", a, b, got, c.want)
		}
	}
}

/*
func TestSpanOfSet(t *testing.T) {
	cases := []struct {
		dim   uint
		a     []uint
		span  uint
		spDim int
	}{
		// 0
		{1, []uint{0}, 0, 0},
		{1, []uint{1}, 1, 1},
		{2, []uint{2}, 2, 1},
		{3, []uint{4}, 4, 1},
		{3, []uint{7}, 7, 3},
		{1, []uint{0, 0}, 0, 0},
		{1, []uint{0, 1}, 1, 1},
		{1, []uint{1, 0}, 1, 1},
		{1, []uint{1, 1}, 1, 1},
		{2, []uint{0, 0}, 0, 0},
		// 10
		{2, []uint{0, 1}, 1, 1},
		{2, []uint{0, 2}, 2, 1},
		{2, []uint{0, 3}, 3, 2},
		{2, []uint{1, 0}, 1, 1},
		{2, []uint{1, 1}, 1, 1},
		{2, []uint{1, 2}, 3, 2},
		{2, []uint{1, 3}, 3, 2},
		{2, []uint{2, 0}, 2, 1},
		{2, []uint{2, 1}, 3, 2},
		{2, []uint{2, 2}, 2, 1},
		// 20
		{2, []uint{2, 3}, 3, 2},
		{2, []uint{3, 0}, 3, 2},
		{2, []uint{3, 1}, 3, 2},
		{2, []uint{3, 2}, 3, 2},
		{2, []uint{3, 3}, 3, 2},
		{3, []uint{0, 0}, 0, 0},
		{3, []uint{0, 1}, 1, 1},
		{3, []uint{0, 2}, 2, 1},
		{3, []uint{0, 3}, 3, 2},
		{3, []uint{0, 4}, 4, 1},
		// 30
		{3, []uint{0, 5}, 5, 2},
		{3, []uint{0, 6}, 6, 2},
		{3, []uint{0, 7}, 7, 3},
		{3, []uint{1, 0}, 1, 1},
		{3, []uint{1, 1}, 1, 1},
		{3, []uint{1, 2}, 3, 2},
		{3, []uint{1, 3}, 3, 2},
		{3, []uint{1, 4}, 5, 2},
		{3, []uint{1, 5}, 5, 2},
		{3, []uint{1, 7}, 7, 3},
		// 40
		{3, []uint{2, 0}, 2, 1},
		{3, []uint{2, 1}, 3, 2},
		{3, []uint{2, 2}, 2, 1},
		{3, []uint{2, 3}, 3, 2},
		{3, []uint{2, 4}, 6, 2},
		{3, []uint{2, 5}, 7, 3},
		{3, []uint{2, 6}, 6, 2},
		{3, []uint{2, 7}, 7, 3},
		{3, []uint{3, 0}, 3, 2},
		{3, []uint{3, 1}, 3, 2},
		// 50
		{3, []uint{3, 2}, 3, 2},
		{3, []uint{3, 3}, 3, 2},
		{3, []uint{3, 4}, 7, 3},
		{3, []uint{3, 5}, 7, 3},
		{3, []uint{3, 6}, 7, 3},
		{3, []uint{3, 7}, 7, 3},
		{3, []uint{4, 0}, 4, 1},
		{3, []uint{4, 1}, 5, 2},
		{3, []uint{4, 2}, 6, 2},
		{3, []uint{4, 3}, 7, 3},
		// 60
		{3, []uint{4, 4}, 4, 1},
		{3, []uint{4, 5}, 5, 2},
		{3, []uint{4, 6}, 6, 2},
		{3, []uint{4, 7}, 7, 3},
		{3, []uint{5, 0}, 5, 2},
		{3, []uint{5, 1}, 5, 2},
		{3, []uint{5, 2}, 7, 3},
		{3, []uint{5, 3}, 7, 3},
		{3, []uint{5, 4}, 5, 2},
		{3, []uint{5, 5}, 5, 2},
		// 70
		{3, []uint{5, 6}, 7, 3},
		{3, []uint{5, 7}, 7, 3},
		{3, []uint{6, 0}, 6, 2},
		{3, []uint{6, 1}, 7, 3},
		{3, []uint{6, 2}, 6, 2},
		{3, []uint{6, 3}, 7, 3},
		{3, []uint{6, 4}, 6, 2},
		{3, []uint{6, 5}, 7, 3},
		{3, []uint{6, 6}, 6, 2},
		{3, []uint{6, 7}, 7, 3},
		// 80
		{3, []uint{7, 0}, 7, 3},
		{3, []uint{7, 1}, 7, 3},
		{3, []uint{7, 2}, 7, 3},
		{3, []uint{7, 3}, 7, 3},
		{3, []uint{7, 4}, 7, 3},
		{3, []uint{7, 5}, 7, 3},
		{3, []uint{7, 6}, 7, 3},
		{3, []uint{7, 7}, 7, 3},
		{3, []uint{2, 7, 7}, 7, 3},
		// 90
		{3, []uint{1, 2, 3}, 3, 2},
		{3, []uint{2, 4, 6}, 6, 2},
		{3, []uint{1, 4, 5}, 5, 2},
	}
	for ic, c := range cases {
		//fmt.Println(i)
		sp := NewGF2VectorSpace(c.dim)
		s := make(GF2VectorSet, len(c.a))
		for i := range c.a {
			s[i] = sp.NewGF2Vector(c.a[i])
		}
		span, spDim := s.SpanOfSet()
		if span.Val() != c.span || spDim != c.spDim {
			t.Errorf("%v: %v.SpanOfSet() = %v, %v, want %b, %v",
				ic, s, span, spDim, c.span, c.spDim)
		}
	}
}

func TestBitDecomp(t *testing.T) {
	cases := []struct {
		dim  uint
		val  uint
		base []uint
	}{
		{1, 0, []uint{}},
		{1, 1, []uint{1}},
		{2, 0, []uint{}},
		{2, 1, []uint{1}},
		{2, 2, []uint{2}},
		{2, 3, []uint{1, 2}},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		vc := sp.NewGF2Vector(c.val)
		vBase := sp.NewGF2VectorSet(c.base)
		base := vc.BitDecomp()
		if !reflect.DeepEqual(base, vBase) {
			t.Errorf("%v.BitDecomp() = \n%v, want \n%v", vc, base, vBase)
		}
	}
}

func TestSubspaceOfSet(t *testing.T) {
	cases := []struct {
		dim     uint
		sbstDim int
		set     []uint
		val     uint
		found   bool
	}{
		{4, 1, []uint{2, 3, 7, 5}, 2, true},
		{4, 2, []uint{3, 3, 7, 5}, 3, true},
		{4, 2, []uint{3, 5, 9, 5}, 5, true},
		{4, 2, []uint{3, 5, 9, 6}, 0, false},
		{4, 2, []uint{7, 9, 7, 8}, 9, true},
		{4, 3, []uint{7, 9, 7, 8}, 0, false},
		{5, 3, []uint{7, 9, 7, 3, 8}, 7, true},
		{5, 3, []uint{5, 9, 7, 3, 8}, 13, true},
		{5, 3, []uint{16, 5, 7, 3, 8}, 7, true},
		{5, 3, []uint{9, 5, 7, 3, 8}, 13, true},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		set := sp.NewGF2VectorSet(c.set)
		v, found := set.Subspace(c.sbstDim)
		val := uint(0)
		if found {
			val = v.subOnes
		}
		if found != c.found || val != c.val {
			t.Errorf("%v.Subspace(%v, %v) =\n%v, %v, want\n%v, %v",
				c.set, set, c.sbstDim, val, found, c.val, c.found)
		}
	}
}
*/
/*
func TestReduceOtherInUnit(t *testing.T) {
	cases := []struct {
		val         int
		allDigits   int
		u           []int
		grid        []int
		gridReduced []int
	}{
		{3, 0b1111, []int{0, 1, 2, 3}, []int{3, 3, 7, 11}, []int{3, 3, 4, 8}},
		{5, 0b1111, []int{0, 1, 4, 5}, []int{3, 5, 7, 6, 9, 5}, []int{2, 5, 7, 6, 8, 5}},
		{1, 0b1111, []int{0, 1, 4, 5}, []int{4, 6, 7, 5, 8, 6}, []int{4, 6, 7, 5, 8, 6}},
		{7, 0b11111, []int{0, 1, 2, 3, 4}, []int{9, 5, 7, 3, 8}, []int{8, 5, 7, 3, 8}},
	}
	for _, c := range cases {
		gridWork := make([]int, len(c.grid))
		copy(gridWork, c.grid)
		reduceOtherInUnit(c.val, c.allDigits, c.u, gridWork)
		if !reflect.DeepEqual(gridWork, c.gridReduced) {
			t.Errorf("reduceOtherInUnit(%v, %0b, %v, %v) =\n%v, want\n%v",
				c.val, c.allDigits, c.u, c.grid, gridWork, c.gridReduced)
		}
	}
}

func TestUnitBits(t *testing.T) {
	cases := []struct {
		vec uint
		res []int
	}{
		{0, []int{}},
		{1, []int{1}},
		{2, []int{2}},
		{3, []int{1, 2}},
		{4, []int{4}},
		{5, []int{1, 4}},
		{6, []int{2, 4}},
		{7, []int{1, 2, 4}},
	}
	for _, c := range cases {
		res := unitBits(c.vec)
		if !reflect.DeepEqual(res, c.res) {
			t.Errorf("unitBits(%v) = \n%v, want\n%v", c.vec, res, c.res)
		}
	}
}

func TestIsValueInSet(t *testing.T) {
	cases := []struct {
		value int
		set   []int
		grid  []int
		isIn  bool
	}{
		{0, []int{0, 1}, []int{1, 2}, false},
		{1, []int{0, 1}, []int{1, 2}, true},
		{2, []int{0, 1}, []int{1, 2}, true},
		{3, []int{0, 1}, []int{1, 2}, true},
		{4, []int{0, 1}, []int{1, 2}, false},
		{4, []int{0, 1}, []int{1, 4}, true},
		{4, []int{0, 1}, []int{1, 5}, true},
		{4, []int{0, 1}, []int{1, 6}, true},
		{5, []int{0, 1}, []int{1, 2}, true},
	}
	for _, c := range cases[0:] {
		isIn := isValueInSet(c.value, c.set, c.grid)
		if isIn != c.isIn {
			t.Errorf("isValueInSet(%v, %v, %v) = %v, want %v",
				c.value, c.set, c.grid, isIn, c.isIn)
		}
	}
}

func TestRemoveValueFromSet(t *testing.T) {
	cases := []struct {
		d         int
		allDigits int
		set       []int
		grid      []int
		gridNew   []int
	}{
		{1, 3, []int{0, 1}, []int{1, 2}, []int{0, 2}},
		{2, 3, []int{0, 1}, []int{1, 2}, []int{1, 0}},
		{2, 7, []int{0, 1}, []int{3, 7}, []int{1, 5}},
		{2, 15, []int{0, 1}, []int{3, 7}, []int{1, 5}},
		{2, 15, []int{0, 1}, []int{6, 7}, []int{4, 5}},
		{2, 15, []int{0, 1}, []int{5, 7}, []int{5, 5}},
		{2, 15, []int{0, 1}, []int{13, 7}, []int{13, 5}},
		{2, 15, []int{0, 1}, []int{13, 12}, []int{13, 12}},
	}
	for _, c := range cases[0:] {
		grds := fmt.Sprint(c.grid)
		removeValueFromSet(c.d, c.allDigits, c.set, c.grid)
		if !reflect.DeepEqual(c.grid, c.gridNew) {
			t.Errorf("removeValueFromSet(%v, %v, %v, %v) grid = \n%v, want\n%v",
				c.d, c.allDigits, c.set, grds, c.grid, c.gridNew)
		}
	}
}
*/

/////////////////////////////////////// utility functions /////////////////////

func TestBinominalCoefficient(t *testing.T) {
	cases := uint(5)
	binCoeff := [][]uint{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}
	for n := range cases {
		for j := range n {
			bc := binominalCoefficient(n, j)
			if bc != binCoeff[n][j] {
				t.Errorf("binominalCoefficient(%v, %v) = %v, wwant %v",
					n, j, bc, binCoeff[n][j])
			}
		}
	}
}

// combinations return all k-element subsets from elements
func TestCombinations(t *testing.T) {
	cases := []struct {
		elements []int
		k        int
		combs    [][]int
	}{
		{[]int{1, 2}, 0,
			[][]int{{}}},
		{[]int{1, 2}, 1,
			[][]int{{1}, {2}}},
		{[]int{1, 2}, 2,
			[][]int{{1, 2}}},
		{[]int{1, 2}, 3,
			[][]int{}},
		{[]int{1, 2, 3}, 1,
			[][]int{{1}, {2}, {3}}},
		{[]int{1, 2, 3}, 2,
			[][]int{{1, 2}, {1, 3}, {2, 3}}},
		{[]int{1, 2, 3}, 3,
			[][]int{{1, 2, 3}}},
		{[]int{1, 2, 3, 4}, 1,
			[][]int{{1}, {2}, {3}, {4}}},
		{[]int{1, 2, 3, 4}, 2,
			[][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{[]int{1, 2, 3, 4}, 3,
			[][]int{{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4}}},
		{[]int{1, 2, 3, 4}, 4,
			[][]int{{1, 2, 3, 4}}},
	}
	for _, c := range cases {
		combs := combinations(c.elements, c.k)
		if !reflect.DeepEqual(combs, c.combs) {
			t.Errorf("combinations(%v, %v) = \n%v, want\n%v",
				c.elements, c.k, combs, c.combs)
		}
	}
}
