package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("diegofrleal@gmail.com")
	fmt.Println(erro)
}

/*
Para lembrar:
- ./modulo "o ./ referencia a pasta que vc está"
- go build
- go install (modulo vai para raiz onde o go foi instalado)
- go mod tidy (retira todas as dependencias não utilizadas)
*/
