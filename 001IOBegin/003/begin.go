// Даны стороны прямоугольника a и b. Найти его площадь S = a·b и периметр P = 2·(a + b)
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y float64
	fmt.Println("Необходимо найти площадь и периметр прямоугольника")
	x = ioutil.Number("Введите значение стороны 1")
	y = ioutil.Number("Введите значение стороны 2")
	fmt.Println("Площадь равна:\t", x*y)
	fmt.Println("Периметр равен:\t", 2*(x+y))
}
