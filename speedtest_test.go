package gost

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSpeedTestProvider(t *testing.T) {
	s, err := NewSpeedTestProvider()
	assert.Nil(t, err)
	assert.NotEmpty(t, s.servers)
}

func TestSpeedTestProvider_Download(t *testing.T) {
	s, _ := NewSpeedTestProvider()
	v, err := s.Download(context.Background())
	assert.Nil(t, err)
	assert.NotEqualValues(t, 0, v)
}

func TestSpeedTestProvider_Upload(t *testing.T) {
	s, _ := NewSpeedTestProvider()
	v, err := s.Upload(context.Background())
	assert.Nil(t, err)
	assert.NotEqualValues(t, 0, v)
}
