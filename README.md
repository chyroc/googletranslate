# googletranslate

## Installation

```bash
go get github.com/chyroc/googletranslate
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/chyroc/googletranslate"
)

func main() {
    fmt.Println(googletranslate.Translate("Hello", "en", "zh"))
    // output: 你好, nil
}
```