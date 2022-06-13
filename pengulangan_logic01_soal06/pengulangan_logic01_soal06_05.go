package pengulangan_logic01_soal06

import (
	"fmt"
	"math"
	"testing"
)

func TestSoalNo6Ke05(t *testing.T) {
	r := 15
	nt := r / 2
	s := 3
	for i := 1; i <= r; i++ {
		if i%4 == 0 {
			t := math.Pow(float64(i), 2)
			fmt.Print(t, "\t")
		} else {
			fmt.Print(s, "\t")
		}
		if i <= nt {
			s += 3
		} else {
			s -= 3
		}
	}
}
