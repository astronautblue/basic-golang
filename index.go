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

	// switch case/ ifelse style

	// var point = 6
	// switch {
	// case point == 8:
	// 	fmt.Println("perfect")
	// case (point < 8) && (point > 3):
	// 	fmt.Println("awesome")
	// default:
	// 	{
	// 		fmt.Println("not bad")
	// 		fmt.Println("you need to learn more")
	// 	}
	// }

	// 	Keyword fallthrough digunakan untuk memaksa proses pengecekkan diteruskan ke case
	// selanjutnya

	// var point = 6
	// switch {
	// case point == 8:
	// 	fmt.Println("perfect")
	// case (point < 8) && (point > 3):
	// 	fmt.Println("awesome")
	// 	fallthrough // next ke kondisi berikutnya
	// case point < 5:
	// 	fmt.Println("you need to learn more")
	// default:
	// 	{
	// 		fmt.Println("not bad")
	// 		fmt.Println("you need to learn more")
	// 	}
	// }

	// nested loop if/else beranak

	// var point = 10
	// if point > 7 {
	// 	switch point {
	// 	case 10:
	// 		fmt.Println("perfect!")
	// 	default:
	// 		fmt.Println("nice!")
	// 	}
	// } else {
	// 	if point == 5 {
	// 		fmt.Println("not bad")
	// 	} else if point == 3 {
	// 		fmt.Println("keep trying")
	// 	} else {
	// 		fmt.Println("you can do it")
	// 		if point == 0 {
	// 			fmt.Println("try harder!")
	// 		}
	// 	}
	// }

	// for i := 1; i <= 5; i++ {
	// 	for k := 1; k <= i; k++ {
	// 		fmt.Print(" ", i)
	// 	}
	// 	fmt.Println()
	// }

	// var i = 0
	// var k = 0
	// for i < 5 {
	// 	i++
	// 	for k < 5 {
	// 		k++
	// 		fmt.Print("* ")
	// 	}

	// 	fmt.Println()
	// }

	// var i = 0
	// for {
	// 	fmt.Println("Angka", i)
	// 	i++
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 1 { //jika ganjil lanjutkan
	// 		continue
	// 	}
	// 	if i > 8 { // jika i nya lebih besar dari 8 berhenti
	// 		break
	// 	}
	// 	fmt.Println("Angka", i)
	// }

	// for i := 0; i < 5; i++ {
	// 	for j := i; j < 5; j++ {
	// 		fmt.Print(j, " ")
	// 	}
	// 	fmt.Println()
	// }

	// outerLoop:
	// 	for i := 0; i < 5; i++ {
	// 		for j := 0; j < 5; j++ {
	// 			if i == 3 {
	// 				break outerLoop
	// 			}
	// 			fmt.Print("matriks [", i, "][", j, "]", "\n")
	// 		}
	// 	}

	// array

	// var names [4]string
	// names[0] = "trafalgar"
	// names[1] = "d"
	// names[2] = "water"
	// names[3] = "law"
	// fmt.Println(names[0], names[1], names[2], names[3])

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}
	// fmt.Println("Jumlah element \t\t", len(fruits))
	// fmt.Println("Isi semua element \t", fruits)

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}
	// fmt.Println("Jumlah element \t\t", len(fruits))
	// fmt.Println("Isi semua element \t", fruits)

	// var fruits [4]string

	// // cara horizontal
	// fruits = [4]string{"apple", "grape", "banana", "melon"}
	// // cara vertikal
	// fruits = [4]string{
	// 	"apple",
	// 	"grape",
	// 	"banana",
	// 	"melon",
	// }

	// var numbers = [...]int{2, 3, 2, 4, 3}
	// fmt.Println("data array \t:", numbers)
	// fmt.Println("jumlah elemen \t:", len(numbers))

	// var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	// var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	// fmt.Println("numbers1", numbers1)
	// fmt.Println("numbers2", numbers2)

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}
	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("elemen %d : %s\n", i, fruits[i])
	// }

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}
	// for i, fruit := range fruits {
	// 	fmt.Printf("elemen %d : %s\n", i, fruit)
	// }

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}
	// for _, fruit := range fruits {
	// 	fmt.Printf("nama buah : %s\n", fruit)
	// }

	// var fruits = make([]string, 4)
	// fruits[0] = "apple"
	// fruits[1] = "manggo"
	// fruits[2] = "2424242"
	// fmt.Println(fruits) // [apple manggo]

	// slice

	// var fruits = []string{"apple", "grape", "banana", "melon"}
	// fmt.Println(fruits[0])

	// var fruits = []string{"apple", "grape", "banana", "melon"}
	// var aFruits = fruits[0:3]
	// var bFruits = fruits[1:4]
	// var aaFruits = aFruits[1:2]
	// var baFruits = bFruits[0:1]
	// fmt.Println(fruits)   // [apple grape banana melon]
	// fmt.Println(aFruits)  // [apple grape banana]
	// fmt.Println(bFruits)  // [grape banana melon]
	// fmt.Println(aaFruits) // [grape]
	// fmt.Println(baFruits) // [grape]
	// // Buah "grape" diubah menjadi "pinnaple"
	// baFruits[0] = "pinnaple"
	// fmt.Println(fruits)   // [apple pinnaple banana melon]
	// fmt.Println(aFruits)  // [apple pinnaple banana]
	// fmt.Println(bFruits)  // [pinnaple banana melon]
	// fmt.Println(aaFruits) // [pinnaple]
	// fmt.Println(baFruits) // [pinnaple]

	// var fruits = []string{"apple", "grape", "banana", "melon"}
	// fmt.Println(len(fruits))

	// var fruits = []string{"apple", "grape", "banana", "melon"}
	// fmt.Println(len(fruits)) // len: 4
	// fmt.Println(cap(fruits)) // cap: 4
	// var aFruits = fruits[0:3]
	// fmt.Println(len(aFruits), aFruits) // len: 3
	// fmt.Println(cap(aFruits), aFruits) // cap: 4
	// var bFruits = fruits[1:4]
	// fmt.Println(len(bFruits)) // len: 3
	// fmt.Println(cap(bFruits)) // cap: 3

	// append: tambah index array baru
	// var fruits = []string{"apple", "grape", "banana"}
	// var cFruits = append(fruits, "papaya")
	// fmt.Println(fruits)  // ["apple", "grape", "banana"]
	// fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]

	// var fruits = []string{"apple", "grape", "banana"}
	// var bFruits = fruits[0:2]
	// fmt.Println(cap(bFruits)) // 3
	// fmt.Println(len(bFruits)) // 2
	// fmt.Println(fruits)       // ["apple", "grape", "banana"]
	// fmt.Println(bFruits)      // ["apple", "grape"]
	// var cFruits = append(bFruits, "papaya")
	// fmt.Println(fruits)  // ["apple", "grape", "papaya"]
	// fmt.Println(bFruits) // ["apple", "grape"]
	// fmt.Println(cFruits) // ["apple", "grape", "papaya"]

	//copy: mengambil jumlah elemen/index yang berhasil di gabungkan
	// var fruits = []string{"apple", "jambu"}
	// var aFruits = []string{"watermelon", "pinnaple"}
	// var copiedElemen = copy(fruits, aFruits)
	// fmt.Println(fruits)       // ["apple", "watermelon", "pinnaple"]
	// fmt.Println(aFruits)      // ["watermelon", "pinnaple"]
	// fmt.Println(copiedElemen) // 1

	// var fruits = []string{"apple", "grape", "banana"}
	// var aFruits = fruits[0:2]
	// var bFruits = fruits[0:2:2]
	// fmt.Println(fruits)       // ["apple", "grape", "banana"]
	// fmt.Println(len(fruits))  // len: 3
	// fmt.Println(cap(fruits))  // cap: 3
	// fmt.Println(aFruits)      // ["apple", "grape"]
	// fmt.Println(len(aFruits)) // len: 2
	// fmt.Println(cap(aFruits)) // cap: 3
	// fmt.Println(bFruits)      // ["apple", "grape"]
	// fmt.Println(len(bFruits)) // len: 2
	// fmt.Println(cap(bFruits)) // cap: 2

	// var chicken map[string]int
	// chicken = map[string]int{}
	// chicken["januari"] = 50
	// chicken["februari"] = 40
	// fmt.Println("januari", chicken["januari"]) // januari 50
	// fmt.Println("mei", chicken["mei"])         // mei 0

	// var chicken1 = map[string]int{"januari": 50, "februari": 40}
	// var chicken2 = map[string]int{
	// 	"januari":  50,
	// 	"februari": 40,
	// }

	// var chicken3 = map[string]int{}
	// var chicken4 = make(map[string]int)
	// var chicken5 = *new(map[string]int)

	// chicken4["januari"] = 50
	// chicken4["mei"] = 40

	// // chicken5["maret"] = 90
	// // chicken5["april"] = 80

	// fmt.Println(chicken4["januari"])
	// fmt.Println(chicken5)

	// var chicken = map[string]int{
	// 	"januari":  50,
	// 	"februari": 40,
	// 	"maret":    34,
	// 	"april":    67,
	// }
	// for key, val := range chicken {
	// 	fmt.Println(key, " \t:", val)
	// }

	// halaman 68
}
