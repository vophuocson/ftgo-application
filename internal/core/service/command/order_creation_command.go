package command

import "delivery-food/order/internal/core/service/receiver"

type OrderCreationCommand struct {
	Receiver *receiver.OrderCreation
}

func (c *OrderCreationCommand) Execute() error {
	return c.Receiver.CreateOrder()
}

func (c *OrderCreationCommand) Compensate() error {
	return nil
}

type ConsumerVerificationCommand struct {
	Receiver *receiver.OrderCreation
}

func (c *ConsumerVerificationCommand) Execute() error {
	return c.Receiver.VerifyConsumer()
}

func (c *ConsumerVerificationCommand) Compensate() error {
	return nil
}

type TicketCreationCommand struct {
	Receiver *receiver.OrderCreation
}

func (c *TicketCreationCommand) Execute() error {
	return c.Receiver.CreateTicket()
}

func (c *TicketCreationCommand) Compensate() error {
	return c.Receiver.CompensateTicket()
}

type CardAuthenticationCommand struct {
	Receiver *receiver.OrderCreation
}

func (c *CardAuthenticationCommand) Execute() error {
	return c.Receiver.AuthenticateCard()
}

func (c *CardAuthenticationCommand) Compensate() error {
	return nil
}

type OrderVerificationCommand struct {
	Receiver *receiver.OrderCreation
}

func (c *OrderVerificationCommand) Execute() error {
	return c.Receiver.Verify()
}

func (c *OrderVerificationCommand) Compensate() error {
	return nil
}

type TicketApprovalCommand struct {
	Receiver *receiver.OrderCreation
}

func (c *TicketApprovalCommand) Execute() error {
	return c.Receiver.ApproveTicket()
}

func (c *TicketApprovalCommand) Compensate() error {
	return nil
}

type OrderApprovalCommand struct {
	Receiver *receiver.OrderCreation
}

func (c *OrderApprovalCommand) Execute() error {
	return c.Receiver.ApproveOrder()
}

func (c *OrderApprovalCommand) Compensate() error {
	return nil
}
