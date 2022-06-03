package main

func diagonalDifference(arr [][]int32) int32 {
	var result int32
	var kanan int32 = 0
	var kiri int32 = 0

	// looping baris
	for i := 0; i < len(arr); i++ {
		// looping kolom
		for j := 0; j < len(arr[i]); j++ {
			if i == j {
				kanan = kanan + arr[i][j]
			}
			if i+j == len(arr)-1 {
				kiri = kiri + arr[i][j]
			}
		}
	}
	if kanan > kiri {
		result = kanan - kiri
	} else {
		result = kiri - kanan
	}
	return result
}
