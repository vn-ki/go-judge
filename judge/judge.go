package judge

/*
GetJudge returns a runner for the language
Available languages and usage
	C++: GetJudge("cpp")
*/
func GetJudge(s string) Judge {
	runners := map[string]func() Judge{
		"cpp":     GetCPPRunner,
		"python3": GetPython3Runner,
	}

	return runners[s]() // TODO: ERROR handling
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
