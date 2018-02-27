package main

import "fmt"

func main(){
	var a [5] int
	fmt.Println(a)

	b := [5] int{1,3,4,5,2}
	fmt.Println(b)

	fmt.Println(b[2])
	fmt.Println(b[1])

	fmt.Println("lenght = ",len(b))
	var dd [3][3] int
	fmt.Println(dd)

	for i:= 0 ; i < 3 ; i++{
		for j := 0; j < 3; j++{
			dd[i][j]=i+j
		}
	}

	fmt.Println("filled tow d : ",dd)


}	