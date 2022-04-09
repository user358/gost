package gost

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFastProvider(t *testing.T) {
	s, err := NewFastProvider()
	assert.Nil(t, err)
	assert.NotEmpty(t, s.urls)
}

func TestFastProvider_Download(t *testing.T) {
	s, _ := NewFastProvider()
	v, err := s.Download(context.Background())
	assert.Nil(t, err)
	assert.NotEqualValues(t, 0, v)
}

func TestFastProvider_Upload(t *testing.T) {
	s, _ := NewFastProvider()
	v, err := s.Upload(context.Background())
	assert.Nil(t, err)
	assert.NotEqualValues(t, 0, v)
}
