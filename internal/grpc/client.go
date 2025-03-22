package client

import (
	"context"
	"fmt"
	"time"

	product "github.com/berezovskyivalerii/server-rpc-csv/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn          *grpc.ClientConn
	productClient product.ProductServiceClient
}

func NewClient(port int) (*Client, error) {
	var conn *grpc.ClientConn

	addr := fmt.Sprintf(":%d", port)

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:          conn,
		productClient: product.NewProductServiceClient(conn),
	}, nil
}

func (c *Client) CloseConnection() error {
	return c.conn.Close()
}

func (c *Client) Fetch(url string) (*product.FetchResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &product.FetchRequest{
		Url: url,
	}
	fmt.Println("[request ready]")

	resp, err := c.productClient.Fetch(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println("[data received]")
	return resp, nil
}

func (c *Client) List(page int32, pageSize int32, sortField, sortOrder string) (*product.ListResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &product.ListRequest{
		PageNumber: page,
		PageSize:   pageSize,
		SortField:  sortField,
		SortOrder:  sortOrder,
	}

	resp, err := c.productClient.List(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}