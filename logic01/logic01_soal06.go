package logic01

import (
	"fmt"
	"math"
)

func Logic01Soal06() {
	n := 15
	nt := n / 2 // nt : nilai tengah
	x := 3      //x := 3 , krn angka awal 3
	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			z := math.Pow(float64(i), 2) //math.pow  = kuadrat
			// z deklarasi variable baru , nilail math pow disimpan ke z
			// math.pow harus integer
			fmt.Print(z, "\t")
		} else {
			fmt.Print(x, "\t")
		}
		if i <= nt {
			x += 3
		} else {
			x -= 3
		}
	}
}
