package logic01

import (
	"fmt"
	"testing"
)

func TestLogic01Soalo2(t *testing.T) {
	n := 10
	angka := 1
	// for i := 0 artinya perulangan mulainya dari variable i=0
	// i kurang dr n artinya perulangan berlanjut selama nilai i kurang dr n
	//dan dihentikan ketika nilai i = nilai n
	// i++ artinya ketika semua pernyataan dalam for selesai dijalankan , maka i++
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(3, "\t")
		}
	}
}
