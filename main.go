package main

import (
	d "github.com/nbedregal/godesde0/middleware"
)

func main() {
	/*canal1 := make(chan bool)
	go d.MiNombreLento("Nico Bedregal", canal1)
	defer func() {
		<-canal1
	}()
	fmt.Println("Estoy acá")*/

	d.MiMiddleware()

}
