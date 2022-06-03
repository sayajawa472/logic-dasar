package pengulangan_logic01_soal02

import (
	"fmt"
	"testing"
)

func TestLogic01Soalo2ke03(t *testing.T) {
	c := 10
	angka := 1
	// for i := 0 artinya perulangan mulainya dari variable i=0
	// i kurang dr c artinya perulangan berlanjut selama nilai i kurang dr c
	//dan dihentikan ketika nilai i = nilai c
	// i++ artinya ketika semua pernyataan dalam for selesai dijalankan , maka i++
	for i := 0; i < c; i++ {
		if i%2 == 1 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(3, "\t")
		}
	}
}
