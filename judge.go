package judge

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
