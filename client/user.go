package client

import (
	"context"
	"github.com/madyar997/user-client/models"
)

type Config struct {
	Address  string
	Protocol string
	Insecure bool
}

// add ping later
type Client interface {
	Connect() error
	Close() error
	GetUserByID(ctx context.Context, id int32) (*models.User, error)
}
