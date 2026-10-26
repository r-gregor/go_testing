package main

import (
	"fmt"
	"strings"
)

type LCS struct {
	S1 string
	S2 string
	n  int
	ss string
	rt float64
}

func main() {

	// Case INSensitive with NO spaces!!
	p1 := LCS{S1: "Gregor R", S2: "Gregor Redelonghi"}
	p2 := LCS{S1: "Gregor R", S2: "Gregor Redelonghi"}
	p3 := LCS{S1: "Tadeja Mali", S2: "Tadeja Mali Redelonghi"}
	p4 := LCS{S1: "York", S2: "New York"}
	p5 := LCS{S1: "york", S2: "New York"}

	lcss := []LCS{p1, p2, p3, p4, p5}

	for _, el := range lcss {
		el.PrintLCS()
	}

	p6 := LCS{S1: "Mex", S2: "New Mexico City"}
	p6.PrintLCS()

} // end main

func Ratio(n int, Y string) float64 {
	L1 := float64(n)
	L2 := float64(len(Y))
	ratio := L1 / L2
	return ratio
}

func (lcs *LCS) CalculateLCS() {
	lcs.n, lcs.ss = LCS1(lcs.S1, lcs.S2)
	lcs.rt = Ratio(lcs.n, lcs.S2)
	// fmt.Printf("LCS of [%s] and [%s] is a sttring [%s] of length %d\n", lcs.S1, lcs.S2, lcs.ss, lcs.n)
	// fmt.Printf("Ratio of substring/string lengths is: %.3f\n", lcs.rt)
}

func (lcs *LCS) PrintLCS() {
	lcs.CalculateLCS()
	fmt.Printf("LCS of [%s] and [%s] is a string [%s] of length %d\n", lcs.S1, lcs.S2, lcs.ss, lcs.n)
	fmt.Printf("Ratio of substring/string lengths is: %.3f\n", lcs.rt)
}

func LCS1(a1, b1 string) (int, string) {

	a := strings.TrimSpace(strings.ToLower(a1))
	b := strings.TrimSpace(strings.ToLower(a1))

	arunes := []rune(a)
	brunes := []rune(b)
	aLen := len(arunes)
	bLen := len(brunes)
	lengths := make([][]int, aLen+1)
	for i := 0; i <= aLen; i++ {
		lengths[i] = make([]int, bLen+1)
	}

	// row 0 and column 0 are initialized to 0 already
	for i := 0; i < aLen; i++ {
		for j := 0; j < bLen; j++ {
			if arunes[i] == brunes[j] {
				lengths[i+1][j+1] = lengths[i][j] + 1
			} else if lengths[i+1][j] > lengths[i][j+1] {
				lengths[i+1][j+1] = lengths[i+1][j]
			} else {
				lengths[i+1][j+1] = lengths[i][j+1]
			}
		}
	}

	// read the substring out from the matrix
	s := make([]rune, 0, lengths[aLen][bLen])
	for x, y := aLen, bLen; x != 0 && y != 0; {
		if lengths[x][y] == lengths[x-1][y] {
			x--
		} else if lengths[x][y] == lengths[x][y-1] {
			y--
		} else {
			s = append(s, arunes[x-1])
			x--
			y--
		}
	}
	// reverse string
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return len(s), strings.ToLower(string(s))
}
