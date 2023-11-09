package grpc

import (
	"context"
	"fmt"
	"github.com/madyar997/user-client/client"
	"github.com/madyar997/user-client/models"
	"github.com/madyar997/user-client/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	address  string
	conn     *grpc.ClientConn
	client   protobuf.UserClient
	dialOpts []grpc.DialOption
}

func (c *Client) Connect() error {
	conn, err := grpc.Dial(c.address, c.dialOpts...)
	if err != nil {
		return fmt.Errorf("error establishing gRPC connection to nats-streaming-reader: %s", err.Error())
	}

	c.conn = conn

	c.setupClient()

	return nil
}

func (c *Client) setupClient() {
	c.client = protobuf.NewUserClient(c.conn)
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) GetUserByID(ctx context.Context, id int32) (*models.User, error) {
	resp, err := c.client.GetUserByID(ctx, &protobuf.UserRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:    resp.Id,
		Name:  resp.Name,
		Email: resp.Email,
		Age:   resp.Age,
	}, nil
}

func NewClient(conf *client.Config) (client.Client, error) {
	cli := &Client{
		address:  conf.Address,
		dialOpts: []grpc.DialOption{},
	}

	if conf.Insecure {
		cli.dialOpts = append(cli.dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return cli, nil
}
