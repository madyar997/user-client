package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Option func(*Client)

func WithAddress(address string) Option {
	return func(client *Client) {
		client.address = address
	}
}

func WithInsecure() Option {
	return func(client *Client) {
		client.dialOpts = append(client.dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
}
