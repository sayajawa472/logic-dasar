package pengulangan_logic01_soal03

import "fmt"

func main() {
	Logic01Soal03Ke10()
}
func Logic01Soal03Ke10() {
	//membuat variable y nilainya 10
	y := 10
	//membuat variable z nilainya 99
	z := 99
	// membuat deret angka menggunakan for
	// for i := 0 artinya perulangan mulainya dari variable i=0
	// i kurang dr y artinya perulangan dihentikan ketikan nilai i kurang dr y
	// i++ artinya ketika semua pernyataan dalam for selesai dijalankan , maka i++
	angka := 1

	for i := 0; i < y; i++ {
		fmt.Print(angka, "\t", z, "\t")
		angka++

	}

}
