package logic02

import "fmt"

func Logic02Soal01(n int) {
	f := 3

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(f, "\t")
		}
		fmt.Println("\n")
		f += 3
	}

}

func Logic02Soal02(n int) {
	for i := 0; i < n; i++ {
		f := 3
		for j := 0; j < n; j++ {
			fmt.Print(f, "\t")
			f += 3

		}
		fmt.Println("\n")
	}

}

func Logic02Soal03(n int) {
	for i := 0; i < n; i++ {
		f := 3 * n
		for j := 0; j < n; j++ {
			fmt.Print(f, "\t")
			f -= 3

		}
		fmt.Println("\n")

	}
}

func Logic02Soal04(n int) {
	f := n * 3
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(f, "\t")
		}
		fmt.Println("\n") // pindah baris
		f -= 3

	}
}

func Logic02Soal05(n int) {
	f := 3
	nt := n / 2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(f, "\t")
		}
		fmt.Println("\n")
		if i < nt {
			f += 3
		} else {
			f -= 3
		}

	}
}

func Logic02Soal06(n int) {

	nt := n / 2
	//looping baris
	for i := 0; i < n; i++ {
		a := 3
		//looping kolom
		for j := 0; j < n; j++ {
			if j < nt {
				fmt.Print(a, "\t")
				a += 3
			} else {
				fmt.Print(a, "\t")
				a -= 3
			}
		}
		//pindah baris
		fmt.Print("\n")
	}
}

func Logic02Soal07(n int) {
	//ini variabel a, nilainya 3
	a := 3
	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		for j := 0; j < n; j++ {
			if i < j {
				fmt.Print(" ", "\t")
			} else {
				//ini menampilkan kolom
				fmt.Print(a, "\t")
			}
		}
		//kolom selesai
		//kemudain pindah ke baris selanjutnya
		fmt.Print("\n")
		//saat pindah baris variabel a ditambah 3
		a += 3
	}
}

func Logic02Soal08(n int) {
	a := 3
	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		for j := 0; j < n; j++ {
			//membuat kondisi untuk kotak yg kosong
			if i > j {
				fmt.Print(" ", "\t")
			} else {
				fmt.Print(a, "\t")
			}
		}
		//kolom selesai
		//pindah kenbaris selanjutnya
		fmt.Print("\n")
		//nilai variabel a ditambah 3
		a += 3
	}
}

func Logic02Soal09(n int) {
	for i := 0; i < n; i++ {
		a := 0
		for j := 0; j < n; j++ {
			a += 3
			if i+j == n-1 {
				fmt.Print(a, "\t")
			} else if i+j <= n-1 {
				fmt.Print(" ", "\t")
			} else {
				fmt.Print(a, "\t")
			}
		}
		fmt.Print("\n")
	}
}

func Logic02Soal10(n int) {
	for i := 0; i < n; i++ {
		a := 0
		for j := 0; j < n; j++ {
			a += 3
			if i+j == n-1 {
				fmt.Print(a, "\t")
			} else if i+j <= n-1 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Print("\n")
	}
}
