package main

// Developed By Dwi Cipta Nugraha
// Adhi Tama Institute Technology Surabaya

import (
	"fmt"
	"math"
	"strconv"
)

// Fungsi Luas Permukaan Tabung (Challenge 01)
func luas_tabung() {
	var tinggi float64
	var phi float64 = 1.61803398874989484820458683436563811772030917980576286213544862
	var jari2 float64
	var lp float64
	fmt.Print("Tinggi = ")
	fmt.Scanln(&tinggi)
	fmt.Print("Jari-jari = ")
	fmt.Scanln(&jari2)
	lp = (2 * (phi * math.Pow(jari2, 2))) + (2 * phi * jari2 * tinggi)
	notif := "Luas Permukaan Tabung = " + strconv.FormatFloat(lp, 'E', -1, 64)
	fmt.Println(notif)
}

// Fungsi Grade (Challenge 02)
func grade() {
	var grade int
	var huruf string
	fmt.Print("Nilai = ")
	fmt.Scanln(&grade)
	if grade > 85 {
		huruf = "A"
	} else if grade > 75 {
		huruf = "B"
	} else if grade > 65 {
		huruf = "C"
	} else if grade > 55 {
		huruf = "D"
	} else if grade < 55 {
		huruf = "TIDAK LULUS"
	}
	notif := "Nilai " + strconv.Itoa(grade) + " Adalah " + huruf
	fmt.Println(notif)
}

// Fungsi Jumlah (Challenge 03)
func jumlah() {
	var jumlah_angka int
	var i int
	var total = 0
	fmt.Print("Input Jumlah Angka = ")
	fmt.Scanln(&jumlah_angka)
	for i = 0; i <= jumlah_angka; i++ {
		total += i
	}
	fmt.Print("Total = ")
	fmt.Println(total)
}

// Fungsi Segitiga Bintang (Challenge 04)
func segitiga() {
	var i int
	var j int
	var angka int
	fmt.Print("Input Jumlah Angka = ")
	fmt.Scanln(&angka)
	for i = 0; i <= angka-1; i++ {
		for j = 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var pilih int
	notif := "Silahkan Pilih Menu : \n"
	notif += "1. Challenge 01\n"
	notif += "2. Challenge 02\n"
	notif += "3. Challenge 03\n"
	notif += "4. Challenge 04\n"
	fmt.Println(notif)
	fmt.Print("Pilih Menu = ")
	fmt.Scanln(&pilih)
	switch pilih {
	case 1:
		luas_tabung()
	case 2:
		grade()
	case 3:
		jumlah()
	case 4:
		segitiga()
	default:
		fmt.Println("Menu Hanya 1-4")
	}
}
