package gost

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeasureSpeedTest(t *testing.T) {
	p, _ := NewSpeedTestProvider()
	v, err := Measure(context.Background(), p)
	if assert.Nil(t, err) {
		assert.NotEqualValues(t, 0, v.Download)
		assert.NotEqualValues(t, 0, v.Upload)
	}
}

func TestMeasureFast(t *testing.T) {
	p, _ := NewFastProvider(false, nil)
	v, err := Measure(context.Background(), p)
	if assert.Nil(t, err) {
		assert.NotEqualValues(t, 0, v.Download)
		assert.NotEqualValues(t, 0, v.Upload)
	}
}

func BenchmarkSpeedTest(b *testing.B) {
	p, _ := NewSpeedTestProvider()
	for i := 0; i < b.N; i++ {
		_, _ = Measure(context.Background(), p)
	}
}

func BenchmarkSpeedFast(b *testing.B) {
	p, _ := NewFastProvider(false, nil)
	for i := 0; i < b.N; i++ {
		_, _ = Measure(context.Background(), p)
	}
}
