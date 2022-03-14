
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

### slackit package
```go
package main

import (
	"bitbucket.org/shadowchef/utils/slackit"
	"fmt"
)

func main() {
	webhookUrl := "https://hooks.slack.com/services/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"
	summary := "Alert Summary"
	details := "Details"
	serviceName := "Storage"
	metadata := "metadata"
	slackitClient := slackit.NewSlackitClient(webhookUrl)

	clientReq := slackit.ClientRequest{
		ServiceName: serviceName,
		Summary: summary,
		Metadata: metadata,
		Details: details,
		Status: slackit.Alert,
	}

	err := slackitClient.Send(clientReq)
	if err != nil {
		fmt.Print("Alert occurred sending message to slack: ", err)
	}
}
```


To run tests, run the following command

```bash
go run main.go
```

