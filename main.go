package main

// import "fmt"
import (
	"os"

	"github.com/Adnan327/grader/grade"
)

func main() {

	args := os.Args

	switch args[1] {
	case "add":
		grade.AddSubject(args[2])
		grade.AddHomework(args[2], args[3], args[4], args[5])
		grade.AddTest(args[2], args[3], args[4], args[5])

	case "rm":

	}

}
