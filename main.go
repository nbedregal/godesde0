package main

import (
	"fmt"

	d "github.com/nbedregal/godesde0/goroutines"
)

func main() {
	canal1 := make(chan bool)
	go d.MiNombreLento("Nico Bedregal", canal1)
	defer func() {
		<-canal1
	}()
	fmt.Println("Estoy acá")

}
