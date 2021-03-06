// Даны числа A , B , C (число A не равно 0). Рассмотрев дискриминант
// D = B^2 – 4· A · C , проверить истинность высказывания: «Квадратное уравне-
// ние A · x^2 + B · x + C = 0 имеет вещественные корни»
package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var a, b, c float64
	var x bool
	a = ioutil.NotNullNumber("число A, не равно 0")
	b = ioutil.Number("число B")
	c = ioutil.Number("число C")
	d := math.Pow(b, 2) - 4*a*c
	x = d > 0
	fmt.Println(x)
}
