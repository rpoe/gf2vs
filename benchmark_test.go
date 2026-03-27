// Ralf Poeppel, 2026

// Bnechmark tests to select library functions for implementation
package gf2vs

import (
	"flag"
	"math/bits"
	"testing"
)

const N = 7
const M = 100

var VSP [N]*GF2VectorSpace
var V [N][M]*GF2Vector

func TestMain(m *testing.M) {
	j := uint(1)
	for i := 0; i < N; i++ {
		k := j
		if j > 8 {
			k = j - 1
		}
		VSP[i] = NewGF2VectorSpace(k)
		ones := VSP[i].ones
		for l := 0; l < M; l++ {
			V[i][l] = VSP[i].NewGF2Vector(uint(l) % (ones + 1))
		}
		j <<= 1
	}

	flag.Parse()
	m.Run()
}

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

func BenchmarkVectorSpaceNewVector(b *testing.B) {
	var i, j uint
	var sp *GF2VectorSpace
	for b.Loop() {
		for j = 1; j < N; j <<= 1 {
			//sp := NewGF2VectorSpace(j)
			sp = VSP[j]
			//fmt.Println(sp)
			for i = 1; i < j; i++ {
				sp.NewGF2Vector(i)
			}
		}
	}
}

func BenchmarkVectorSpaceVectorIndexMap(b *testing.B) {
	var j uint
	for b.Loop() {
		for j = 1; j < N; j <<= 1 {
			vs := V[j]
			for _, v := range vs {
				v.IndexMap()
			}
		}
	}
}

func BenchmarkVectorSpaceVectorIndex(b *testing.B) {
	var j uint
	for b.Loop() {
		for j = 1; j < N; j <<= 1 {
			vs := V[j]
			for _, v := range vs {
				v.Index()
			}
		}
	}
}

func BenchmarkVectorSpaceVectorIsBaseVector(b *testing.B) {
	var j uint
	for b.Loop() {
		for j = 1; j < N; j <<= 1 {
			vs := V[j]
			for _, v := range vs {
				v.IsBaseVector()
			}
		}
	}
}
