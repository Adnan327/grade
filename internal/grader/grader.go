package grader

import (
	"encoding/json"
	"fmt"
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
	Homework  int // this is the number of the homework
	Points    int
	MaxPoints int
}

type Test struct {
	Test      int // this is the number of the test
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
			for j := 0; j < len(s.subjects[i].TestList); j++ {
				if s.subjects[i].TestList[j].Test == test {
					s.subjects[i].TestList = slices.Delete(s.subjects[i].TestList, j, j+1)
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
			for j := 0; j < len(s.subjects[i].HomeWorkList); j++ {
				if s.subjects[i].HomeWorkList[j].Homework == Homework {
					s.subjects[i].HomeWorkList = slices.Delete(s.subjects[i].HomeWorkList, j, j+1)
				}
			}
		}
	}

	s.writeJson()
}

func (s *SubjectStore) GetTotalPercentage() {
	s.readJson()
	fmt.Printf("~Total points of all subjects~\n\n")
	fmt.Printf("%-15s | %-15s | %-20s\n", "Subject", "Homeworks", "Tests")
	fmt.Println(strings.Repeat("-", 55))
	var sumPointsHomework, sumTotalPointsHomework, sumPointsTest, sumTotalPointsTest float32
	for i := 0; i < len(s.subjects); i++ {
		sumPointsHomework = 0
		sumTotalPointsHomework = 0
		sumPointsTest = 0
		sumTotalPointsTest = 0
		for j := 0; j < len(s.subjects[i].HomeWorkList); j++ {
			sumPointsHomework += float32(s.subjects[i].HomeWorkList[j].Points)
			sumTotalPointsHomework += float32(s.subjects[i].HomeWorkList[j].MaxPoints)
		}

		for j := 0; j < len(s.subjects[i].TestList); j++ {
			sumPointsTest += float32(s.subjects[i].TestList[j].Points)
			sumTotalPointsTest += float32(s.subjects[i].TestList[j].MaxPoints)
		}

		hwstr := "N/A"
		ttstring := "N/A"
		if sumTotalPointsHomework != 0 {
			hwstr = fmt.Sprintf("%.2f %%", (sumPointsHomework*100)/sumTotalPointsHomework)
		}
		if sumTotalPointsTest != 0 {
			ttstring = fmt.Sprintf("%.2f %%", (sumPointsTest*100)/sumTotalPointsTest)
		}
		fmt.Printf("%-15s | %-15s | %-20s\n", s.subjects[i].Name, hwstr, ttstring)
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
