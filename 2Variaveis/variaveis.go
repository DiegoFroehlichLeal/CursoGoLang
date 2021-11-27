package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2" //inferencia de tipo
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "aaaaaa"
		variavel4 string = "bbbbbb"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
