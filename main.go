package main

// import "fmt"
import (
	"flag"
	"os"

	"github.com/Adnan327/grader/grade"
)

func main() {
	var (
		args         = os.Args
		subjectList  grade.SubjectStore
		subjectFlag  string
		homeworkFlag string
		testFlag     string
	)
	flag.StringVar(&subjectFlag, "-s", "", "help message for flagname")
	flag.StringVar(&homeworkFlag, "-h", "", "help message for flagname")
	flag.StringVar(&testFlag, "-t", "", "help message for flagname")

	switch args[1] {
	case "add":
		if subjectFlag != "" {
			subjectList.AddSubject(subjectFlag)
		} else if homeworkFlag != "" {
			subjectList.AddHomework(homeworkFlag)
		} else if testFlag != "" {
			subjectList.AddTest(testFlag)
		}

	case "rm":
		if homeworkFlag != "" {
			subjectList.RemoveHomework(homeworkFlag)
		} else if testFlag != "" {
			subjectList.RemoveTest(testFlag)
		}
	}
}
