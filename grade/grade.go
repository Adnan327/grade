package grade

import (
	"strconv"
	"strings"
)

type Subject struct {
	name string
	// homework
	homework          []int
	pointsHomework    []int
	maxPointsHomework []int // check later if I need a int or a float number
	// tests
	test          []int
	pointsTests   []int // check later if I need a int or a float number
	maxPointsTest []int
}

type SubjectStore struct {
	subjects []Subject
}

func (s *SubjectStore) AddSubject(name string) {
	newSubject := Subject{
		name: name,
	}
	s.subjects = append(s.subjects, newSubject)
}

func (s *SubjectStore) AddTest(data string) {
	subject, test, points, maxPoints := convert(data)

	for _, element := range s.subjects {
		if element.name == subject {
			element.test = append(element.test, test)
			element.pointsTests = append(element.pointsTests, points)
			element.maxPointsTest = append(element.maxPointsTest, maxPoints)
		}
	}
}

func (s *SubjectStore) AddHomework(data string) {
	subject, homework, points, maxPoints := convert(data)

	for _, element := range s.subjects {
		if element.name == subject {
			element.test = append(element.test, homework)
			element.pointsTests = append(element.pointsTests, points)
			element.maxPointsTest = append(element.maxPointsTest, maxPoints)
		}
	}
}

func convert(s string) (string, int, int, int) {
	splitted := strings.Split(s, " ")
	subject := splitted[0]
	firstArg, _ := strconv.Atoi(splitted[1])
	secondArg, _ := strconv.Atoi(splitted[2])
	thirdArg, _ := strconv.Atoi(splitted[3])

	return subject, firstArg, secondArg, thirdArg
}

func (s *SubjectStore) RemoveTest(data string) {

}

func (s *SubjectStore) RemoveHomework(data string) {

}

// extend the programm at the end with a change-function
