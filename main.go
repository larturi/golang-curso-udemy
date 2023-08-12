package main

import (
	"fmt"

	"github.com/larturi/golang-curso-udemy/variables"
)

func main() {
	// variables.RestoVariables()
	estado, texto := variables.ConvertIntToString(12345)
	fmt.Println(estado)
	fmt.Println(texto)
}
