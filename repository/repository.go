package repository

import "time"

type Repository interface {
	Get(id string) error
	Add(model interface{}, duration time.Duration) error
}

type Client struct {
	Context Repository
}

func (circle *Client) Get(id string) error {
	return circle.Context.Get(id)
}

func (circle *Client) Add(model interface{}, duration time.Duration) error {
	return circle.Context.Add(model, duration)
}
