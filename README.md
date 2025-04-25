# Gradr - Points Manager CLI App

This is a small CLI tool to help you track your test and homework points. It shows you the percentage score of each subject at a glance.

Go must be installed. When it is installed, execute the following command:
```shell
go install github.com/Adnan327/grader/cmd/gradr@latest
```


# Commands

Add a new Subject
```shell
gradr add -s "<subject-name>"
```
Add a new homework
```shell
gradr add -h "<subject-name> <number> <points> <max-points>"
```
Add a new test
```shell
gradr add -t "<subject-name> <number> <points> <max-points>"
```

Remove homework
```shell
gradr rm -h "<subject-name> <number>"
```
Remove test
```shell
gradr rm -t "<subject-name> <number>"
```

List all subject with its total percentages in table form
```shell
gradr list
```