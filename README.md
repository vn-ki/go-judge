# go-judge


```go
package main

import (
  "fmt"
  judge "github.com/vn-ki/go-judge"
)

func main() {
  judge := judge.GetJudge("cpp")
  judge.AddFile("a.cpp")
  judge.AddInputFile("input1.txt")
  
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
