package http

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestNewListenerTcp(t *testing.T) {
	assert := assert2.New(t)
	l, err := NewListenerTcp("")
	assert.Nil(err)
	l.Close()

	l, err = NewListenerTcp(":oo")
	assert.NotNil(err)
}
