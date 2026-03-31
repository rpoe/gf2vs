// Ralf Poeppel, 2026

package gf2vs

import (
	"fmt"
	"math/big"
	"testing"
)

func TestText(t *testing.T) {
	cases := []struct {
		rows []string
		base int
		want string
	}{
		{[]string{"0"}, 2, "0\n"},
		{[]string{"1", "0"}, 2, "1\n0\n"},
		{[]string{"3", "1"}, 2, "11\n01\n"},
		{[]string{"7", "1", "3"}, 2, "111\n001\n011\n"},
	}
	for _, c := range cases {
		n := len(c.rows)
		bm := make(BitMatrix, n)
		var ok bool
		for i, r := range c.rows {
			if bm[i], ok = big.NewInt(0).SetString(r, 0); !ok {
				t.Fatalf("SetString(%v, 0) = false", r)
			}
		}
		got := bm.Text(c.base, "\n")
		if got != c.want {
			t.Errorf("%v.Text(%v, \\n) =\n%v, want\n%v", bm, c.base, got, c.want)
		}
	}
}

func TestCmp(t *testing.T) {
	cases := []struct {
		x BitMatrix
		y BitMatrix
		r int
	}{
		{BitMatrix{}, BitMatrix{}, 0},
		{BitMatrix{}, BitMatrix{big.NewInt(1)}, -1},
		{BitMatrix{big.NewInt(1)}, BitMatrix{}, 1},
		{BitMatrix{big.NewInt(0)}, BitMatrix{big.NewInt(1)}, -1},
		{BitMatrix{big.NewInt(1)}, BitMatrix{big.NewInt(1)}, 0},
		{BitMatrix{big.NewInt(1)}, BitMatrix{big.NewInt(0)}, 1},
		{BitMatrix{big.NewInt(0)},
			BitMatrix{big.NewInt(1), big.NewInt(0)}, -1},
		{BitMatrix{big.NewInt(0), big.NewInt(1)},
			BitMatrix{big.NewInt(1)}, 1},
		{BitMatrix{big.NewInt(0), big.NewInt(1)},
			BitMatrix{big.NewInt(1), big.NewInt(0)}, -1},
		{BitMatrix{big.NewInt(1), big.NewInt(1)},
			BitMatrix{big.NewInt(1), big.NewInt(1)}, 0},
		{BitMatrix{big.NewInt(1), big.NewInt(0)},
			BitMatrix{big.NewInt(0), big.NewInt(1)}, 1},
	}
	for _, c := range cases {
		g := (&c.x).Cmp(&c.y)
		if g != c.r {
			t.Errorf("\n%v.Cmp(\n%v) = %v, want %v", &c.x, &c.y, g, c.r)
		}
	}
}

func TestSet(t *testing.T) {
	cases := []struct {
		x BitMatrix
		z BitMatrix
	}{
		{BitMatrix{}, BitMatrix{}},
		{BitMatrix{big.NewInt(1)}, BitMatrix{big.NewInt(1)}},
		{BitMatrix{big.NewInt(1), big.NewInt(2)}, BitMatrix{big.NewInt(1), big.NewInt(2)}},
	}
	for _, c := range cases {
		z := BitMatrix{}
		g := (&z).Set(&c.x)
		if g.Cmp(&c.z) != 0 {
			t.Errorf("Set(\n%v) = \n%v, want\n%v", &c.x, g, &c.z)
		}
	}
}

type RrefTestCase struct {
	in   BitMatrix
	mr   int
	want BitMatrix
	rank int
	ok   bool
}

func RunTestRowReducedEcholonForm(t *testing.T, cases []RrefTestCase) {
	for _, c := range cases {
		cinstr := c.in.Text(2, "\n")
		rank, ok := c.in.RowReducedEcholonForm(c.mr)
		sgot := c.in.Text(2, "\n")
		swant := c.want.Text(2, "\n")
		if sgot != swant || rank != c.rank || ok != c.ok {
			t.Errorf("\n%v.RowReducedEcholonForm(%v) == %v, %v\n%v\nwant %v, %v\n%v",
				cinstr, c.mr, rank, ok, sgot, c.rank, c.ok, swant)
		}
	}
}

