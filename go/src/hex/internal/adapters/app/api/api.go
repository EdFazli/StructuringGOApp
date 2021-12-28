//

package api

import (
	"github.com/EdFazli/StructuringGOApp/internal/ports"
)

type Adapter struct {
	db    ports.DbPort
	arith ports.ArithmeticPort
}

func NewAdapter(db ports.DbPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{db: db, arith: arith}
}

func (apiA Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apiA.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = apiA.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (apiA Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := apiA.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = apiA.db.AddToHistory(answer, "subtraction")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (apiA Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apiA.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	err = apiA.db.AddToHistory(answer, "multiplication")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (apiA Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apiA.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = apiA.db.AddToHistory(answer, "division")
	if err != nil {
		return 0, err
	}

	return answer, nil
}
