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
	Name         string
	HomeWorkList []Homework
	TestList     []Test
}

type Homework struct {
	Homework  int
	Points    int
	MaxPoints int
}

type Test struct {
	Test      int
	Points    int
	MaxPoints int
}

type SubjectStore struct {
	subjects []Subject
}

func (s *SubjectStore) AddSubject(name string) {
	s.readJson()

	newSubject := Subject{
		Name:         name,
		HomeWorkList: []Homework{},
		TestList:     []Test{},
	}
	s.subjects = append(s.subjects, newSubject)

	s.writeJson()
}

func (s *SubjectStore) AddTest(data string) {
	s.readJson()

	subject, test, points, maxPoints := convert(data)
	newTest := Test{
		Test:      test,
		Points:    points,
		MaxPoints: maxPoints,
	}
	for i := 0; i < len(s.subjects); i++ {
		if s.subjects[i].Name == subject {
			s.subjects[i].TestList = append(s.subjects[i].TestList, newTest)
		}
	}

	s.writeJson()
}

func (s *SubjectStore) AddHomework(data string) {
	s.readJson()

	subject, homework, points, maxPoints := convert(data)
	newHomework := Homework{
		Homework:  homework,
		Points:    points,
		MaxPoints: maxPoints,
	}
	for i := 0; i < len(s.subjects); i++ {
		if s.subjects[i].Name == subject {
			s.subjects[i].HomeWorkList = append(s.subjects[i].HomeWorkList, newHomework)
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
