package judge

/*
GetPython2Runner returns a Runner with python configuration
*/
func GetPython2Runner() Judge {
	return Judge{
		compileCmd:       "",
		compileArgs:      []string{},
		runCmd:           "python2",
		runArgs:          []string{"{source_file}"},
		compileBeforeRun: false,
		extension:        ".py",
	}
}
