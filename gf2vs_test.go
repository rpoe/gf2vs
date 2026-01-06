// Copyright 2024 Ralf Poeppel. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gf2vs

import (
	"fmt"
	"math/bits"
	"testing"
)

func TestGF2VectorSpace(t *testing.T) {
	v := GF2VectorSpace{3, 7, map[uint]uint{1: 1, 2: 2, 4: 3}, map[uint]uint{1: 1, 2: 2, 3: 4}}
	want := "{3 7 map[1:1 2:2 4:3] map[1:1 2:2 3:4]}"
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
		{0, "NewGF2VectorSpace(dim): dim = 0 < 1"},
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
		{1, "GF(2)sp {1 1 map[1:1] map[1:1]}"},
		{2, "GF(2)sp {2 3 map[1:1 2:2] map[1:1 2:2]}"},
		{3, "GF(2)sp {3 7 map[1:1 2:2 4:3] map[1:1 2:2 3:4]}"},
		{4, "GF(2)sp {4 15 map[1:1 2:2 4:3 8:4] map[1:1 2:2 3:4 4:8]}"},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.in)
		got := fmt.Sprint(sp)
		if got != c.want {
			t.Errorf("NewGF2VectorSpace(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestGF2Vector(t *testing.T) {
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

// TestNewGF2VectorString test NewGF2Vector and String
func TestNewGF2VectorString(t *testing.T) {
	cases := []struct {
		spin uint
		vin  uint
		want string
	}{
		//{1, -1, "NewGF2Vector(value): value = -1 < 0"},
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
			sp.NewGF2Vector(c.vin)
		}(c.spin, c.vin, c.want)
	}

	cases = []struct {
		spin uint
		vin  uint
		want string
	}{
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
			t.Errorf("%v.GF2Ones() = %v, want %v", c.dim, got, c.want)
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

func TestZeros(t *testing.T) {
	cases := []struct {
		dim  uint
		val  uint
		want uint
	}{
		{1, 0, 0},
		{1, 1, 0},
		{2, 0, 0},
		{2, 1, 0},
		{2, 2, 0},
		{2, 3, 0},
	}
	for _, c := range cases {
		sp := NewGF2VectorSpace(c.dim)
		v := sp.NewGF2Vector(c.val)
		vs1 := fmt.Sprint(v)
		w := sp.NewGF2Vector(c.want)
		ws := fmt.Sprint(w)
		zo := v.Zeros()
		zs1 := fmt.Sprint(zo)
		if zs1 != ws {
			t.Errorf("%v.Zero() = %v, want %v", vs1, zo, ws)
		}
		// make sure change of source is not affecting copy
		v.val++
		zs2 := fmt.Sprint(zo)
		if zs2 != zs1 {
			t.Errorf("%v.Zero() source changed to %v", zs1, zs2)
		}
		// make sure change of copy is not affecting source
		v = sp.NewGF2Vector(c.val)
		vs1 = fmt.Sprint(v)
		zo = v.Zeros()
		zo.val++
		vs2 := fmt.Sprint(v)
		if vs2 != vs1 {
			t.Errorf("%v.Zero() copy changed to %v", vs1, vs2)
		}
	}
}

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
