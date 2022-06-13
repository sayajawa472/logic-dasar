package logic03

import (
	"fmt"
	array2 "github.com/sayajawa472/logic-dasar/array"
)

func Logic03Soal01(n int) {
	nt := n / 2
	a := 3

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i >= j && i+j <= n-1 {
				fmt.Print(a, "\t")
			} else if i <= j && i+j >= n-1 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Print("\n")
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
	}
}

func Logic03Soal02(n int) {
	nt := n / 2
	a := 3

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i <= j && i+j <= n-1 {
				fmt.Print(a, "\t")
			} else if i >= j && i+j >= n-1 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Print("\n")
		if i < nt {
			a += 3
		} else {
			a -= 3
		}

	}
}

func Logic03Soal03(n int) {
	a := 3

	for i := 0; i < n; i++ {
		nt := n / 2
		for j := 0; j < n; j++ {
			if i < j && j <= nt {
				fmt.Print(a, "\t")
			} else if i > j && j >= nt {
				fmt.Print(a, "\t")
			} else if i+j < n-1 && i >= nt {
				fmt.Print(a, "\t")
			} else if i+j > n-1 && i <= nt {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		println("\n")
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
	}
}

func Logic03Soal03WithArray(n int) {
	array := array2.NewNumberArray(n, n) //line ini untuk membuat array

	//mengisi array
	a := 3 //Tipe variable int
	nt := n / 2
	for i := 0; i < len(array); i++ { // Looping baris
		for j := 0; j < len(array[i]); j++ { //Looping kolom
			if i < j && j <= nt {
				array[i][j] = int32(a) //Konvert variable a ke int32
			} else if i > j && j >= nt {
				array[i][j] = int32(a) //Konvert variable a ke int32
			} else if i+j < n-1 && i >= nt {
				array[i][j] = int32(a) //Konvert variable a ke int32
			} else if i+j > n-1 && i <= nt {
				array[i][j] = int32(a) //Konvert variable a ke int32
			}
		} //Looping kolom selesai
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
	} //Looping baris selesai
	array2.PrintNumberArrayZeroEmpty(array) //Print array
}

func Logic03Soal04(n int) {
	a := 3
	nt := n / 2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%4 == 0 {
				fmt.Print(a, "\t")
			} else if i%4 == 1 && j == n-1 {
				fmt.Print(a, "\t")
			} else if i%4 == 2 {
				fmt.Print(a, "\t")
			} else if i%4 == 3 && j == 0 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Println("\n")
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
	}
}

func Logic03Soal05(n int) {
	a := 3
	nt := n / 2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%4 == 0 {
				fmt.Print(a, "\t")
			} else if i%4 == 1 && j == 0 {
				fmt.Print(a, "\t")
			} else if i%4 == 2 {
				fmt.Print(a, "\t")
			} else if i%4 == 3 && j == n-1 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Println("\n")
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
	}
}
