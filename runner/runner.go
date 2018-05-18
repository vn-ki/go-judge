package runner

import (
	"errors"
	"io/ioutil"
	"os/exec"
	"strings"
)

// Runner can compile and run a program
type Runner struct {
	compileCmd       string
	compileArgs      []string
	runCmd           string
	runArgs          []string
	compileBeforeRun bool

	file      string
	extension string

	inputs []string

	output []string
}

func checkExtension(fileName string, extension string) bool {
	return strings.HasSuffix(fileName, extension)
}

// AddFile adds a file to the Runner which can then be (compiled and) run.
func (runner *Runner) AddFile(file string) error {

	if checkExtension(file, runner.extension) {
		runner.file = file
	} else {
		return errors.New("File doesnot match the extension")
	}

	return nil
}

// AddInputFile adds an input file to run the program with
func (runner *Runner) AddInputFile(file string) error {
	if checkExtension(file, ".txt") {
		runner.inputs = append(runner.inputs, file)
	} else {
		return errors.New("Input file extension doesnot match '.txt'")
	}

	return nil
}

// Compile compiles the given file
func (runner *Runner) Compile() error {
	args := append(runner.compileArgs, runner.file)
	cmd := exec.Command(runner.compileCmd, args...)
	_, err := cmd.Output()
	return err
}

// Output (compiles and) runs a given file with the given input files and returns
// the output if there is no previous runs. If there has been a previous run,
// Output returns the previous output. Use Runner.Run to rerun the file
func (runner *Runner) Output() ([]string, error) {
	if len(runner.output) > 0 {
		return runner.output, nil
	}

	return runner.Run()
}

// Run (compiles and) runs a given file with the given input files and returns
// the output.
func (runner *Runner) Run() ([]string, error) {

	if runner.compileBeforeRun {
		err := runner.Compile()
		if err != nil {
			return nil, err
		}
	}

	var ret []string

	for _, input := range runner.inputs {

		out, err := runner.run(input)
		if err != nil {
			ret = append(ret, err.Error())
		} else {
			ret = append(ret, string(out))
		}
	}

	runner.output = ret

	return runner.output, nil
}

func (runner *Runner) run(inputFile string) ([]byte, error) {
	cmd := exec.Command(runner.runCmd, runner.runArgs...)
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()

	cmd.Start()

	file, _ := ioutil.ReadFile(inputFile) // TODO: Error handling
	stdin.Write(file)

	out := make([]byte, 1024)

	n, err := stdout.Read(out)

	return out[:n], err
}
