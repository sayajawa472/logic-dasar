package logic01

import "fmt"

func Logic01Soal03() {
	//membuat variable n nilainya 10
	n := 10
	//membuat variable x nilainya 99
	x := 99
	// membuat deret angka menggunakan for
	// for i := 0 artinya perulangan mulainya dari variable i=0
	// i kurang dr n artinya perulangan dihentikan ketikan nilai i kurang dr n
	// i++ artinya ketika semua pernyataan dalam for selesai dijalankan , maka i++
	angka := 1

	for i := 0; i < n; i++ {
		fmt.Print(angka, "\t", x, "\t")
		angka++

	}

}
