package main

import "fmt"

func main() {

	// Case Sensitive with spaces!!
	CalculateLCS("Gregor R", "Gregor Redelonghi")
	CalculateLCS("Tadeja Mali", "Tadeja Mali Redelonghi")
	CalculateLCS("York", "New York")
	CalculateLCS("york", "New York")
}

func Ratio(n int, Y string) float64 {
	L1 := float64(n)
	L2 := float64(len(Y))
	ratio := L1 / L2
	return ratio
}

func CalculateLCS(S1, S2 string) {
	n, s := LCS1(S1, S2)
	ratio := Ratio(n, S2)
	fmt.Printf("LCS of [%s] and [%s] is a sttring [%s] of length %d\n", S1, S2, s, n)
	fmt.Printf("Ratio of substring/string lengths is: %.3f\n", ratio)
}

func LCS1(a, b string) (int, string) {
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
	return len(s), string(s)
}
