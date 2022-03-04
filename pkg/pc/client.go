package pc

import "task/pkg/client"

type Params struct {
	Name   string
	Amount int
}

func New(p Params) client.Product {
	return &pc{
		name:   p.Name,
		amount: p.Amount,
		step:   4,
	}
}

type pc struct {
	name   string
	amount int
	step   int
}

func (pc *pc) Buy(period int) (float64, int, float64, string) {
	pc.step *= period / 3

	amount := float64(pc.amount) * (float64(pc.step/100) + 1)

	return amount, period, amount / float64(period), pc.name
}
