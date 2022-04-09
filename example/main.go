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

	fastProvider, err := gost.NewFastProvider()
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
