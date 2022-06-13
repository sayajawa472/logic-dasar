package pengulangan_logic01_soal03

import "fmt"

func Logic01Soal03Ke05() {
	//membuat variable j nilainya 10
	j := 10
	//membuat variable k nilainya 99
	k := 99
	// membuat deret angka menggunakan for
	// for i := 0 artinya perulangan mulainya dari variable i=0
	// i kurang dr j artinya perulangan dihentikan ketikan nilai i kurang dr j
	// i++ artinya ketika semua pernyataan dalam for selesai dijalankan , maka i++
	angka := 1

	for i := 0; i < j; i++ {
		fmt.Print(angka, "\t", k, "\t")
		angka++

	}

}
