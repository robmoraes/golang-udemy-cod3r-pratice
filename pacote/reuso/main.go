package main

import (
	"fmt"

	"github.com/robmoraes/golang-udemy-cod3r-pkg/abc"
	"github.com/robmoraes/golang-udemy-cod3r-pkg/area"
)

func main() {
	fmt.Println("Practice", "pacote", "reuso", "main")
	fmt.Println(area.Circ(6.0))
	abc.Epa()
	fmt.Println(area.Rect(5.0, 2.0))
}
