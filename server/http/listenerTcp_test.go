package http

import "testing"

func TestNewListenerTcp(t *testing.T) {
	l, err := NewListenerTcp("")
	if err != nil {
		panic("create fail")
	}
	l.Close()

	l, err = NewListenerTcp(":oo")
	if err == nil {
		panic("listen error")
	}
}
