package judge

// Config can be passed to the decorator WithConfig to get a Judge using
// configuration
type Config struct {
	CompileArgs []string
	RunArgs     []string

	SourceFile     string
	InputFiles     []string
	ExpectedOutput []string
}

// WithConfig is a decorator to be used with getLanguageJudge
func WithConfig(getJudge func() Judge) func(Config) Judge {
	return func(config Config) Judge {
		judge := getJudge()

		judge.AddSourceFile(config.SourceFile)

		for _, file := range config.InputFiles {
			judge.AddInputFile(file) // TODO Exception handling
		}

		// judge.AddExpectedOutput() // TODO Once we add checking of output

		return judge
	}
}
