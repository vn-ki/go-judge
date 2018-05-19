package judge

/*
GetPython3Runner returns a Runner with python configuration
*/
func GetPython3Runner() Judge {
	return Judge{
		compileCmd:       "",
		compileArgs:      []string{},
		runCmd:           "python3",
		runArgs:          []string{"{source_file}"},
		compileBeforeRun: false,
		extension:        ".py",
	}
}
