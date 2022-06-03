package pengulangan_logic01_soal03

import "fmt"

func main() {
	Logic01Soal03Ke07()
}
func Logic01Soal03Ke07() {
	//membuat variable o nilainya 10
	o := 10
	//membuat variable p nilainya 99
	p := 99
	// membuat deret angka menggunakan for
	// for i := 0 artinya perulangan mulainya dari variable i=0
	// i kurang dr o artinya perulangan dihentikan ketikan nilai i kurang dr o
	// i++ artinya ketika semua pernyataan dalam for selesai dijalankan , maka i++
	angka := 1

	for i := 0; i < o; i++ {
		fmt.Print(angka, "\t", p, "\t")
		angka++

	}

}
