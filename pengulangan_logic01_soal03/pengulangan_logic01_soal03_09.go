package pengulangan_logic01_soal03

import "fmt"

func main() {
	Logic01Soal03Ke09()
}
func Logic01Soal03Ke09() {
	//membuat variable t nilainya 10
	t := 10
	//membuat variable u nilainya 99
	u := 99
	// membuat deret angka menggunakan for
	// for i := 0 artinya perulangan mulainya dari variable i=0
	// i kurang dr t artinya perulangan dihentikan ketikan nilai i kurang dr t
	// i++ artinya ketika semua pernyataan dalam for selesai dijalankan , maka i++
	angka := 1

	for i := 0; i < t; i++ {
		fmt.Print(angka, "\t", u, "\t")
		angka++

	}

}
