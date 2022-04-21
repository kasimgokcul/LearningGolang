package dictionary

import (
	"fmt"
)

func TheWords() {

	var Sozluk = map[string]string{
		"book":     "kitap",
		"table":    "masa",
		"pencil":   "kalem",
		"school":   "okul",
		"able":     "yapabilen",
		"also":     "ayrica",
		"done":     "tamamlamak",
		"computer": "bilgisayar",
		"foot":     "ayak",
		"finger":   "parmak",
		"noise":    "burun",
	}
	var input string
	var f bool
	var w string
	fmt.Println("Cevirmek istediginiz kelimeyi girin")
	fmt.Scanln(&input)

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
