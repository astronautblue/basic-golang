// package main

// import (
// 	"fmt"
// )

// func main() {
// 	runApplication(true)
// 	fmt.Println("Jalankan setelah terjadi panic")
// }

// func runApplication(error bool) {
// 	// call defer func
// 	defer endApp()

// 	if error {
// 		panic("Upss Error")
// 	}
// }

// func endApp() {
// 	fmt.Println("END APPS")
// 	message := recover()
// 	fmt.Println("Terjadi error!", message)
// }
