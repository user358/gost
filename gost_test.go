package gost

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeasureSpeedTest(t *testing.T) {
	p, _ := NewSpeedTestProvider()
	v, err := Measure(context.Background(), p)
	if assert.Nil(t, err) {
		assert.NotEqualValues(t, 0, v.Download)
		assert.NotEqualValues(t, 0, v.Upload)
		fmt.Println(v.Download, v.Upload)
	}
}

func TestMeasureFast(t *testing.T) {
	p, _ := NewFastProvider()
	v, err := Measure(context.Background(), p)
	if assert.Nil(t, err) {
		assert.NotEqualValues(t, 0, v.Download)
		assert.NotEqualValues(t, 0, v.Upload)
		fmt.Println(v.Download, v.Upload)
	}
}
