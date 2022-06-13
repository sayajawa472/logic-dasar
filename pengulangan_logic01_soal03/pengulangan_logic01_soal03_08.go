package pengulangan_logic01_soal03

import "fmt"

func Logic01Soal03Ke08() {
	//membuat variable q nilainya 10
	q := 10
	//membuat variable s nilainya 99
	s := 99
	// membuat deret angka menggunakan for
	// for i := 0 artinya perulangan mulainya dari variable i=0
	// i kurang dr q artinya perulangan dihentikan ketikan nilai i kurang dr q
	// i++ artinya ketika semua pernyataan dalam for selesai dijalankan , maka i++
	angka := 1

	for i := 0; i < q; i++ {
		fmt.Print(angka, "\t", s, "\t")
		angka++

	}

}
