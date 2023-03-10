package main

import "fmt"

func main() {
	// deklarasi kalimat
	str := "selamat malam"

	// deklarasi map kosong
	data := map[string]int{}

	for _, v := range str {
		fmt.Println(string(v))
	}

	for i := 0; i < len(str); i++ {
		karakter := string(str[i])
		if string(str[i]) == karakter {
			data[karakter] += 1
		} else {
			data[karakter] = 0
		}
	}
	fmt.Println(data)
}
