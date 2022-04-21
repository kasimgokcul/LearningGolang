package numberguessing

import "fmt"

func Guessing() {

	sayi, tahmin, tahminsayisi := 123, 0, 0

	for tahmin != sayi {
		fmt.Println("Tahmininizi giriniz")
		fmt.Scanln(&tahmin)
		tahminsayisi = tahminsayisi + 1
		if tahmin > sayi {
			fmt.Println("Daha kucuk bir sayi deneyiniz")
		} else if tahmin < sayi {
			fmt.Println("Daha buyuk bir sayi deneyiniz")
		} else if tahminsayisi <= 3 {
			fmt.Printf("Mukemmel %d tahminde bildiniz\n", tahminsayisi)
		} else if tahminsayisi > 3 && tahminsayisi <= 6 {
			fmt.Printf("Guzel %d tahminde bildiniz\n", tahminsayisi)
		} else {
			fmt.Printf("Fena degil %d tahminde bildiniz\n", tahminsayisi)
		}
	}

}
