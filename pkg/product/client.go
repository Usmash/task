package product

import "task/pkg/client"

type Params struct {
	Name   string
	Amount int
	Step   int
}

func New(p Params) client.Product {
	return &product{
		name:   p.Name,
		amount: p.Amount,
		step:   p.Step,
	}
}

type product struct {
	name   string
	amount int
	step   int
}

func (p *product) Buy(period int) (float64, int, float64, string) {
	p.step *= period / 3

	amount := float64(p.amount) * (float64(p.step/100) + 1)

	return amount, period, amount / float64(period), p.name
}
