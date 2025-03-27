package ordercreation

import (
	order "delivery-food/order/internal/core/domain"
	"delivery-food/order/internal/core/port"
	"delivery-food/order/internal/core/service/receiver/common"
)

// using abstract factory pattern

type OrderCreationFactory interface {
	CreateOrderCreation() common.AbsTractCreation
	CreateTicketCreation() common.AbsTractCreation
	CreateConsumerVerification() common.AbsTractVerification
	CreateOrderVerification() common.AbsTractVerification
	CreateCardAuthentication() common.AbsTractVerification
	CreateTicketApproval() common.AbstractApproval
	CreateOrderApproval() common.AbstractApproval
}

type StandardOrderProcessFactory struct {
	order *order.Order
	repo  port.OrderRepository
	p     port.OrderProducer
	c     port.OrderConsumer
}

func CreateStandardOrderProcessFactory(order *order.Order, repo port.OrderRepository, p port.OrderProducer, c port.OrderConsumer) OrderCreationFactory {
	return &StandardOrderProcessFactory{order: order, repo: repo, p: p, c: c}
}

func (f *StandardOrderProcessFactory) CreateOrderCreation() common.AbsTractCreation {
	return &orderCreation{order: f.order, repo: f.repo}
}

func (f *StandardOrderProcessFactory) CreateTicketCreation() common.AbsTractCreation {
	return &ticketCreation{order: f.order, p: f.p}
}

func (f *StandardOrderProcessFactory) CreateConsumerVerification() common.AbsTractVerification {
	return &consumerVerification{order: f.order, p: f.p}
}

func (f *StandardOrderProcessFactory) CreateOrderVerification() common.AbsTractVerification {
	return &orderVerification{order: f.order, c: f.c}
}

func (f *StandardOrderProcessFactory) CreateCardAuthentication() common.AbsTractVerification {
	return &cardAuthentication{order: f.order, p: f.p}
}

func (f *StandardOrderProcessFactory) CreateTicketApproval() common.AbstractApproval {
	return &ticketApproval{order: f.order, p: f.p}
}

func (f *StandardOrderProcessFactory) CreateOrderApproval() common.AbstractApproval {
	return &orderApproval{order: f.order, repo: f.repo}
}
