# gost
**Go API to Test Internet Speed using [speedtest.net](http://speedtest.net/)
or [fast.com](http://fast.com/)**


### API Usage
The code below tests download and upload speeds using speedtest.net and fast.com.
```go
package main

import (
	"context"
	"fmt"

	"github.com/user358/gost"
)

func main() {
	stProvider, err := gost.NewSpeedTestProvider()
	if err != nil {
		panic(err)
	}

	stResult, err := gost.Measure(context.Background(), stProvider)
	if err != nil {
		panic(err)
	}

	fmt.Printf("speedtest.net download: %5.2f Mbit/s\n", stResult.Download)
	fmt.Printf("speedtest.net upload: %5.2f Mbit/s\n\n", stResult.Upload)

	fastProvider, err := gost.NewFastProvider(false, nil)
	if err != nil {
		panic(err)
	}

	fastResult, err := gost.Measure(context.Background(), fastProvider)
	if err != nil {
		panic(err)
	}

	fmt.Printf("fast.com download: %5.2f Mbit/s\n", fastResult.Download)
	fmt.Printf("fast.com upload: %5.2f Mbit/s\n\n", fastResult.Upload)
}
```

## Test

```sh
go test -v
```

## Benchmark

```sh
go test -bench=.
```

## LICENSE

[MIT](https://github.com/user358/gost/blob/master/LICENSE)