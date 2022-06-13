package logic01

import "fmt"

func logic01Soal10() {
	n := 12
	nilaiTambah := 1
	nilaiTambah2 := 2
	nilaiTambah3 := 3

	for i := 1; i <= n; i++ {
		if i%4 == 1 {
			fmt.Print(nilaiTambah3, " | ")
			nilaiTambah3 += 3
		} else if i%4 == 2 {
			fmt.Print(nilaiTambah2, " | ")
			nilaiTambah2 += 2
		} else if i%4 == 3 {
			fmt.Print(nilaiTambah, " | ")
			nilaiTambah += 1
		} else {
			fmt.Print(999, " | ")
		}
	}
}
