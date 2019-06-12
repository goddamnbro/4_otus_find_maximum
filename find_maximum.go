package find_maximum

func FindMax(items []interface{}, greater func(currentIndex, maxValueIndex int) bool) interface{} {
	indexWithMaxValue := -1

	lastIndex := len(items) - 1

	for index := range items {
		if indexWithMaxValue == -1 {
			indexWithMaxValue = index
		}

		if index <= lastIndex {
			if greater(index, indexWithMaxValue) {
				indexWithMaxValue = index
			}
		}
	}

	return items[indexWithMaxValue]
}
