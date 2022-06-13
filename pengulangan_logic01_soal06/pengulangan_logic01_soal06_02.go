package pengulangan_logic01_soal06

import (
	"fmt"
	"math"
	"testing"
)

func TestSoalNo6Ke02(t *testing.T) {
	d := 15
	nt := d / 2
	e := 3

	for i := 1; i <= d; i++ {
		if i%4 == 0 {
			f := math.Pow(float64(i), 2)
			fmt.Print(f, "\t")
		} else {
			fmt.Print(e, "\t")
		}
		if i <= nt {
			e += 3
		} else {
			e -= 3
		}

	}
}
