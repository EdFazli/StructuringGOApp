//resposible to orchestrating the startup of application, dependency injection and etc

package main

import (
	"fmt"

	"github.com/EdFazli/StructuringGOApp/internal/adapters/core/arithmetic"
)

func main() {
	fmt.Println("Main file")

	arithAdapter := arithmetic.NewAdapter()
	result, err := arithAdapter.Addition(3, 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
