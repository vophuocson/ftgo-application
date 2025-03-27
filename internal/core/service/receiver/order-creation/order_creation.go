package ordercreation

import (
	order "delivery-food/order/internal/core/domain"
	"delivery-food/order/internal/core/port"
	"fmt"
)

type orderCreation struct {
	order *order.Order
	repo  port.OrderRepository
}

func (o *orderCreation) Create() error {
	fmt.Println("this is CreateOrder")
	return nil
}

func (o *orderCreation) Compensate() error {
	fmt.Println("this is CompensateOrder")
	return nil
}

type ticketCreation struct {
	order *order.Order
	p     port.OrderProducer
}

func (o *ticketCreation) Create() error {
	fmt.Println("this is CreateTicket ticket")
	return nil
}

func (o *ticketCreation) Compensate() error {
	fmt.Println("this is CompensateTicket ticket")
	return nil
}

type consumerVerification struct {
	order *order.Order
	p     port.OrderProducer
}

func (o *consumerVerification) Verify() error {
	fmt.Println("this is VerifyConsumer()")
	return nil
}

type orderVerification struct {
	order *order.Order
	c     port.OrderConsumer
}

func (c *orderVerification) Verify() error {
	fmt.Println("this is Verify")
	return nil
}

type cardAuthentication struct {
	order *order.Order
	p     port.OrderProducer
}

func (c *cardAuthentication) Verify() error {
	fmt.Println("this is AuthenticateCard card")
	return nil
}

type ticketApproval struct {
	order *order.Order
	p     port.OrderProducer
}

func (c *ticketApproval) Approve() error {
	fmt.Println("this is ApproveTicket")
	return nil
}

type orderApproval struct {
	order *order.Order
	repo  port.OrderRepository
}

func (c *orderApproval) Approve() error {
	fmt.Println("this is ApproveOrder")
	return nil
}
