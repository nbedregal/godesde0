package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func TablaNumerica() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese número: ")

	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			TablaNumerica()
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(numero, "*", i, " : ", i*numero)
	}

}
