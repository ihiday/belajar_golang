package main

import "fmt"

func main() {
	// deklarasi kalimat
	str := "selamat malam"

	// deklarasi map kosong
	data := map[string]int{}

	for i := 0; i < len(str); i++ {
		karakter := string(str[i])
		fmt.Println(karakter)
		if string(str[i]) == karakter {
			data[karakter] += 1
		}
	}
	fmt.Println(data)
}
