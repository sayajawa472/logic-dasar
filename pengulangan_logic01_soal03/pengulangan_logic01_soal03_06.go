package pengulangan_logic01_soal03

import "fmt"

func Logic01Soal03Ke06() {
	//membuat variable l nilainya 10
	l := 10
	//membuat variable m nilainya 99
	m := 99
	// membuat deret angka menggunakan for
	// for i := 0 artinya perulangan mulainya dari variable i=0
	// i kurang dr l artinya perulangan dihentikan ketikan nilai i kurang dr l
	// i++ artinya ketika semua pernyataan dalam for selesai dijalankan , maka i++
	angka := 1

	for i := 0; i < l; i++ {
		fmt.Print(angka, "\t", m, "\t")
		angka++

	}

}
