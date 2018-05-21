package judge

/*
GetPython2Judge returns a Runner with python configuration
*/
func GetPython2Judge(config Config) Judge {
	return addConfig(
		Judge{
			compileCmd:       "",
			compileArgs:      []string{},
			runCmd:           "python2",
			runArgs:          []string{"{source_file}"},
			compileBeforeRun: false,
			extension:        ".py",
		},
		config)
}
