package judge

/*
GetPython3Judge returns a Judge with python configuration
*/
func GetPython3Judge() Judge {
	return Judge{
		compileCmd:       "",
		compileArgs:      []string{},
		runCmd:           "python3",
		runArgs:          []string{"{source_file}"},
		compileBeforeRun: false,
		extension:        ".py",
	}
}
