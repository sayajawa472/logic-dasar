package main

import "fmt"

func main() {
	Logic01Soal09()
}

func Logic01Soal09() {
	n := 12
	nilaiTetap := 1
	nilaiKali2 := 2
	nilaiKali3 := 3

	for i := 1; i <= n; i++ {
		if i%3 == 2 {
			fmt.Print(nilaiKali2, "\t")
			nilaiKali2 *= 2
		} else if i%3 == 0 {
			fmt.Print(nilaiKali3, "\t")
			nilaiKali3 *= 3
		} else {
			fmt.Print(nilaiTetap, "\t")
		}
	}
}