func TestRowReducedEcholonForm(t *testing.T) {
	cases := []RrefTestCase{
		{BitMatrix{}, 0, BitMatrix{}, 0, true},
		{BitMatrix{big.NewInt(1)}, 0, BitMatrix{big.NewInt(1)}, 1, true},
		{BitMatrix{big.NewInt(1)}, 1, BitMatrix{big.NewInt(1)}, 1, false},
		{BitMatrix{big.NewInt(1)}, 2, BitMatrix{big.NewInt(1)}, 1, false},
		{BitMatrix{big.NewInt(1), big.NewInt(2)}, 0,
			BitMatrix{big.NewInt(2), big.NewInt(1)}, 2, true},
		{BitMatrix{big.NewInt(1), big.NewInt(3)}, 0,
			BitMatrix{big.NewInt(2), big.NewInt(1)}, 2, true},
		{BitMatrix{big.NewInt(3), big.NewInt(1)}, 0,
			BitMatrix{big.NewInt(2), big.NewInt(1)}, 2, true},
		{BitMatrix{big.NewInt(1), big.NewInt(2)}, 1,
			BitMatrix{big.NewInt(2), big.NewInt(1)}, 2, false},
		{BitMatrix{big.NewInt(1), big.NewInt(2)}, 2,
			BitMatrix{big.NewInt(2), big.NewInt(1)}, 2, false},
		{BitMatrix{big.NewInt(4), big.NewInt(5)}, 0,
			BitMatrix{big.NewInt(4), big.NewInt(1)}, 2, true},
		{BitMatrix{big.NewInt(1), big.NewInt(3), big.NewInt(7)}, 0,
			BitMatrix{big.NewInt(4), big.NewInt(2), big.NewInt(1)}, 3, true},
		// example on 2024-06-01 from
		// https://en.wikipedia.org/w/index.php?title=Boolean_satisfiability_problem&oldid=1219369085
		// Associated coefficient matrix
		//
		// a	b	c	d		line
		//
		// 1	0	1	1	1	A
		// 0	1	1	1	0	B
		// 1	1	0	1	0	C
		// 1	1	1	0	1	D
		{BitMatrix{
			big.NewInt(0b10111),
			big.NewInt(0b01110),
			big.NewInt(0b11010),
			big.NewInt(0b11101)}, 1,
			BitMatrix{big.NewInt(16), big.NewInt(9), big.NewInt(4), big.NewInt(3)}, 4, true},
		// example Gauss3Bits.txt
		{
			BitMatrix{
				big.NewInt(0b11000110),
				big.NewInt(0b00100100),
				big.NewInt(0b00011011),
				big.NewInt(0b00110110),
				big.NewInt(0b10000010),
				big.NewInt(0b01001101),
			}, 3,
			BitMatrix{
				big.NewInt(0b10000010),
				big.NewInt(0b01000100),
				big.NewInt(0b00100100),
				big.NewInt(0b00010010),
				big.NewInt(0b00001001),
			}, 5, true,
		},
		// one line test 2022-10-05
		{
			BitMatrix{
				big.NewInt(0b1_1),
			}, 0,
			BitMatrix{
				big.NewInt(0b1_1),
			}, 1, true,
		},
		// 2024-04-22 Inverse of 2x2 Matrix
		{
			BitMatrix{
				big.NewInt(0b10_10),
				big.NewInt(0b01_01),
			}, 2,
			BitMatrix{
				big.NewInt(0b10_10),
				big.NewInt(0b01_01),
			}, 2, true,
		},
		{
			BitMatrix{
				big.NewInt(0b10_10),
				big.NewInt(0b11_01),
			}, 2,
			BitMatrix{
				big.NewInt(0b10_10),
				big.NewInt(0b01_11),
			}, 2, true,
		},
		{
			BitMatrix{
				big.NewInt(0b11_10),
				big.NewInt(0b01_01),
			}, 2,
			BitMatrix{
				big.NewInt(0b10_11),
				big.NewInt(0b01_01),
			}, 2, true,
		},
		{
			BitMatrix{
				big.NewInt(0b01_10),
				big.NewInt(0b10_01),
			}, 2,
			BitMatrix{
				big.NewInt(0b10_01),
				big.NewInt(0b01_10),
			}, 2, true,
		},
		{
			BitMatrix{
				big.NewInt(0b11_10),
				big.NewInt(0b10_01),
			}, 2,
			BitMatrix{
				big.NewInt(0b10_01),
				big.NewInt(0b01_11),
			}, 2, true,
		},
		{
			BitMatrix{
				big.NewInt(0b01_10),
				big.NewInt(0b11_01),
			}, 2,
			BitMatrix{
				big.NewInt(0b10_11),
				big.NewInt(0b01_10),
			}, 2, true,
		},
		{
			BitMatrix{
				big.NewInt(0x00_00_00_00_0F_FF_FF_FF),
				big.NewInt(0x00_00_00_00_0F_00_00_0F),
			}, 0,
			BitMatrix{
				big.NewInt(0x00_00_00_00_0F_00_00_0F),
				big.NewInt(0x00_00_00_00_00_FF_FF_F0),
			}, 2, true,
		},
		// 2026-03-25 https://en.wikipedia.org/w/index.php?title=XOR-SAT&oldid=1322926005
		{
			BitMatrix{
				big.NewInt(0b1101_0),
				big.NewInt(0b0111_0),
				big.NewInt(0b1110_0),
				big.NewInt(0b1101_1),
			}, 1,
			BitMatrix{
				big.NewInt(0b10010),
				big.NewInt(0b01000),
				big.NewInt(0b00110),
				big.NewInt(0b00001),
			}, 4, false,
		},
		// linear dependent line added
		{
			BitMatrix{
				big.NewInt(0b1101_0),
				big.NewInt(0b0111_0),
				big.NewInt(0b1110_0),
				big.NewInt(0b1101_1),
				big.NewInt(0b0011_0),
			}, 1,
			BitMatrix{
				big.NewInt(0b10010),
				big.NewInt(0b01000),
				big.NewInt(0b00110),
				big.NewInt(0b00001),
			}, 4, false,
		},
		// solvable version
		{
			BitMatrix{
				big.NewInt(0b1101_0),
				big.NewInt(0b0111_0),
				big.NewInt(0b1110_0),
			}, 1,
			BitMatrix{
				big.NewInt(0b10010),
				big.NewInt(0b01000),
				big.NewInt(0b00110),
			}, 3, true,
		},
		// linear dependent line added
		{
			BitMatrix{
				big.NewInt(0b1101_0),
				big.NewInt(0b0111_0),
				big.NewInt(0b1110_0),
				big.NewInt(0b0011_0),
			}, 1,
			BitMatrix{
				big.NewInt(0b10010),
				big.NewInt(0b01000),
				big.NewInt(0b00110),
			}, 3, true,
		},
	}

	RunTestRowReducedEcholonForm(t, cases[0:])
}

