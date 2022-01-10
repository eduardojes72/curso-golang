package main

import "fmt"

// O desafio é refatorar as seguintes condições e colocar como switch

// if n >= 9 && n <= 10 {
// 	return "A"
// } else if n >= 8 {
// 	return "B"
// } else if n >= 5 {
// 	return "C"
// } else if n >= 3 {
// 	return "D"
// } else {
// 	return "E"
// }

func notaParaConceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
