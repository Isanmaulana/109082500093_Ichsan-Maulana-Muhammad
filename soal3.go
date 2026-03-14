package main

import "fmt"

func main() {
	var kg, gram, sisa, total, biayakirim, biayasisa int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&gram)
	kg = gram / 1000
	sisa = gram % 1000
	biayakirim = kg * 10000

	if sisa >= 500 {
		biayasisa = sisa * 5
	} else {
		biayasisa = sisa * 15
	}
	total = biayakirim + biayasisa
	if kg > 10 {
		total = biayakirim
	}
	fmt.Println("Detail berat:", kg, "kg +", sisa, "gr")
	fmt.Println("Detail biaya: Rp.", biayakirim, "+ Rp.", biayasisa)
	fmt.Println("Total biaya: Rp.", total)
}
