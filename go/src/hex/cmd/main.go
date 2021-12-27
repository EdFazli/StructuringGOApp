//resposible to orchestrating the startup of application, dependency injection and etc

package main

import (
	"fmt"

	"github.com/EdFazli/StructuringGOApp/internal/adapters/core/arithmetic"
)

func main() {
	fmt.Println("Main file")

	arithAdapter := arithmetic.NewAdapter()
	arithAdapter.Addition(3, 5)
}
