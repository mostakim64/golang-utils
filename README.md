
# Shadowchef Utils

Some common utils functions and log functions have been written here for common usage on other services. 

## Installation
Configure `.gitconfig` to include the below code
```
[url "git@bitbucket.org:"]
        insteadOf = https://bitbucket.org/

# this will enable the use of ssh version of the repo url instead of https.
```

To setup private bitbucket cloud
```
git config --global url."git@bitbucket.org:shadowchef".insteadOf "https://bitbucket.org/shadowchef"
```

Install Utils with `go get`.

```bash
GOPRIVATE=bitbucket.org/shadowchef/utils go get bitbucket.org/shadowchef/utils
```

## Usage
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
	header := "Alert"
	summary := "Alert Summary"
	details := "Details"
	serviceName := "Storage"
	metadata := "metadata"
	mentions := []string{"@here", "@there"}
	
	slackitClient := slackit.NewSlackitClient(webhookUrl)

	clientReq := slackit.ClientRequest{
		Header: header,
		ServiceName: serviceName,
		Summary: summary,
		Metadata: metadata,
		Details: details,
		Status: slackit.Alert,
		Mentions:    mentions,
	}

	err := slackitClient.Send(clientReq)
	if err != nil {
		fmt.Print("Alert occurred sending message to slack: ", err)
	}
}
```

### monitor package

```go
package main

import (
	"bitbucket.org/shadowchef/utils/monitor"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	monitor.NewEchoPrometheusClient(e, "/metrics")
}
```


To run tests, run the following command

```bash
go run main.go
```

