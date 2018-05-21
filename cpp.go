package judge

/*
GetCPPJudge returns a Judge with C++ configuration
*/
func GetCPPJudge(config Config) Judge {
	return addConfig(
		Judge{
			compileCmd:       "g++",
			compileArgs:      []string{"{source_file}"},
			runCmd:           "./{compiled_program}",
			runArgs:          []string{},
			compileBeforeRun: true,
			extension:        ".cpp",
		}, config)
}
