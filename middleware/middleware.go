package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}
func restar(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}
func dividir(a, b int) int {
	return a / b
}

func MiMiddleware() {
	fmt.Println("Inicio")
	resultadoS := operacionesMidd(sumar)(2, 3)
	fmt.Println(resultadoS)
	resultadoR := operacionesMidd(restar)(10, 6)
	fmt.Println(resultadoR)
	resultadoM := operacionesMidd(multiplicar)(2, 4)
	fmt.Println(resultadoM)
	resultadoD := operacionesMidd(dividir)(10, 2)
	fmt.Println(resultadoD)
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de Operación")
		return f(a, b)
	}
}
