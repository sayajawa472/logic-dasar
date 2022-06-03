package main

import "fmt"

func main() {
	Logic01Soal08()

}

func Logic01Soal08() {
	n := 10
	nilaiKali2 := 2
	nilaiKali5 := 5

	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Print(nilaiKali2, "\t")
			nilaiKali2 *= 2
		} else {
			fmt.Print(nilaiKali5, "\t")
			nilaiKali5 *= 5
		}

	}
}
