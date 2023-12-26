package main

import (
	"fmt"
	"math"
)

// funcion para hallar el area de un ciruclo
func cirCulo(radio float64) float64 {
	area := math.Pi * math.Pow(radio, 2)
	return area

}
func main() {
	areaCirculo := cirCulo(5)
	fmt.Println(" el area  es", areaCirculo)

}
