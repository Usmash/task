package tv

import "task/pkg/client"

type Params struct {
	Name   string
	Amount int
}

func New(p Params) client.Product {
	return &tv{
		name:   p.Name,
		amount: p.Amount,
		step:   5,
	}
}

type tv struct {
	name   string
	amount int
	step   int
}

func (tv *tv) Buy(period int) (float64, int, float64, string) {
	tv.step *= period / 3

	amount := float64(tv.amount) * (float64(tv.step/100) + 1)

	return amount, period, amount / float64(period), tv.name
}
