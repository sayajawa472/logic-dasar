package pengulangan_logic01_soal04

import "fmt"

func main() {
	Logic01Soal04Ke03()
}

func Logic01Soal04Ke03() {
	n := 10
	x := 777
	// membuat deret angka menggunakan for
	// for i := 0 artinya perulangan mulainya dari variable i=0
	// i kurang dr n artinya perulangan dihentikan ketikan nilai i kurang dr n
	// i++ artinya ketika semua pernyataan dalam for selesai dijalankan , maka i++
	angka := 1

	for i := 0; i < n; i++ {
		fmt.Print(x, "\t", angka, "\t")
		angka++

	}
}
