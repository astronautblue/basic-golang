package main

import "fmt"

func main() {
	// komentar kode
	// menampilkan pesan hello world
	// fmt.Println("hello world")
	// fmt.Println("baris ini tidak akan di eksekusi")

	// variable
	// var firstName string = "wick"
	// var lastName string
	// lastName = "john"

	// fmt.Println(firstName + " " + lastName)

	//	Skema penggunaan keyword var:
	//
	// var <nama-variabel> <tipe-data>
	// var <nama-variabel> <tipe-data> = <nilai>
	// Contoh:
	// var lastName string
	// var firstName string = "john"

	// var firstName string = "john"
	// lastName := "wick"
	// type data inference := tipe var otomatis membaca nilai nya
	// fmt.Printf("halo %s %s!\n", firstName, lastName)

	// menggunakan var, tanpa tipe data, menggunakan perantara "="
	// var firstName = "john"
	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	// lastName := "wick"

	//deklarasi multi variable
	// var first, second, third string
	// first, second, third = "satu", "dua", "tiga"

	// fmt.Println(first + second + third)

	// var fourth, fifth, sixth string = "empat", "lima", "enam"
	// seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	//variable inference multi tipe data
	// one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	// 	Golang memiliki aturan unik yang tidak dimiliki bahasa lain, yaitu tidak boleh ada satupun
	// variabel yang menganggur. Artinya, semua variabel yang dideklarasikan harus digunakan.
	// Jika ada variabel yang tidak digunakan tapi dideklarasikan, program akan gagal dikompilasi.

	name := new(string)
	fmt.Println(name)  // 0x20818a220
	fmt.Println(*name) // ""

}
