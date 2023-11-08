# didyoumean

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![codecov](https://codecov.io/gh/AL-Kost/didyoumean/branch/main/graph/badge.svg)](https://codecov.io/gh/AL-Kost/didyoumean)
[![Go Reference](https://pkg.go.dev/badge/github.com/AL-Kost/didyoumean.svg)](https://pkg.go.dev/github.com/AL-Kost/didyoumean)
[![Go Report Card](https://goreportcard.com/badge/github.com/AL-Kost/didyoumean)](https://goreportcard.com/report/github.com/AL-Kost/didyoumean)


<p align="center">
  <img src="utils/logo.png" alt="Application Logo" width="300">
</p>

`didyoumean` is a Go package that suggests the most likely string based on the Levenshtein distance algorithm. It is particularly useful for correcting user typos in command-line tools or offering suggestions in search inputs.

## Getting Started

### Prerequisites

You need to have Go 1.15 or higher installed on your machine.

### Installing

To use `didyoumean` in your Go project, run:

```shell
go get -u github.com/AL-Kost/didyoumean
```
This will retrieve the library and add it to your project's dependencies.

## Usage

Import DidYouMean in your Go code:
```go
import "github.com/AL-Kost/didyoumean"
```

Example of suggesting a string:

```go
package main

import (
	"fmt"
	"github.com/AL-Kost/didyoumean"
)

func main() {
	input := "appl"
	candidates := []string{"apple", "application", "apology", "apply", "happy"}
	threshold := 2
	caseInsensitive := true
	
	closestMatch := didyoumean.Suggest(input, candidates, threshold, caseInsensitive)
	
	if closestMatch != "" {
		fmt.Printf("Did you mean '%s'?\n", closestMatch)
	} else {
		fmt.Println("No close match found.")
	}
}

```

For more examples, please refer to the `examples` directory.