package client

import "fmt"

type Client interface {
	SendSMS(float64, int, float64, string)
	Buy(Product, int) (float64, int, float64, string)
}

type client struct {
	phone string
}

type Params struct {
	Phone string
}

func New(p Params) Client {
	return &client{phone: p.Phone}
}

type Product interface {
	Buy(int) (float64, int, float64, string)
}

func (c *client) SendSMS(sum float64, period int, monthly float64, product string) {
	fmt.Printf(
		"Dear %s. Product %s has been bought for %.2f for period of %d months. Monthly payment is %.2f\n",
		c.phone,
		product,
		sum/100,
		period,
		monthly/100,
	)
}

func (c *client) Buy(p Product, period int) (float64, int, float64, string) {
	return p.Buy(period)
}
