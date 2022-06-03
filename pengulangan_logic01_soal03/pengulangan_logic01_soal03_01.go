package pengulangan_logic01_soal03

import "fmt"

func main() {
	Logic01Soal03Ke01()
}
func Logic01Soal03Ke01() {
	//membuat variable a nilainya 10
	a := 10
	//membuat variable b nilainya 99
	b := 99
	// membuat deret angka menggunakan for
	// for i := 0 artinya perulangan mulainya dari variable i=0
	// i kurang dr a artinya perulangan dihentikan ketikan nilai i kurang dr a
	// i++ artinya ketika semua pernyataan dalam for selesai dijalankan , maka i++
	angka := 1

	for i := 0; i < a; i++ {
		fmt.Print(angka, "\t", b, "\t")
		angka++

	}

}
