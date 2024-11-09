package utils

type StringArray []string

func IntContains(array []int, element int) bool {

	for i := 0; i < len(array); i++ {
		if array[i] == element {
			return true
		}
	}

	return false
}

func StrContains(array []string, element string) bool {

	for i := 0; i < len(array); i++ {
		if array[i] == element {
			return true
		}
	}

	return false
}

func (array StringArray) Contains(element string) bool {
	return StrContains(array, element)
}
