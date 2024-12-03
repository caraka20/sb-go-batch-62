package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// SOAL 1
	// Mendefinisikan 5 variabel untuk setiap kata
	word1 := "Bootcamp"
	word2 := "Digital"
	word3 := "Skill"
	word4 := "Sanbercode"
	word5 := "Golang"

	// Menggabungkan variabel menjadi satu kalimat
	sentence := word1 + " " + word2 + " " + word3 + " " + word4 + " " + word5

	// Menampilkan kalimat
	fmt.Println(sentence)

	// SOAL 2
	halo := "Halo Dunia"

	// Mengganti kata "Dunia" menjadi "Golang"
	halo = strings.Replace(halo, "Dunia", "Golang", 1)

	// Menampilkan hasil
	fmt.Println(halo)

	// SOAL 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	// Format kata sesuai dengan yang diminta
	kataKeduaFormatted := strings.Title(kataKedua)                  // "Senang"
	kataKetigaFormatted := strings.Replace(kataKetiga, "r", "R", 1) // "belajaR"
	kataKeempatFormatted := strings.ToUpper(kataKeempat)            // "GOLANG"

	// Gabungkan kata
	output := fmt.Sprintf("%s %s %s %s", kataPertama, kataKeduaFormatted, kataKetigaFormatted, kataKeempatFormatted)

	// Cetak hasil
	fmt.Println(output)

	// SOAL 5
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	// Ubah string menjadi integer
	angka1, _ := strconv.Atoi(angkaPertama)
	angka2, _ := strconv.Atoi(angkaKedua)
	angka3, _ := strconv.Atoi(angkaKetiga)
	angka4, _ := strconv.Atoi(angkaKeempat)

	// Jumlahkan semua angka
	total := angka1 + angka2 + angka3 + angka4

	// Cetak hasil
	fmt.Println("Jumlah total:", total)

	// SOAL 6
	kalimat := "halo halo bandung"
	angka := 2021

	// Ganti kata "halo" menjadi "Hi"
	kalimatFormatted := strings.Replace(kalimat, "halo", "Hi", -1)

	// Gabungkan dengan angka
	output = fmt.Sprintf("\"%s\" - %d", kalimatFormatted, angka)

	// Cetak hasil
	fmt.Println(output)
}
