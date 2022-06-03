package main

import "fmt"

func PrintNumberArray(array [][]int32) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Print(array[i][j], "\t")
		}
		fmt.Println("\n")
	}

}