func TestLeftRightSplitter(t *testing.T) {
	type Case struct {
		bg     *big.Int
		wleft  *big.Int
		wright *big.Int
	}
	splits := []struct {
		mr    int // number of bits on right side
		cases []Case
	}{
		{
			0,
			[]Case{
				{big.NewInt(0b0), big.NewInt(0b0), big.NewInt(0b0)},
				{big.NewInt(0b1), big.NewInt(0b1), big.NewInt(0b0)},
			},
		},
		{
			1,
			[]Case{
				{big.NewInt(0b0), big.NewInt(0b0), big.NewInt(0b0)},
				{big.NewInt(0b1), big.NewInt(0b0), big.NewInt(0b1)},
				{big.NewInt(0b10), big.NewInt(0b10), big.NewInt(0b0)},
				{big.NewInt(0b11), big.NewInt(0b10), big.NewInt(0b1)},
				{big.NewInt(0b110), big.NewInt(0b110), big.NewInt(0b0)},
				{big.NewInt(0b111), big.NewInt(0b110), big.NewInt(0b1)},
			},
		},
		{
			2,
			[]Case{
				{big.NewInt(0b110), big.NewInt(0b100), big.NewInt(0b10)},
				{big.NewInt(0b111), big.NewInt(0b100), big.NewInt(0b11)},
				{big.NewInt(0b1110), big.NewInt(0b1100), big.NewInt(0b10)},
				{big.NewInt(0b1111), big.NewInt(0b1100), big.NewInt(0b11)},
			},
		},
	}
	for _, s := range splits {
		lrSplitter := LeftRightSplitter(s.mr)
		for _, c := range s.cases {
			left, right := lrSplitter(c.bg)
			if left.Cmp(c.wleft) != 0 || right.Cmp(c.wright) != 0 {
				cbgs := c.bg.Text(2)
				lefts := left.Text(2)
				rights := right.Text(2)
				clefts := c.wleft.Text(2)
				crights := c.wright.Text(2)
				t.Errorf("LeftRightSplitter(%v) lrSplitter(0b%v) = \n0b%v, 0b%v, want \n0b%v, 0b%v",
					s.mr, cbgs, lefts, rights, clefts, crights)
			}
		}
	}
}

