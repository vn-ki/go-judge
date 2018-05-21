# go-judge
[![GoDoc](https://godoc.org/github.com/vn-ki/go-judge?status.svg)](https://godoc.org/github.com/vn-ki/go-judge)
[![Build Status](https://travis-ci.com/vn-ki/go-judge.svg?branch=master)](https://travis-ci.com/vn-ki/go-judge)
[![codecov](https://codecov.io/gh/vn-ki/go-judge/branch/master/graph/badge.svg)](https://codecov.io/gh/vn-ki/go-judge)

A go package which (compiles and) executes a program and checks it with given of inputs. This can be used to build an online judge (or a web IDE).

### Supported Languages
- C++
- Python 3
- Python 2


```go
package main

import (
  "fmt"
  gojudge "github.com/vn-ki/go-judge/judge"
)

func main() {
  judge := gojudge.GetCPPJudge(gojudge.Config{})
  judge.AddSourceFile("a.cpp")
  judge.AddInputFile("input1.txt")
  
  out, err := judge.Output()
  
  if err != nil {
    panic(err)
  }
  
  fmt.Println(out)
}

```

You can remove the use of `AddSourceFile` and `AddInputFile` by using [`Config`](https://godoc.org/github.com/vn-ki/go-judge#Config).

```go
judge := gojudge.GetCPPJudge(gojudge.Config{
  SourceFile: "a.cpp",
  InputFiles: []string{"input1.txt"},
})
```

See [documentation](https://godoc.org/github.com/vn-ki/go-judge#Judge) for more info on `Judge`.


## How to add more languages

- Make a new file under `go-judge/judge` in the following format

```go
package judge

/*
GetLanguageRunner returns a Runner with Language configuration
*/
func GetLanguageRunner(config Config) Judge {
	return addConfig(
    Judge{
      compileCmd:       "LanguageCompiler",
      compileArgs:      []string{"--compile-lang", "{source_file}"},
      runCmd:           "./{compiled_program}",
      runArgs:          []string{},
      compileBeforeRun: true,
      extension:        ".lang",
	  }, config)
}

```

- Profit!!

### TODO

- Support more languages
- Check output
- Write tests
- Add support for compiler flags
