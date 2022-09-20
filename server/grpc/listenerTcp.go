package grpc

import "net"

func NewListenerTcp(address string) (net.Listener, error) {
	if address == "" {
		address = ":9090"
	}
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return listener, err
	}
	return listener, nil
}
