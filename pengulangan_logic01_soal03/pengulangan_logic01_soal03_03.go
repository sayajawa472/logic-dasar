package pengulangan_logic01_soal03

import (
	"fmt"
)

func Logic01Soal03Ke03() {
	//membuat variable e nilainya 10
	e := 10
	//membuat variable f nilainya 99
	f := 99
	// membuat deret angka menggunakan for
	// for i := 0 artinya perulangan mulainya dari variable i=0
	// i kurang dr e artinya perulangan dihentikan ketikan nilai i kurang dr e
	// i++ artinya ketika semua pernyataan dalam for selesai dijalankan , maka i++
	angka := 1

	for i := 0; i < e; i++ {
		fmt.Print(angka, "\t", f, "\t")
		angka++

	}

}
