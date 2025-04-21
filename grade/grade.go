package grade

import (
	"encoding/json"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Subject struct {
	Name string
	// Homework
	Homework          []int
	PointsHomework    []int
	MaxPointsHomework []int // check later if I need a int or a float number
	// Tests
	Test          []int
	PointsTests   []int // check later if I need a int or a float number
	MaxPointsTest []int
}

type SubjectStore struct {
	subjects []Subject
}

func (s *SubjectStore) AddSubject(name string) {
	s.readJson()

	newSubject := Subject{
		Name:              name,
		Homework:          []int{},
		PointsHomework:    []int{},
		MaxPointsHomework: []int{},
		Test:              []int{},
		PointsTests:       []int{},
		MaxPointsTest:     []int{},
	}
	s.subjects = append(s.subjects, newSubject)

	s.writeJson()
}

func (s *SubjectStore) AddTest(data string) {
	s.readJson()

	subject, test, points, maxPoints := convert(data)
	for i := 0; i < len(s.subjects); i++ {
		if s.subjects[i].Name == subject {
			s.subjects[i].Test = append(s.subjects[i].Test, test)
			s.subjects[i].PointsTests = append(s.subjects[i].PointsTests, points)
			s.subjects[i].MaxPointsTest = append(s.subjects[i].MaxPointsTest, maxPoints)
		}
	}

	s.writeJson()
}

func (s *SubjectStore) AddHomework(data string) {
	s.readJson()

	subject, homework, points, maxPoints := convert(data)

	for i := 0; i < len(s.subjects); i++ {
		if s.subjects[i].Name == subject {
			s.subjects[i].Homework = append(s.subjects[i].Homework, homework)
			s.subjects[i].PointsTests = append(s.subjects[i].PointsHomework, points)
			s.subjects[i].MaxPointsTest = append(s.subjects[i].MaxPointsHomework, maxPoints)
		}
	}

	s.writeJson()
}

func (s *SubjectStore) RemoveTest(data string) {
	s.readJson()

	subject, test := convertTwoArg(data)

	for i := 0; i < len(s.subjects); i++ {
		if s.subjects[i].Name == subject {
			for j := 0; j < len(s.subjects[i].Test); j++ {
				if s.subjects[i].Test[j] == test {
					s.subjects[i].Test = slices.Delete(s.subjects[i].Test, j, j+1)
					s.subjects[i].PointsTests = slices.Delete(s.subjects[i].PointsTests, j, j+1)
					s.subjects[i].MaxPointsTest = slices.Delete(s.subjects[i].MaxPointsTest, j, j+1)
				}
			}
		}
	}

	s.writeJson()
}

func (s *SubjectStore) RemoveHomework(data string) {
	s.readJson()

	subject, Homework := convertTwoArg(data)

	for i := 0; i < len(s.subjects); i++ {
		if s.subjects[i].Name == subject {
			for j := 0; j < len(s.subjects[i].Homework); j++ {
				if s.subjects[i].Homework[j] == Homework {
					s.subjects[i].Homework = slices.Delete(s.subjects[i].Homework, j, j+1)
					s.subjects[i].PointsHomework = slices.Delete(s.subjects[i].MaxPointsHomework, j, j+1)
					s.subjects[i].MaxPointsHomework = slices.Delete(s.subjects[i].MaxPointsHomework, j, j+1)
				}
			}
		}
	}

	s.writeJson()
}

func convert(s string) (string, int, int, int) {
	splitted := strings.Split(s, " ")
	subject := splitted[0]
	firstArg, _ := strconv.Atoi(splitted[1])
	secondArg, _ := strconv.Atoi(splitted[2])
	thirdArg, _ := strconv.Atoi(splitted[3])

	return subject, firstArg, secondArg, thirdArg
}

func convertTwoArg(s string) (string, int) {
	splitted := strings.Split(s, " ")
	subject := splitted[0]
	firstArg, _ := strconv.Atoi(splitted[1])

	return subject, firstArg
}

func (s *SubjectStore) readJson() {
	// read file
	jsonData, err := os.ReadFile("subject.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(jsonData, &s.subjects)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (s *SubjectStore) writeJson() {
	jsonData, err := json.MarshalIndent(s.subjects, "", " ")
	if err != nil {
		log.Fatal(err)
		return
	}
	// write file
	err = os.WriteFile("subject.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
}
