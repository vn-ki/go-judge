# go-judge
[![GoDoc](https://godoc.org/github.com/vn-ki/go-judge?status.svg)](https://godoc.org/github.com/vn-ki/go-judge/judge)
[![Build Status](https://travis-ci.com/vn-ki/go-judge.svg?branch=master)](https://travis-ci.com/vn-ki/go-judge)
[![codecov](https://codecov.io/gh/vn-ki/go-judge/branch/master/graph/badge.svg)](https://codecov.io/gh/vn-ki/go-judge)

```go
package main

import (
  "fmt"
  gojudge "github.com/vn-ki/go-judge/judge"
)

func main() {
  judge := gojudge.GetJudge("cpp")
  judge.AddFile("a.cpp")
  judge.AddInputFile("input1.txt")
  
  out, err := judge.Output()
  
  if err != nil {
    panic(err)
  }
  
  fmt.Println(out)
}

```

See [documentation](https://godoc.org/github.com/vn-ki/go-judge/judge#Judge) for more info on `Judge`.

## How to add more languages

- Make a new file under `go-judge/judge` in the following format

```go
package judge

/*
GetLanguageRunner returns a Runner with Language configuration
*/
func GetLanguageRunner() Judge {
	return Judge{
		compileCmd:       "LanguageCompiler",
		compileArgs:      []string{"--compile-lang", "{source_file}"},
		runCmd:           "./{compiled_program}",
		runArgs:          []string{},
		compileBeforeRun: true,
		extension:        ".lang",
	}
}

```

- Add an entry to the `map` in `go-judge/judge/judge.go`.
- Profit!!

### TODO

- Support more languages
- Check output
- Write tests
- Add support for compiler flags
