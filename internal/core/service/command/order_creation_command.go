package command

import (
	"delivery-food/order/internal/core/service/receiver/common"
)

type OrderCreationCommand struct {
	Receiver common.AbsTractCreation
}

func (c *OrderCreationCommand) Execute() error {
	return c.Receiver.Create()
}

func (c *OrderCreationCommand) Compensate() error {
	return nil
}

type ConsumerVerificationCommand struct {
	Receiver common.AbsTractVerification
}

func (c *ConsumerVerificationCommand) Execute() error {
	return c.Receiver.Verify()
}

func (c *ConsumerVerificationCommand) Compensate() error {
	return nil
}

type TicketCreationCommand struct {
	Receiver common.AbsTractCreation
}

func (c *TicketCreationCommand) Execute() error {
	return c.Receiver.Create()
}

func (c *TicketCreationCommand) Compensate() error {
	return c.Receiver.Compensate()
}

type CardAuthenticationCommand struct {
	Receiver common.AbsTractVerification
}

func (c *CardAuthenticationCommand) Execute() error {
	return c.Receiver.Verify()
}

func (c *CardAuthenticationCommand) Compensate() error {
	return nil
}

type OrderVerificationCommand struct {
	Receiver common.AbsTractVerification
}

func (c *OrderVerificationCommand) Execute() error {
	return c.Receiver.Verify()
}

func (c *OrderVerificationCommand) Compensate() error {
	return nil
}

type TicketApprovalCommand struct {
	Receiver common.AbstractApproval
}

func (c *TicketApprovalCommand) Execute() error {
	return c.Receiver.Approve()
}

func (c *TicketApprovalCommand) Compensate() error {
	return nil
}

type OrderApprovalCommand struct {
	Receiver common.AbstractApproval
}

func (c *OrderApprovalCommand) Execute() error {
	return c.Receiver.Approve()
}

func (c *OrderApprovalCommand) Compensate() error {
	return nil
}
