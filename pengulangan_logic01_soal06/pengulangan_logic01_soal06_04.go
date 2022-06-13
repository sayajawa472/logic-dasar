package pengulangan_logic01_soal06

import (
	"fmt"
	"math"
	"testing"
)

func TestSoalNo6Ke04(t *testing.T) {
	o := 15
	nt := o / 2
	p := 3
	for i := 1; i <= o; i++ {
		if i%4 == 0 {
			z := math.Pow(float64(i), 2)
			fmt.Print(z, "\t")
		} else {
			fmt.Print(p, "\t")
		}
		if i <= nt {
			p += 3
		} else {
			p -= 3
		}
	}
}
