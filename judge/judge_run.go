package judge

import (
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Compile compiles the given file
func (judge *Judge) Compile() error {
	judge.compiledName = strconv.FormatInt(time.Now().Unix(), 10)

	args := append(judge.compileArgs, judge.file, "-o", judge.compiledName)

	cmd := exec.Command(judge.compileCmd, args...)
	_, err := cmd.Output()
	return err
}

// Output (compiles and) runs a given file with the given input files and returns
// the output if there is no previous runs. If there has been a previous run,
// Output returns the previous output. Use Judge.Run to rerun the file
func (judge *Judge) Output() ([]string, error) {
	if len(judge.output) > 0 {
		return judge.output, nil
	}

	return judge.Run()
}

// Run (compiles and) runs a given file with the given input files and returns
// the output.
func (judge *Judge) Run() ([]string, error) {
	judge.replaceCompileCommands()

	if judge.compileBeforeRun {
		err := judge.Compile()
		if err != nil {
			return nil, err
		}
	}

	// Should be used after compile
	judge.replaceRunCommands()

	var ret []string

	for _, input := range judge.inputs {

		out, err := judge.run(input)
		if err != nil {
			ret = append(ret, err.Error())
		} else {
			ret = append(ret, string(out))
		}
	}

	judge.output = ret

	return judge.output, nil
}

func (judge *Judge) run(inputFile string) ([]byte, error) {
	cmd := exec.Command(judge.runCmd, judge.runArgs...)
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()

	cmd.Start()

	file, _ := ioutil.ReadFile(inputFile) // TODO: Error handling
	stdin.Write(file)

	out := make([]byte, 1024)

	n, err := stdout.Read(out)

	return out[:n], err
}

func (judge *Judge) replaceCommands(cmd string) string {

	replacer := strings.NewReplacer("{source_file}", judge.file, "{compiled_program}", judge.compiledName)

	return replacer.Replace(cmd)
}

func (judge *Judge) replaceCompileCommands() {
	judge.compileCmd = judge.replaceCommands(judge.compileCmd)

	for idx, val := range judge.compileArgs {
		judge.compileArgs[idx] = judge.replaceCommands(val)
	}
}

func (judge *Judge) replaceRunCommands() {
	judge.runCmd = judge.replaceCommands(judge.runCmd)

	for idx, val := range judge.runArgs {
		judge.runArgs[idx] = judge.replaceCommands(val)
	}
}
