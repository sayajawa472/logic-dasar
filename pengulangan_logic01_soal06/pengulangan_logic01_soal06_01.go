package pengulangan_logic01_soal06

import (
	"fmt"
	"math"
	"testing"
)

func TestSoalNo6Ke01(t *testing.T) {
	a := 15
	nt := a / 2
	b := 3

	for i := 1; i <= a; i++ {
		if i%4 == 0 {
			c := math.Pow(float64(i), 2)
			fmt.Print(c, "\t")
		} else {
			fmt.Print(b, "\t")
		}
		if i <= nt {
			b += 3
		} else {
			b -= 3
		}

	}
}
