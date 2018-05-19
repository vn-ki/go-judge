package judge

/*
GetCPPRunner returns a Runner with C++ configuration
*/
func GetCPPRunner() Judge {
	return Judge{
		compileCmd:       "g++",
		compileArgs:      []string{},
		runCmd:           "./{compiled_program}",
		runArgs:          []string{},
		compileBeforeRun: true,
		extension:        ".cpp",
	}
}