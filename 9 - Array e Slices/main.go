package main

import "fmt"

func main() {
	var arr1 [5]string
	arr1[0] = "Posicao 1"
	fmt.Println(arr1)

	arr2 :=[5]int{1,2,3,4,5}
	fmt.Println(arr2)

	arr3 := [...]string{"1","1","1","1","1"}
	fmt.Println(arr3)

	slice := []int{1,2,3,4,5,6,7,8}
	fmt.Println(slice)

	slice = append(slice, 9)
	fmt.Println(slice)

	slice2 := arr2[1:3]
	fmt.Println(slice2)

	arr2[1] = 10
	fmt.Println(slice2)

	// Array internos
	slice3:=make([]float32, 10, 15)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 10)
	fmt.Println(slice3)

	fmt.Println(len(slice)) // length
	fmt.Println(cap(slice)) // capacidade



}