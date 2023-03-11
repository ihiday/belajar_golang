package main

import "fmt"

func main() {

	// deklarasi kalimat
	str := "selamat malam"

	// deklarasi map kosong
	data := map[string]int{}

	for _, v := range str {

		// deklarasi untuk memudahkan memasukkan key
		karakter := string(v)

		// menampilkan karakter
		fmt.Print(karakter)

		// melakukan input dan penambahan nilai jika karakter sama
		data[karakter] += 1

	}
	fmt.Println(data)
}
