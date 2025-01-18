// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func balapan(motor string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Motor %s mulai balapan...\n", motor)
// 	for lap := 1; lap <= 3; lap++ {
// 		time.Sleep(time.Second) // Simulasi setiap lap
// 		fmt.Printf("Motor %s selesai lap %d\n", motor, lap)
// 	}

// 	fmt.Printf("Motor %s selesai balapan!\n", motor)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	motorList := []string{"Motor A", "Motor B", "Motor C"}
// 	for _, motor := range motorList {
// 		wg.Add(1)
// 		go balapan(motor, &wg)
// 	}

// 	wg.Wait()

// 	fmt.Println("Balapan selesai!")
// }
