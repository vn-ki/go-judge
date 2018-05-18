# go-judge


```go
package main

import (
  "fmt"
  judge "github.com/vn-ki/go-judge"
)

func main() {
  runner := judge.GetJudge("cpp")
  runner.AddFile("a.cpp")
  runner.AddInputFile("input1.txt")
  
  out, err := runner.Output()
  
  if err != nil {
    panic(err)
  }
  
  fmt.Println(out)
}

```

### TODO

- Support more languages
- Check output
- Write tests
- Add support for compiler flags
