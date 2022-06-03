package pengulangan_logic01_soal02

import (
	"fmt"
	"testing"
)

func TestLogic01Soalo2ke09(t *testing.T) {
	j := 10
	angka := 1
	// for i := 0 artinya perulangan mulainya dari variable i=0
	// i kurang dr j artinya perulangan berlanjut selama nilai i kurang dr j
	//dan dihentikan ketika nilai i = nilai j
	// i++ artinya ketika semua pernyataan dalam for selesai dijalankan , maka i++
	for i := 0; i < j; i++ {
		if i%2 == 1 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(3, "\t")
		}
	}
}
