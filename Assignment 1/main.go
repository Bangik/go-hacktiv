package main

import (
	"fmt"
	"os"
)

type Teman struct {
	Nama         string
	Alamat       string
	Pekerjaan    string
	AlasanGolang string
}

func main() {
	// Mendapatkan argumen dari terminal
	absen := os.Args[1]

	// Konversi argumen ke integer
	absenInt := 0
	fmt.Sscanf(absen, "%d", &absenInt)

	// Data dummy teman
	teman := []Teman{
		{
			Nama:         "Budi",
			Alamat:       "Jl. Merdeka 12",
			Pekerjaan:    "Mahasiswa",
			AlasanGolang: "Ingin belajar bahasa pemrograman yang modern dan powerful",
		},
		{
			Nama:         "Ani",
			Alamat:       "Jl. Sudirman 23",
			Pekerjaan:    "Freelancer",
			AlasanGolang: "Ingin membuat aplikasi web yang scalable",
		},
		{
			Nama:         "Cici",
			Alamat:       "Jl. Diponegoro 34",
			Pekerjaan:    "Software Engineer",
			AlasanGolang: "Ingin meningkatkan skill programming",
		},
		{
			Nama:         "Doni",
			Alamat:       "Jl. Gajah Mada 45",
			Pekerjaan:    "Data Scientist",
			AlasanGolang: "Ingin belajar bahasa pemrograman yang powerful untuk analisis data",
		},
		{
			Nama:         "Euis",
			Alamat:       "Jl. Asia Afrika 56",
			Pekerjaan:    "UI/UX Designer",
			AlasanGolang: "Ingin membuat aplikasi web yang scalable",
		},
	}

	// Validasi absen
	if absenInt < 1 || absenInt > len(teman) {
		fmt.Println("Absen tidak ditemukan")
		return
	}

	// Mencari data teman berdasarkan absen
	var dataTeman Teman
	dataTeman = teman[absenInt-1]

	// Menampilkan data teman
	fmt.Println("Nama:", dataTeman.Nama)
	fmt.Println("Alamat:", dataTeman.Alamat)
	fmt.Println("Pekerjaan:", dataTeman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", dataTeman.AlasanGolang)
}
