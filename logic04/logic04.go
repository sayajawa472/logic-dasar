package logic04

import array2 "github.com/sayajawa472/logic-dasar/array"

func Logic04Soal01(n int) {
	array := array2.NewNumberArray(n, n) //Untuk membuat array

	a := 1

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if i%4 == 0 { //Membuat deret dari kiri ke kanan
				array[i][j] = int32(a)
				a++
			} else if i%4 == 1 && j == n-1 { //Membuat kotak kosong pada baris 1 dan berisi pada kolom 8
				array[i][j] = int32(a)
				a++
			} else if i%4 == 2 { //Membuat baris 3 dan baris 6 bergerak mundur
				array[i][n-1-j] = int32(a)
				a++
			} else if i%4 == 3 && j == 0 { //Membuat berisi pada kolom 0
				array[i][j] = int32(a)
				a++
			}
		}
	}
	array2.PrintNumberArrayZeroEmpty(array) //Print array
}

func Logic04Soal02(n int) {
	array := array2.NewNumberArray(n, n)
	a := 1

	for i := 0; i < len(array); i++ {

		for j := 0; j < len(array); j++ {
			if i%4 == 0 { //Membuat deret dari kiri ke kanan
				array[j][i] = int32(a)
				a += 3
			} else if i%4 == 1 && j == n-1 { //Membuat kotak kosong pada baris 1 dan berisi pada kolom 8
				array[j][i] = int32(a)
				a += 3
			} else if i%4 == 2 { //Membuat baris 3 dan baris 6 bergerak mundur
				array[n-1-j][i] = int32(a)
				a += 3
			} else if i%4 == 3 && j == 0 { //Membuat berisi pada kolom 0
				array[j][i] = int32(a)
				a += 3
			}
		}
	}
	array2.PrintNumberArrayZeroEmpty(array) //Print array
}

func Logic04Soal03(n int) {
	array := array2.NewNumberArray(n, n) //Membuat array

	a := 1

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if i%4 == 0 { //Membuat deret dari kiri ke kanan
				array[n-1-j][i] = int32(a)
				a += 3
			} else if i%4 == 1 && j == 0 { //Membuat berisi pada kolom 0
				array[j][i] = int32(a)
				a += 3
			} else if i%4 == 2 { //Membuat baris 3 dan baris 6 bergerak mundur
				array[j][i] = int32(a)
				a += 3
			} else if i%4 == 3 && j == n-1 { //Membuat kotak kosong pada baris 1 dan berisi pada kolom 8
				array[j][i] = int32(a)
				a += 3
			}
		}
	}
	array2.PrintNumberArrayZeroEmpty(array) //Print array
}

func Logic04Soal04(n int) {
	array := array2.NewNumberArray(n, n) //Membuat array

	a := 1

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if i%4 == 0 { //Membuat deret dari kiri ke kanan
				array[i][n-1-j] = int32(a)
				a += 3
			} else if i%4 == 1 && j == 0 { //Membuat berisi pada kolom 0
				array[i][j] = int32(a)
				a += 3
			} else if i%4 == 2 { //Membuat baris 3 dan baris 6 bergerak mundur
				array[i][j] = int32(a)
				a += 3
			} else if i%4 == 3 && j == n-1 { //Membuat kotak kosong pada baris 1 dan berisi pada kolom 8
				array[i][j] = int32(a)
				a += 3
			}
		}
	}
	array2.PrintNumberArrayZeroEmpty(array) //Print array
}
