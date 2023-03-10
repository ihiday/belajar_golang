package main

import "fmt"

func main() {
	// deklarasi kalimat
	str := "selamat malam"

	// deklarasi map kosong
	data := map[string]int{}

	for i := 0; i < len(str); i++ {
		// deklarasi untuk memudahkan memasukkan key
		karakter := string(str[i])

		// menampilkan karakter
		fmt.Println(karakter)
		if string(str[i]) == karakter {

			// melakukan input dan penambahan nilai jika karakter sama
			data[karakter] += 1
		}
	}
	fmt.Println(data)
}
