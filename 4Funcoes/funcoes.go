package main

import "fmt"

func somar(n1 int32, n2 int32) int32 {
	return n1 + n2
}

//função com mais de um retorno
func calculosMatematicos(n1, n2 int32) (int32, int32) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao

}
func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	//declara uma variavel do tipo função e joga uma função nela.
	var f = func() {
		fmt.Println("Função F")
	}
	f()

	//resultado da função com mais de um retorno
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15) //se quiser usar somente um basta substituir por _ exemplo: resultadoSoma, _ := calculosMatematicos
	fmt.Println(resultadoSoma, resultadoSubtracao)
}
