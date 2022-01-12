package main

import "fmt"

//Cada arquivo pode ter um init e sempre vai começar por essa função
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