type SolutionTestCase struct {
	in   BitMatrix
	mr   int
	want []int
	werr string // expected error
}

func RunTestSolutionFromRref(t *testing.T, cases []SolutionTestCase) {
	for _, c := range cases {
		cinstr := c.in.String()
		got, err := c.in.SolutionFromRref(c.mr)
		if fmt.Sprint(got) != fmt.Sprint(c.want) ||
			(len(c.werr) == 0) != (err == nil) {
			t.Errorf("\n%v.SolutionFromRref(%v) == \n"+
				"%v\n%v\nwant \n%v\n%v",
				cinstr, c.mr, got, err, c.want, c.werr)
		}
	}
}

func TestSolutionFromRref(t *testing.T) {
	cases := []SolutionTestCase{
		{BitMatrix{}, 0, []int{},
			"No values on right side"},
		{BitMatrix{big.NewInt(1)}, 0, []int{},
			"No values on right side"},
		{BitMatrix{big.NewInt(1)}, 1, []int{},
			"Row length=1 to short, #rows:ln=1, #right:mr=1"},
		{BitMatrix{big.NewInt(1)}, 2, []int{},
			"Row length=1 to short, #rows:ln=1, #right:mr=2"},
		{BitMatrix{big.NewInt(3)}, 1, []int{1}, ""},
		// example on 2024-06-01 from
		// https://en.wikipedia.org/w/index.php?title=Boolean_satisfiability_problem&oldid=1219369085
		// Associated coefficient matrix rref
		//
		// a  b  c  d     line
		//
		// 1  0  0  0  0  A
		// 0  1  0  0  1  B
		// 0  0  1  0  0  C
		// 0  0  0  1  1  D
		{BitMatrix{
			big.NewInt(0b1000_0),
			big.NewInt(0b0100_1),
			big.NewInt(0b0010_0),
			big.NewInt(0b0001_1),
		},
			1,
			[]int{0, 1, 0, 1}, ""},
		// example Gauss3Bits.txt
		{BitMatrix{big.NewInt(0x82), big.NewInt(0x44), big.NewInt(0x24),
			big.NewInt(0x12), big.NewInt(9)}, 3,
			[]int{2, 4, 4, 2, 1}, ""},
		// simple test
		{BitMatrix{
			big.NewInt(0b1_1),
		}, 1,
			[]int{1}, "",
		},
		// a row is not there
		{BitMatrix{
			big.NewInt(0b100_1),
			big.NewInt(0b001_1),
		}, 1,
			[]int{1, -1, 1}, "Partial Solution returned",
		},
		// 2 rows missing
		{BitMatrix{
			big.NewInt(0b10000_100),
			big.NewInt(0b00100_010),
			big.NewInt(0b00010_001),
		}, 3,
			[]int{4, -1, 2, 1, -1}, "Partial Solution returned",
		},
		{BitMatrix{
			big.NewInt(0b100000_100),
			big.NewInt(0b010000_000),
			big.NewInt(0b000100_010),
			big.NewInt(0b000010_001),
		}, 3,
			[]int{4, 0, -1, 2, 1, -1}, "Partial Solution returned",
		},
		{BitMatrix{
			big.NewInt(0b100000_100),
			big.NewInt(0b010000_001),
			big.NewInt(0b000100_010),
			big.NewInt(0b000001_000),
		}, 3,
			[]int{4, 1, -1, 2, -1, 0}, "Partial Solution returned",
		},
		{BitMatrix{
			big.NewInt(0b100000000000000_1000),
			big.NewInt(0b010000000000000_0100),
			big.NewInt(0b001000000000000_0010),
			big.NewInt(0b000100000000000_0100),
			big.NewInt(0b000000010000000_1000),
			big.NewInt(0b000000000001000_0010),
			big.NewInt(0b000000000000100_0001),
			big.NewInt(0b000000000000010_1000),
		}, 4,
			[]int{8, 4, 2, 4, -1, -1, -1, 8, -1, -1, -1, 2, 1, 8, -1},
			"Partial Solution returned",
		},
		// 2026-03-25 https://en.wikipedia.org/w/index.php?title=XOR-SAT&oldid=1322926005
		// solvable version
		{
			BitMatrix{
				big.NewInt(0b1001),
				big.NewInt(0b0100),
				big.NewInt(0b0011),
			}, 1,
			[]int{1, 0, 1}, "",
		},
		// original version
		{
			BitMatrix{
				big.NewInt(0b10010),
				big.NewInt(0b01000),
				big.NewInt(0b00110),
				big.NewInt(0b00001),
			}, 1,
			[]int{}, "Contradiction of equations of echolon form",
		},
	}

	RunTestSolutionFromRref(t, cases)
}

