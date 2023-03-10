package main

import (
	fmt "fmt"
	"strconv"
)

func main() {
	i := 21
	j := true
	heksa, _ := strconv.Atoi(fmt.Sprintf("%x", i))
	var k float64 = 123.456
	//menampilkan nilai i
	fmt.Printf("%v \n", i)

	//menampilkan tipe data i
	fmt.Printf("%T \n", i)

	//menampilkan isi boolean j
	fmt.Printf("%t \n", j)

	//menampilkan simbol persen %
	fmt.Printf("%% \n")

	//menampilkan karakter ya
	fmt.Println("\u042f")

	//menampilkan desimal i
	fmt.Printf("%d \n", i)

	//menampilkan octal i
	fmt.Printf("%o \n", i)

	//menampilkan biner i
	fmt.Printf("%b \n", i)

	//menampilkan heksa huruf kecil
	fmt.Printf("%x \n", heksa)

	//menampilkan heksa dengan huruf kapital
	fmt.Printf("%X \n", heksa)

	//menampilkan unicode dari ya
	fmt.Printf("%U \n", 'Я')

	//menampilkan float k
	fmt.Printf("%f \n", k)

	//menampilkan scientific k
	fmt.Printf("%e \n", k)

}
