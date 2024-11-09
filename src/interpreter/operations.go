package interpreter

import (
	"evo-compiler/src/lexer"
	"strconv"
)

func GetConstantValue(v ConstantValue) (value any, t int) {
	var returnType = v.constantType

	switch returnType {

	case lexer.INTEGER:
		lValue, ok := strconv.Atoi(v.value)

		if ok != nil {
			return nil, -1
		}

		return lValue, t

	case lexer.FLOAT:
		lValue, ok := strconv.ParseFloat(v.value, 64)

		if ok != nil {
			return nil, -1
		}

		return lValue, t
	}

	return nil, -1
}

func SumOf(left any, right any) any {
	var leftValue any = left
	var rightValue any = right

	if leftConstant, ok := left.(ConstantValue); ok {
		leftValue, _ = GetConstantValue(leftConstant)
	}

	if rightConstant, ok := right.(ConstantValue); ok {
		rightValue, _ = GetConstantValue(rightConstant)
	}

	// BOTH INTEGERS

	leftInteger, isLeftInteger := leftValue.(int)
	rightInteger, isRightInteger := rightValue.(int)

	if isLeftInteger && isRightInteger {
		return leftInteger + rightInteger
	}

	// INTEGER AND FLOAT

	rightFloat, isRightFloat := rightValue.(float64)
	if isLeftInteger && isRightFloat {
		return float64(leftInteger) + rightFloat
	}

	leftFloat, isLeftFloat := leftValue.(float64)

	if isLeftFloat && isRightInteger {
		return leftFloat + float64(rightInteger)
	}

	// BOTH FLOATS

	if isLeftFloat && isRightFloat {
		return leftFloat + rightFloat
	}

	return nil
}

func SubOf(left any, right any) any {
	var leftValue any = left
	var rightValue any = right

	if leftConstant, ok := left.(ConstantValue); ok {
		leftValue, _ = GetConstantValue(leftConstant)
	}

	if rightConstant, ok := right.(ConstantValue); ok {
		rightValue, _ = GetConstantValue(rightConstant)
	}

	// BOTH INTEGERS

	leftInteger, isLeftInteger := leftValue.(int)
	rightInteger, isRightInteger := rightValue.(int)

	if isLeftInteger && isRightInteger {
		return leftInteger - rightInteger
	}

	// INTEGER AND FLOAT

	rightFloat, isRightFloat := rightValue.(float64)
	if isLeftInteger && isRightFloat {
		return float64(leftInteger) - rightFloat
	}

	leftFloat, isLeftFloat := leftValue.(float64)

	if isLeftFloat && isRightInteger {
		return leftFloat - float64(rightInteger)
	}

	// BOTH FLOATS

	if isLeftFloat && isRightFloat {
		return leftFloat - rightFloat
	}

	return nil
}
