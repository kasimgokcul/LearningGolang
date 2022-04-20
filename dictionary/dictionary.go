package dictionary

import (
	"fmt"
)

const (
	book string = "kitap"
	kagit
	deftere
)

func TheWords() {

	var Sozluk = map[string]string{
		"book":   "kitap",
		"table":  "masa",
		"pencil": "kalem",
		"school": "okul",
	}
	var input string
	var f bool
	var w string
	fmt.Println("Cevirmek istediginiz kelimeyi girin")
	fmt.Scanln(&input)

	// if _, ok := Sozluk[input]; ok {
	// 	fmt.Println(Sozluk[input])
	// } else {
	// 	fmt.Println("Karsilik bulunamadi")
	// }

	for k, v := range Sozluk {
		if input == k {
			f = true
			w = v
			break
		} else if input == v {
			f = true
			w = k
			break
		}
	}
	if f {
		fmt.Println(w)
	} else {
		fmt.Println("Karsilik Bulunamadi")
	}

}
