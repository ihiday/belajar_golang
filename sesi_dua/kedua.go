package main

import "fmt"

func main() {
	// perulangan nilai i
	for i := 0; i <= 4; i++ {
		fmt.Println("Nilai i = ", i)
	}

	// perulangan nilai j
	for j := 0; j <= 10; j++ {
		if j == 5 {

			// deklarasi teks untuk perulangan
			teks := []rune{'С', 'А', 'Ш', 'А', 'Р', 'В', 'О'}

			// perulangan karakter
			for indeks := 0; indeks < len(teks); indeks++ {
				char := teks[indeks]
				fmt.Printf("character %U '%s' starts at byte position %d \n", char, string(char), indeks*2)
			}
		} else {
			fmt.Println("Nilai j = ", j)
		}
	}

}
