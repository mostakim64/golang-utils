
# Shadowchef Utils

Some utils functions and logger have been integrated here. 

## Installation

Install Utils with go mod.

```bash
  GOPRIVATE=bitbucket.org/shadowchef/utils go get bitbucket.org/shadowchef/utils
```
    
## Running Tests



```go
package main
import (
	"fmt"

	"bitbucket.org/shadowchef/utils/methods"
)

func main() {
	fmt.Println(methods.GenerateRandomStringOfLength(20))
}
```

To run tests, run the following command

```bash
  go run main.go
```

