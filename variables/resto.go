package variables

import (
	"fmt"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float64
var Fecha time.Time

func RestoVariables() {
	Nombre = "Juan"
	Estado = true
	Sueldo = 1500.50
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}
