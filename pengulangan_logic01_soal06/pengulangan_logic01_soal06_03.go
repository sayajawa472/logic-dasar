package pengulangan_logic01_soal06

import (
	"fmt"
	"math"
	"testing"
)

func TestSoalNo6Ke03(t *testing.T) {
	g := 15
	nt := g / 2
	h := 3

	for i := 1; i <= g; i++ {
		if i%4 == 0 {
			j := math.Pow(float64(i), 2)
			fmt.Print(j, "\t")
		}
		if i <= nt {
			h += 3
		} else {
			h -= 3

		}
	}

}
