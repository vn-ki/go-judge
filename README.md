# go-judge
[![GoDoc](https://godoc.org/github.com/vn-ki/go-judge?status.svg)](https://godoc.org/github.com/vn-ki/go-judge)

```go
package main

import (
  "fmt"
  gojudge "github.com/vn-ki/go-judge"
)

func main() {
  judge := gojudge.GetJudge("cpp")
  judge.AddFile("a.cpp")
  judge.AddInputFile("input1.txt")
  
  out, err := runner.Output()
  
  if err != nil {
    panic(err)
  }
  
  fmt.Println(out)
}

```

See [documentation](https://godoc.org/github.com/vn-ki/go-judge/runner#Runner) for more info on judge (Runner)

### TODO

- Support more languages
- Check output
- Write tests
- Add support for compiler flags
- Rename `Runner` to `Judge`
