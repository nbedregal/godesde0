package main

import (
	e "github.com/nbedregal/godesde0/ejer_interfaces"
	m "github.com/nbedregal/godesde0/modelos"
)

func main() {
	Pedro := new(m.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(m.Mujer)
	e.HumanosRespirando(Maria)
}
