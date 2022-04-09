package gost

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFastProvider(t *testing.T) {
	s, err := NewFastProvider(false, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, s.urls)
}

func TestFastProvider_Download(t *testing.T) {
	s, _ := NewFastProvider(false, nil)
	v, err := s.Download(context.Background())
	assert.Nil(t, err)
	assert.NotEqualValues(t, 0, v)
}

func TestFastProvider_Upload(t *testing.T) {
	s, _ := NewFastProvider(false, nil)
	v, err := s.Upload(context.Background())
	assert.Nil(t, err)
	assert.NotEqualValues(t, 0, v)
}

func Test_downloadRequest(t *testing.T) {
	s, _ := NewFastProvider(false, nil)
	err := downloadRequest(context.Background(), s.doer, s.urls[0])
	assert.Nil(t, err)
}

func Test_uploadRequest(t *testing.T) {
	s, _ := NewFastProvider(false, nil)
	err := uploadRequest(context.Background(), s.doer, s.urls[0], 0)
	assert.Nil(t, err)
}

func Test_dlWarmUp(t *testing.T) {
	s, _ := NewFastProvider(false, nil)
	err := dlWarmUp(context.Background(), s.doer, s.urls[0])
	assert.Nil(t, err)
}

func Test_ulWarmUp(t *testing.T) {
	s, _ := NewFastProvider(false, nil)
	err := ulWarmUp(context.Background(), s.doer, s.urls[0])
	assert.Nil(t, err)
}
