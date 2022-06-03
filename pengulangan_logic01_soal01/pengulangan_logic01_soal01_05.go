package pengulangan_logic01_soal01

import (
	"fmt"
	"testing"
)

func TestLogic01Soal01Ke05(t testing.T) {
	//looping artiny pengulangan
	e := 10
	// berarti + =1
	angka := 1
	// for artiny looping , balik ke program sampe ke -e kali
	// for i := 0 artiny looping dimulai dr 0
	for i := 0; i < e; i++ {
		// if i%2 == 0 artiny  jk i dibagi 2 hasilny 0 maka true
		//if true masuk ke kurung kurawal pertama if
		//if true  fmt.Println(angka, "\t") angka ++
		//if false yg di dalem kurang skip
		//if i%2 == 0 habis dibagi 2 (genap) maka print angka
		//modulo 2 hasilny 0 maka genap
		if i%2 == 0 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(3, "\t")
		}
		//jk index genap variable "angka" +1
	}
}
