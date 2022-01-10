package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	// versão string de x
	xs := fmt.Sprint(x)

	fmt.Printf("O valor de x é " + xs)

	// sem precisar tranformar x em string
	fmt.Println("O valor de x é ", x)

	fmt.Printf("O valor de x é %f", x)

	//  arredondando e deixando em 2 casas decimais
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	// usando o v para mostrar o valor
	fmt.Printf("\n%v %v %.1v %v %v", a, b, b, c, d)

}
