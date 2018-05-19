package judge

import (
	"errors"
	"strings"
)

func checkExtension(fileName string, extension string) bool {
	return strings.HasSuffix(fileName, extension)
}

// AddFile adds a file to the Judge which can then be (compiled and) run.
func (judge *Judge) AddFile(file string) error {

	if checkExtension(file, judge.extension) {
		judge.file = file
	} else {
		return errors.New("File doesnot match the extension")
	}

	return nil
}

// AddInputFile adds an input file to run the program with
func (judge *Judge) AddInputFile(file string) error {
	if checkExtension(file, ".txt") {
		judge.inputs = append(judge.inputs, file)
	} else {
		return errors.New("Input file extension doesnot match '.txt'")
	}

	return nil
}
