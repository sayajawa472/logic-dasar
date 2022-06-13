package pengulangan_logic01_soal03

import "fmt"

func Logic01Soal03ke04() {
	//membuat variable g nilainya 10
	g := 10
	//membuat variable h nilainya 99
	h := 99
	// membuat deret angka menggunakan for
	// for i := 0 artinya perulangan mulainya dari variable i=0
	// i kurang dr g artinya perulangan dihentikan ketikan nilai i kurang dr g
	// i++ artinya ketika semua pernyataan dalam for selesai dijalankan , maka i++
	angka := 1

	for i := 0; i < g; i++ {
		fmt.Print(angka, "\t", h, "\t")
		angka++

	}

}
