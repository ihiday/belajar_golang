package main

import "fmt"

func main() {
	for i := 0; i <= 4; i++ {
		fmt.Println("Nilai dari i adalah ", i)
	}
	for j := 0; j <= 10; j++ {
		if j == 5 {
			teks := "CAWAPBO"
			for indeks := 0; indeks < len(teks); indeks++ {
				fmt.Printf("character %U '%s' starts at byte position %d \n", rune(teks[indeks]), string(rune(teks[indeks])), indeks*2)
			}
		} else {
			fmt.Println("Nilai dari j adalah ", j)
		}
	}

}
