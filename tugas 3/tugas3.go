package main

import "fmt"

func main() {
	var buah = []string{"apel", "jeruk", "anggur", "melon"}
	buah = append(buah, "pepaya")
	fmt.Println("jumlah element :", len(buah))
	fmt.Println("isi element : ", buah)
	for i := 0; i < len(buah); i++ {
		fmt.Printf("element ke - : %d %s \n", i, buah[i])
	}

}
