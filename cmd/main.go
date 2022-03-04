package main

import (
	"task/pkg/client"
	"task/pkg/pc"
	"task/pkg/phone"
	"task/pkg/product"
	"task/pkg/tv"
)

func main() {
	p := phone.New(phone.Params{
		Name:   "IPhone 6",
		Amount: 1000_00,
	})

	comp := pc.New(pc.Params{
		Name:   "Asus ROG",
		Amount: 20000_00,
	})

	t := tv.New(tv.Params{
		Name:   "Samsung",
		Amount: 2000_00,
	})

	c := client.New(client.Params{Phone: "+992123456789"})

	c.SendSMS(c.Buy(p, 3))
	c.SendSMS(c.Buy(comp, 12))
	c.SendSMS(c.Buy(t, 18))

	pr := product.New(product.Params{
		Name:   "USB Flash Drive 4GB",
		Amount: 120_00,
		Step:   3,
	})

	c.SendSMS(c.Buy(pr, 6))
}
