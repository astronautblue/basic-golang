package main

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

	// name := new(string)
	// fmt.Println(name)  // 0x20818a220
	// fmt.Println(*name) // ""

	// nama := "mm"

	// if nama == "mm" {
	// 	fmt.Println("Aditya")
	// } else {
	// 	fmt.Println("Saputra")
	// }

	// var point = 8
	// if point == 10 {
	// 	fmt.Println("lulus dengan nilai sempurna")
	// } else if point > 5 {
	// 	fmt.Println("lulus")
	// } else if point == 4 {
	// 	fmt.Println("hampir lulus")
	// } else {
	// 	fmt.Printf("tidak lulus. nilai anda %d\n", point)
	// }

	// var point = 7840.0
	// if percent := point / 100; percent >= 100 {
	// 	fmt.Printf("%.1f%s perfect!\n", percent, "%")
	// } else if percent >= 70 {
	// 	fmt.Printf("%.1f%s good\n", percent, "%")
	// } else {
	// 	fmt.Printf("%.1f%s not bad\n", percent, "%")
	// }

	// var point = 6
	// switch point {
	// case 8:
	// 	fmt.Println("perfect")
	// case 7, 6, 5, 4:
	// 	fmt.Println("awesome")
	// default:
	// 	fmt.Println("not bad")
	// }

}
