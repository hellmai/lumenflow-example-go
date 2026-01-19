// LumenFlow Example Go Project
package main

import (
	"fmt"

	"github.com/hellmai/lumenflow-example-go/calculator"
)

func main() {
	fmt.Println("LumenFlow Example Go Project")
	fmt.Printf("2 + 3 = %.0f\n", calculator.Add(2, 3))
	fmt.Printf("5 - 2 = %.0f\n", calculator.Subtract(5, 2))
	fmt.Printf("4 * 3 = %.0f\n", calculator.Multiply(4, 3))

	result, err := calculator.Divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.0f\n", result)
	}
}
