package interpreter

func GetUnConstantValue(value any) any {
	if valueConstant, ok := value.(ConstantValue); ok {
		value, _ = GetConstantValue(valueConstant)
	}

	return value
}

func GetType(value any) int {
	var t = -1

	_, isInteger := value.(int)
	_, isFloat := value.(float64)
	_, isString := value.(string)
	_, isBool := value.(bool)

	if isInteger {
		t = INTEGER
	}

	if isFloat {
		t = FLOAT
	}

	if isString {
		t = STRING
	}

	if isBool {
		t = BOOLEAN
	}

	return t
}

func Not(operand any) any {
	operand = GetUnConstantValue(operand)
	var operandType = GetType(operand)

	if operandType == BOOLEAN {
		return !operand.(bool)
	}

	return false
}

var UnaryOperatorFunctions map[string]func(operand any) any = map[string]func(operand any) any{
	"!": Not,
}
