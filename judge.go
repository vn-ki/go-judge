package judge

import (
	"github.com/vn-ki/go-judge/runner"
)

/*
GetJudge returns a runner for the language
Available languages and usage
	C++: GetJudge("cpp")
*/
func GetJudge(s string) runner.Runner {
	runners := map[string]func() runner.Runner{
		"cpp": runner.GetCPPRunner,
	}

	return runners[s]()
}
