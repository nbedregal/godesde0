package main

import (
	"fmt"

	"github.com/nbedregal/godesde0/ejercicios"
)

func main() {
	resp, num := ejercicios.ConvNumerico("130")

	fmt.Println("Respuesta: ", resp, " Valor: ", num)
}
