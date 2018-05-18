package runner

/*
GetCPPRunner returns a Runner with C++ configuration
*/
func GetCPPRunner() Runner {
	return Runner{
		compileCmd:       "g++",
		compileArgs:      []string{},
		runCmd:           "./a.out",
		runArgs:          []string{},
		compileBeforeRun: true,
		extension:        ".cpp",
	}
}
