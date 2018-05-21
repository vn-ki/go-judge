package judge

/*
GetPython3Judge returns a Judge with python configuration
*/
func GetPython3Judge(config Config) Judge {
	return addConfig(
		Judge{
			compileCmd:       "",
			compileArgs:      []string{},
			runCmd:           "python3",
			runArgs:          []string{"{source_file}"},
			compileBeforeRun: false,
			extension:        ".py",
		}, config)
}
