package find_maximum

func findMax(elements []interface{}, fn func(elements []interface{}) (interface{}, error)) (interface{}, error) {
	return fn(elements)
}
