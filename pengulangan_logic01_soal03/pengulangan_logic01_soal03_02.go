package pengulangan_logic01_soal03

import "fmt"

func main() {
	Logic01Soal03Ke02()
}
func Logic01Soal03Ke02() {
	//membuat variable c nilainya 10
	c := 10
	//membuat variable d nilainya 99
	d := 99
	// membuat deret angka menggunakan for
	// for i := 0 artinya perulangan mulainya dari variable i=0
	// i kurang dr c artinya perulangan dihentikan ketikan nilai i kurang dr c
	// i++ artinya ketika semua pernyataan dalam for selesai dijalankan , maka i++
	angka := 1

	for i := 0; i < c; i++ {
		fmt.Print(angka, "\t", d, "\t")
		angka++

	}

}
