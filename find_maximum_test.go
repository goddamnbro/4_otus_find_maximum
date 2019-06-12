package find_maximum

import (
	"testing"
)

func TestFindMaximumInt(t *testing.T) {
	numbers := []interface{}{10, 7, 3, 6, 6, 1, 100}

	actual := FindMax(numbers, func(currentIndex, maxValueIndex int) bool {
		return numbers[currentIndex].(int) > numbers[maxValueIndex].(int)
	})

	expected := 100
	if actual != expected {
		t.Errorf("Expected: %v, but actual: %v", expected, actual)
	}
}

func TestFindMaximumAge(t *testing.T) {
	type Person struct {
		Name string
		Age  uint
	}

	persons := []interface{}{
		Person{Name: "Ivan", Age: 20},
		Person{Name: "Semen", Age: 27},
		Person{Name: "Petr", Age: 31},
	}

	expected := Person{Name: "Petr", Age: 31}
	actual := FindMax(persons, func(currentIndex, maxValueIndex int) bool {
		person := persons[currentIndex].(Person)
		oldestPerson := persons[maxValueIndex].(Person)
		return person.Age > oldestPerson.Age
	})
	if actual != expected {
		t.Errorf("Expected: %v, but actual: %v", expected, actual)
	}
}
