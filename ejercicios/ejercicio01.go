package ejercicios

import (
	"strconv"
)

func ConvNumerico(texto string) (int, string) {

	var respuesta string
	num, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "Error: "
	}
	if num > 100 {
		respuesta = "Es mayor a 100"
	} else {
		respuesta = "Es menor a 100"
	}
	return num, respuesta
}
