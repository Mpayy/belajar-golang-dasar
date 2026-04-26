package main

import "fmt"

// Buat map untuk menyimpan nilai ujian beberapa siswa, lalu cari rata-ratanya
func rataRata(siswa map[string]int) float64{
	total:= 0
	for _ , value := range siswa {
		total += value
	}
	return float64(total) / float64(len(siswa))

}

func main(){
	siswa := map[string]int{
		"siswa1" : 95,
		"siswa2" : 80,
		"siswa3" : 70,
	}

	fmt.Println("Rata-rata nilai siswa", rataRata(siswa))
}