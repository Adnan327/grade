package main

// import "fmt"
import (
	"flag"
	"os"

	"github.com/Adnan327/grader/internal/grader"
)

func main() {
	command := os.Args[1]
	args := os.Args[2:]
	var (
		subjectList  grader.SubjectStore
		subjectFlag  string
		homeworkFlag string
		testFlag     string
	)
	flag := flag.NewFlagSet(command, flag.ExitOnError)
	flag.StringVar(&subjectFlag, "s", "", "help message for flagname")
	flag.StringVar(&homeworkFlag, "h", "", "help message for flagname")
	flag.StringVar(&testFlag, "t", "", "help message for flagname")
	flag.Parse(args)

	switch command {
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
	case "list":
		subjectList.GetTotalPercentage()
	}
}
