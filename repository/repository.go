package repository

import "time"

type Repository interface {
	Get(model interface{}) error
	Add(model interface{}, duration time.Duration) error
}

type Client struct {
	Context Repository
}

func (circle *Client) Get(model interface{}) error {
	return circle.Context.Get(model)
}

func (circle *Client) Add(model interface{}, duration time.Duration) error {
	return circle.Context.Add(model, duration)
}
