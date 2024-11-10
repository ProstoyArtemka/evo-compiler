package interpreter

import (
	"evo-compiler/src/lexer"
	"fmt"
	"strconv"
	"strings"
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

	case lexer.FALSE:
		return false, BOOLEAN

	case lexer.TRUE:
		return true, BOOLEAN
	}

	return nil, -1
}

func GetConstantValues(left any, right any) (any, any) {
	if leftConstant, ok := left.(ConstantValue); ok {
		left, _ = GetConstantValue(leftConstant)
	}

	if rightConstant, ok := right.(ConstantValue); ok {
		right, _ = GetConstantValue(rightConstant)
	}

	return left, right
}

func GetTypes(left any, right any) (int, int) {
	return GetType(left), GetType(right)
}

func SumOf(left any, right any) any {
	left, right = GetConstantValues(left, right)
	leftType, rightType := GetTypes(left, right)

	// INTEGERS

	if leftType == INTEGER && rightType == INTEGER {
		return left.(int) - right.(int)

	}

	// FLOATS

	if leftType == FLOAT && rightType == FLOAT {
		return left.(float64) - right.(float64)

	}

	// INTEGERS AND FLOATS

	if leftType == INTEGER && rightType == FLOAT {
		return float64(left.(int)) - right.(float64)

	}

	if leftType == FLOAT && rightType == INTEGER {
		return left.(float64) - float64(right.(int))

	}

	// STRINGS

	if leftType == STRING && rightType == STRING {
		return left.(string) + right.(string)
	}

	return nil
}

func SubOf(left any, right any) any {
	left, right = GetConstantValues(left, right)
	leftType, rightType := GetTypes(left, right)

	// INTEGERS

	if leftType == INTEGER && rightType == INTEGER {
		return left.(int) - right.(int)

	}

	// FLOATS

	if leftType == FLOAT && rightType == FLOAT {
		return left.(float64) - right.(float64)

	}

	// INTEGERS AND FLOATS

	if leftType == INTEGER && rightType == FLOAT {
		return float64(left.(int)) - right.(float64)

	}

	if leftType == FLOAT && rightType == INTEGER {
		return left.(float64) - float64(right.(int))

	}

	return nil
}

func MulOf(left any, right any) any {
	left, right = GetConstantValues(left, right)
	leftType, rightType := GetTypes(left, right)

	// INTEGERS

	if leftType == INTEGER && rightType == INTEGER {
		return left.(int) * right.(int)

	}

	// FLOATS

	if leftType == FLOAT && rightType == FLOAT {
		return left.(float64) * right.(float64)

	}

	// INTEGER AND FLOAT

	if leftType == INTEGER && rightType == FLOAT {
		return float64(left.(int)) * right.(float64)

	}

	if leftType == FLOAT && rightType == INTEGER {
		return left.(float64) * float64(right.(int))

	}

	// STRINGS AND INTEGERS

	fmt.Println(left, right)

	if leftType == STRING && rightType == INTEGER {
		var buffer []string = make([]string, 0)

		for i := 0; i < right.(int); i++ {
			buffer = append(buffer, left.(string))
		}

		return strings.Join(buffer, "")
	}

	if leftType == INTEGER && rightType == STRING {
		var buffer []string = make([]string, 0)

		for i := 0; i < left.(int); i++ {
			buffer = append(buffer, right.(string))
		}

		return strings.Join(buffer, "")
	}

	return nil
}

func DivOf(left any, right any) any {
	left, right = GetConstantValues(left, right)
	leftType, rightType := GetTypes(left, right)

	// INTEGERS

	if leftType == INTEGER && rightType == INTEGER {
		return left.(int) / right.(int)

	}

	// FLOATS

	if leftType == FLOAT && rightType == FLOAT {
		return left.(float64) / right.(float64)

	}

	// INTEGER AND FLOAT

	if leftType == INTEGER && rightType == FLOAT {
		return float64(left.(int)) / right.(float64)

	}

	if leftType == FLOAT && rightType == INTEGER {
		return left.(float64) / float64(right.(int))

	}

	return nil
}

func AndOf(left any, right any) any {
	left, right = GetConstantValues(left, right)
	leftType, rightType := GetTypes(left, right)

	fmt.Println(leftType, rightType)

	if leftType == rightType {

		if leftType == BOOLEAN {
			return left.(bool) && right.(bool)
		}

	}

	return false
}

func OrOf(left any, right any) any {
	left, right = GetConstantValues(left, right)
	leftType, rightType := GetTypes(left, right)

	if leftType == rightType {

		if leftType == BOOLEAN {
			return left.(bool) || right.(bool)
		}

	}

	return false
}

func Equals(left any, right any) any {
	left, right = GetConstantValues(left, right)
	leftType, rightType := GetTypes(left, right)

	if leftType == rightType {

		if leftType == INTEGER {
			return left.(int) == right.(int)
		}

		if leftType == FLOAT {
			return left.(float64) == right.(float64)
		}

		if leftType == STRING {
			return left.(string) == right.(string)
		}

		if leftType == BOOLEAN {
			return left.(bool) == right.(bool)
		}
	}

	return false
}

func NotEquals(left any, right any) any {
	return !(Equals(left, right).(bool))
}
