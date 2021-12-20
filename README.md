
# Shadowchef Utils

Some common utils functions and log functions have been written here for common usage on other services. 

## Installation
Configure `.gitconfig` to include the below code
```
[url "git@bitbucket.org:"]
        insteadOf = https://bitbucket.org/

# this will enable the use of ssh version of the repo url instead of https.
```

Install Utils with `go get`.

```bash
GOPRIVATE=bitbucket.org/shadowchef/utils go get bitbucket.org/shadowchef/utils
```

## Useage
### Methods package
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

### Logger package
```go
package main
import (
	"fmt"

	"bitbucket.org/shadowchef/utils/logger"
)

func main() {
	logger.Info("put your message here...")
}
```

To run tests, run the following command

```bash
go run main.go
```

