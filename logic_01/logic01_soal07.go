package main

import "fmt"

func main() {
	Logic01Soal07()
}

func Logic01Soal07() {
	n := 10
	angkatetap := 4
	angkakelipatan3 := 3

	// i kurang dr  = n artinya perulangan berlanjut selama nilai i kurang dr  & = n
	//dan dihentikan ketika nilai i lnh dr nilai n
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Print(angkatetap, "\t")
		} else {
			fmt.Print(angkakelipatan3, "\t")
			angkakelipatan3 += 3
		}
	}

}