func RunTestXorSatSolve(t *testing.T, cases []SolutionTestCase) {
	for _, c := range cases {
		cinstr := c.in.String()
		got, _, err := c.in.XorSatSolve(c.mr)
		if fmt.Sprint(got) != fmt.Sprint(c.want) ||
			(len(c.werr) == 0) != (err == nil) {
			t.Errorf("\n%v.XorSatSolve(%v) == \n"+
				"%v\n%v\nwant \n%v\n%v",
				cinstr, c.mr, got, err, c.want, c.werr)
		}
	}
}

func TestXorSatSolve(t *testing.T) {
	cases := []SolutionTestCase{
		{BitMatrix{}, 0, []int{},
			"No values on right side"},
		{BitMatrix{big.NewInt(1)}, 0, []int{},
			"No values on right side"},
		{BitMatrix{big.NewInt(1)}, 1, []int{},
			"Row length=1 to short, #rows:ln=1, #right:mr=1"},
		{BitMatrix{big.NewInt(1)}, 2, []int{},
			"Row length=1 to short, #rows:ln=1, #right:mr=2"},
		{BitMatrix{big.NewInt(3)}, 1, []int{1}, ""},
		// example on 2019-10-13 from
		// https://en.m.wikipedia.org/wiki/Boolean_satisfiability_problem
		{BitMatrix{big.NewInt(23), big.NewInt(14), big.NewInt(26), big.NewInt(29)}, 1,
			[]int{0, 1, 0, 1}, ""},
		{BitMatrix{big.NewInt(16), big.NewInt(9), big.NewInt(4), big.NewInt(3)}, 1,
			[]int{0, 1, 0, 1}, ""},
		// example Gauss3Bits.txt
		{BitMatrix{big.NewInt(0xc6), big.NewInt(0x24), big.NewInt(0x1b),
			big.NewInt(0x36), big.NewInt(0x82), big.NewInt(0x4d)}, 3,
			[]int{2, 4, 4, 2, 1}, ""},
		{BitMatrix{big.NewInt(0x82), big.NewInt(0x44), big.NewInt(0x24),
			big.NewInt(0x12), big.NewInt(9)}, 3,
			[]int{2, 4, 4, 2, 1}, ""},
		{BitMatrix{
			big.NewInt(0b1_1)},
			1,
			[]int{1}, ""},
		// 2026-03-25 https://en.wikipedia.org/w/index.php?title=XOR-SAT&oldid=1322926005
		// solvable version
		{BitMatrix{
			big.NewInt(0b1101),
			big.NewInt(0b0111),
			big.NewInt(0b1110),
		}, 1,
			[]int{1, 0, 1}, "",
		},
		// original version
		{
			BitMatrix{
				big.NewInt(0b1101_0),
				big.NewInt(0b0111_0),
				big.NewInt(0b1110_0),
				big.NewInt(0b1101_1),
			}, 1,
			[]int{}, "Contradiction of equations of BitMatrix",
		},
	}

	RunTestXorSatSolve(t, cases)
}
