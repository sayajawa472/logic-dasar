package pengulangan_logic01_soal02

import (
	"fmt"
	"testing"
)

func TestLogic01Soalo2ke05(t *testing.T) {
	e := 10
	angka := 1
	// for i := 0 artinya perulangan mulainya dari variable i=0
	// i kurang dr e artinya perulangan berlanjut selama nilai i kurang dr e
	//dan dihentikan ketika nilai i = nilai e
	// i++ artinya ketika semua pernyataan dalam for selesai dijalankan , maka i++
	for i := 0; i < e; i++ {
		if i%2 == 1 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(3, "\t")
		}
	}
}
