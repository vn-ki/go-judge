package judge

/*
GetCPPJudge returns a Judge with C++ configuration
*/
func GetCPPJudge(config Config) Judge {
	return addConfig(
		Judge{
			compileCmd:       "g++",
			compileArgs:      []string{"{source_file}"},
			runCmd:           ".compiled_bin/{compiled_program}", // TODO: Don't hardcode compiled folder
			runArgs:          []string{},
			compileBeforeRun: true,
			extension:        ".cpp",
		}, config)
}
