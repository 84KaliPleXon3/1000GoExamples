// Дано целое число N (> 0). Найти наименьшее целое положительное
// число K , квадрат которого превосходит N : K^2 > N . Функцию извлечения
// квадратного корня не использовать

package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		n int
		x int
	)

	for n <= 0 {
		n = ioutil.Integer("N")
	}
	x = countLoop(n)
	fmt.Printf("K^2 > N, K = %v\n", x)

}

func countLoop(n int) int {
	x := 1
	for math.Pow(float64(x), 2) <= float64(n) {
		x++
	}
	return x
}
