package interpreter

import (
	"evo-compiler/src/lexer"
	"strconv"
	"strings"
)

const (
	INTEGER = 0
	FLOAT   = 1
	STRING  = 2
)

func GetConstantValue(v ConstantValue) (value any, t int) {
	var returnType = v.constantType

	switch returnType {

	case lexer.INTEGER:
		lValue, ok := strconv.Atoi(v.value)

		if ok != nil {
			return nil, -1
		}

		return lValue, INTEGER

	case lexer.FLOAT:
		lValue, ok := strconv.ParseFloat(v.value, 64)

		if ok != nil {
			return nil, -1
		}

		return lValue, FLOAT

	case lexer.STRING:
		return v.value, STRING
	}

	return nil, -1
}

func GetTypes(left any, right any) (int, int) {
	var leftType = -1
	var rightType = -1

	_, isLeftInteger := left.(int)
	_, isRightInteger := right.(int)

	if isLeftInteger {
		leftType = INTEGER
	}

	if isRightInteger {
		rightType = INTEGER
	}

	_, isLeftFloat := left.(float64)
	_, isRightFloat := right.(float64)

	if isLeftFloat {
		leftType = FLOAT
	}

	if isRightFloat {
		rightType = FLOAT
	}

	_, isLeftString := left.(string)
	_, isRightString := right.(string)

	if isLeftString {
		leftType = STRING
	}

	if isRightString {
		rightType = STRING
	}

	return leftType, rightType
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

	leftType, rightType := GetTypes(leftValue, rightValue)

	if leftType == INTEGER && rightType == INTEGER {
		return leftValue.(int) + rightValue.(int)

	}

	if leftType == FLOAT && rightType == FLOAT {
		return leftValue.(float64) + rightValue.(float64)

	}

	if leftType == FLOAT && rightType == INTEGER {
		return leftValue.(float64) + float64(rightValue.(int))

	}

	if leftType == INTEGER && rightType == FLOAT {
		return float64(leftValue.(int)) + rightValue.(float64)

	}

	if leftType == STRING && rightType == STRING {
		return strings.Join([]string{leftValue.(string), rightValue.(string)}, "")
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

	leftType, rightType := GetTypes(leftValue, rightValue)

	if leftType == INTEGER && rightType == INTEGER {
		return leftValue.(int) - rightValue.(int)

	}

	if leftType == FLOAT && rightType == FLOAT {
		return leftValue.(float64) - rightValue.(float64)

	}

	if leftType == INTEGER && rightType == FLOAT {
		return float64(leftValue.(int)) - rightValue.(float64)

	}

	if leftType == FLOAT && rightType == INTEGER {
		return float64(rightValue.(int)) - leftValue.(float64)

	}

	return nil
}
