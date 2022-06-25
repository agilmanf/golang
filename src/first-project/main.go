package main

import "fmt"

func main() {
	simpleCalculator()

}

func simpleCalculator() {
	var pilihan uint

	fmt.Printf("SIMPLE CALCULATOR \n----------------------\n")

	fmt.Println("1. Tambah")
	fmt.Println("2. Kurang")
	fmt.Println("3. Kali")
	fmt.Printf("4. Bagi\n\n")

	fmt.Printf("Pilih : ")
	fmt.Scan(&pilihan)

	if pilihan <= 4 {
		var angka1 float64
		var angka2 float64

		fmt.Printf("Masukan Angka ke-1 : ")
		fmt.Scan(&angka1)

		fmt.Printf("Masukan Angka ke-2 : ")
		fmt.Scan(&angka2)

		switch pilihan {
		case 1:
			pertambahan(angka1, angka2)
			break
		case 2:
			pengurangan(angka1, angka2)
			break
		case 3:
			perkalian(angka1, angka2)
			break
		case 4:
			pembagian(angka1, angka2)
			break
		}
	} else {
		fmt.Printf("Invalid input, masukan angka 1-4!\n\n")
		simpleCalculator()
	}
}

func pertambahan(a float64, b float64) {
	total := a + b
	fmt.Printf("Hasil Pertambahan %v + %v = %v", a, b, total)
}

func pengurangan(a float64, b float64) {
	total := a - b
	fmt.Printf("Hasil Pengurangan %v - %v = %v", a, b, total)
}

func perkalian(a float64, b float64) {
	total := a * b
	fmt.Printf("Hasil Perkalian %v x %v = %v", a, b, total)
}

func pembagian(a float64, b float64) {
	total := a / b
	fmt.Printf("Hasil Pembagian %v : %v = %v", a, b, total)
}
