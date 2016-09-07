// Дни недели пронумерованы следующим образом: 1 — понедельник,
// 2 — вторник, ..., 6 — суббота, 7 — воскресенье. Дано целое число K , ле-
// жащее в диапазоне 1–365. Определить номер дня недели для K -го дня года,
// если известно, что в этом году 1 января было субботой
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x uint
	for x == 0 || x > 365 || x < 1 {
		x = util.UInteger("введите K, лежащее в диапазоне 1–365")
	}
	x = x % 7
	switch {
	case x == 0:
		x = 5
		fmt.Printf("номер дня недели для K-го дня года: %v\n", x)
	case x == 1:
		x = 6
		fmt.Printf("номер дня недели для K-го дня года: %v\n", x)
	case x == 2:
		x = 7
		fmt.Printf("номер дня недели для K-го дня года: %v\n", x)
	case x == 3:
		x = 1
		fmt.Printf("номер дня недели для K-го дня года: %v\n", x)
	case x == 4:
		x = 2
		fmt.Printf("номер дня недели для K-го дня года: %v\n", x)
	case x == 5:
		x = 3
		fmt.Printf("номер дня недели для K-го дня года: %v\n", x)
	case x == 6:
		x = 4
		fmt.Printf("номер дня недели для K-го дня года: %v\n", x)
	}
}