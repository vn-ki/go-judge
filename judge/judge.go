package judge

import (
	"errors"
)

/*
GetJudge returns a runner for the language
Available languages and usage
	C++: GetJudge("cpp")
*/
func GetJudge(s string) (Judge, error) {
	runners := map[string]func() Judge{
		"cpp":     GetCPPRunner,
		"python3": GetPython3Runner,
		"python2": GetPython2Runner,
	}

	judge, ok := runners[s]
	if !ok {
		return Judge{}, errors.New("Judge not found")
	}
	return judge(), nil // TODO: ERROR handling
}

// Judge can compile and run a program
// 	{compiled_file}: The output file of compiler
// 	{source_file}  : The source file
type Judge struct {
	compileCmd       string
	compileArgs      []string
	runCmd           string
	runArgs          []string
	compileBeforeRun bool
	compiledName     string

	file      string
	extension string

	inputs []string

	output []string
}
