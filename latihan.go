package main

import "fmt"

func main() {
	var pil string
	var no1 float32
	var no2 float32
	
	fmt.Print("pilihan operator (+,-,/,*):")
	fmt.Scanln(&pil)


	// if pil == "+" {
	// 	fmt.Print("masukkan number 1 :")
	// 	fmt.Scanln(&no1)
	// 	fmt.Print("masukkan number 2 :")
	// 	fmt.Scanln(&no2)
	// 	fmt.Println("hasil pertambahan", no1+no2)
	// } else if pil == "-" {
	// 	fmt.Print("masukkan number 1 :")
	// 	fmt.Scanln(&no1)
	// 	fmt.Print("masukkan number 2 :")
	// 	fmt.Scanln(&no2)
	// 	fmt.Println("hasil pengurangan", no1-no2)
	// } else if pil == "/" {
	// 	fmt.Print("masukkan number 1 :")
	// 	fmt.Scanln(&no1)
	// 	fmt.Print("masukkan number 2 :")
	// 	fmt.Scanln(&no2)
	// 	fmt.Println("hasil bagi", no1/no2)
	// } else if pil == "*" {
	// 	fmt.Print("masukkan number 1 :")
	// 	fmt.Scanln(&no1)
	// 	fmt.Print("masukkan number 2 :")
	// 	fmt.Scanln(&no2)
	// 	fmt.Println("hasil perkalian", no1*no2)
	// } else {
	// 	fmt.Println("pilihan operator salah")
	// }
	if pil == "+" || pil =="-" || pil =="/" || pil =="*"{
		fmt.Print("masukkan number 1 :")
		fmt.Scanln(&no1)
		fmt.Print("masukkan number 2 :")
		fmt.Scanln(&no2)
		if pil == "+" {
			fmt.Println("hasil penambahan",no1+no2 )
		} else if pil == "-" {
			fmt.Println("hasil pengurangan",no1-no2 )
		} else if pil == "/" {
			fmt.Println("hasil bagi",no1/no2 )
		}else {
			fmt.Println("hasil perkalian",no1*no2 )
		}
		
	} else {
		fmt.Println("pilihan operator salah")
	}
}