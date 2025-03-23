package receiver

import (
	order "delivery-food/order/internal/core/domain"
	"delivery-food/order/internal/core/port"
	"fmt"
)

type OrderCreation struct {
	order *order.Order
	repo  port.OrderRepository
	p     port.OrderProducer
	c     port.OrderConsumer
}

func NewOrderCreation(order *order.Order, repo port.OrderRepository, p port.OrderProducer, c port.OrderConsumer) OrderCreation {
	return OrderCreation{
		order: order,
		repo:  repo,
		p:     p,
		c:     c,
	}
}

func (o *OrderCreation) CreateOrder() error {
	fmt.Println("this is CreateOrder")
	return nil
}

func (o *OrderCreation) CompensateOrder() {
	fmt.Println("this is CompensateOrder")
}

func (o *OrderCreation) VerifyConsumer() error {
	fmt.Println("this is VerifyConsumer()")
	return nil
}

func (o *OrderCreation) CreateTicket() error {
	fmt.Println("this is CreateTicket ticket")
	return nil
}

func (o *OrderCreation) CompensateTicket() error {
	fmt.Println("this is CompensateTicket ticket")
	return nil
}

func (c *OrderCreation) AuthenticateCard() error {
	fmt.Println("this is AuthenticateCard card")
	return nil
}

func (c *OrderCreation) Verify() error {
	fmt.Println("this is Verify")
	return nil
}

func (c *OrderCreation) ApproveTicket() error {
	fmt.Println("this is ApproveTicket")
	return nil
}

func (c *OrderCreation) ApproveOrder() error {
	fmt.Println("this is ApproveOrder")
	return nil
}
