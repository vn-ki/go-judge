package judge

import (
	"strings"
	"testing"
)

func Test_cppfull(t *testing.T) {
	judge := GetCPPJudge(Config{
		SourceFile: "testSourceFiles/cpp_test.cpp",
		InputFiles: []string{"testSourceFiles/cpp_input.txt"},
	})

	out, err := judge.Output()

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if strings.Compare(out[0], "123") != 0 {
		t.Logf("Expected '123', Got '%s'", out)
		t.Fail()
	}
}
