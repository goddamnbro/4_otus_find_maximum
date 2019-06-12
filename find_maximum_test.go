package find_maximum

import (
	"github.com/pkg/errors"
	"testing"
)


func TestFindMaximumInt(t *testing.T) {
	numbers := []interface{}{10, 7, 3, 6, 6, 1, 0}
	expected := 10

	finder := func(elements []interface{}) (interface{}, error) {
		// we can't use default 0 for min value, because we can get slice with negative values
		maxValue := -1 << 63

		for _, element := range elements {
			num, ok := element.(int)
			if !ok {
				return nil, errors.New("Your comparator function doesn't match the type of slice")
			}
			if num > maxValue {
				maxValue = num
			}
		}
		return maxValue, nil
	}

	actual, err := findMax(numbers, finder)
	if err != nil {
		t.Errorf(err.Error())
	}
	if actual != expected {
		t.Errorf("Expected: %v, but actual: %v", expected, actual)
	}
}

func TestFindMaximumAge(t *testing.T) {
	type Person struct {
		Name string
		Age uint
	}

	persons  := []interface{}{
		Person{Name: "Ivan", Age: 20},
		Person{Name: "Semen", Age: 27},
		Person{Name: "Petr", Age: 31},
	}

	finder := func(elements []interface{}) (interface{}, error) {
		var maxValue uint

		for _, element := range elements {
			person, ok := element.(Person)
			if !ok {
				return nil, errors.New("Your comparator function doesn't match the type of slice")
			}
			if person.Age > maxValue {
				maxValue = person.Age
			}
		}
		return maxValue, nil
	}

	expected := uint(31)
	actual, err := findMax(persons, finder)
	if err != nil {
		t.Errorf(err.Error())
	}
	if actual != expected {
		t.Errorf("Expected: %v, but actual: %v", expected, actual)
	}
}
