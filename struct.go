package main

import "fmt"

type Customers struct {
	Name, Addres, NoHP string
	Age                int
}

func (customer Customers) sayHello(name string) {
	fmt.Println("Hallo saya " + name + " adalah " + customer.Name)
}

func main() {
	// var customer Customers

	// customer.Name = "Aditya Saputra"
	// customer.Addres = "Kota Bekasi"
	// customer.NoHP = "0892828828"
	// customer.Age = 30

	customer := Customers{
		Name:   "Aditya",
		Addres: "Bekasi",
		NoHP:   "089292828",
		Age:    20,
	}

	customer1 := Customers{"Aditya", "Bogor", "242424", 30}

	fmt.Println("Customers : ", customer.Name)
	fmt.Println("Customers : ", customer1.Addres)
	fmt.Println("Customers : ", customer.NoHP)
	fmt.Println("Customers : ", customer.Age)

	customer.sayHello("dodit")
}
