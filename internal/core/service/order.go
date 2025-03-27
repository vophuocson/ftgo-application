package service

import (
	order "delivery-food/order/internal/core/domain"
	"delivery-food/order/internal/core/port"
	"delivery-food/order/internal/core/service/command"
	"delivery-food/order/internal/core/service/orchestrator"
	ordercreation "delivery-food/order/internal/core/service/receiver/order-creation"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type orderService struct {
	repo         port.OrderRepository
	p            port.OrderProducer
	c            port.OrderConsumer
	orchestrator port.OrderOrchestratorTransaction
}

func NewOrderService(repo port.OrderRepository, p port.OrderProducer, c port.OrderConsumer, orchestrator port.OrderOrchestratorTransaction) port.OrderService {
	return &orderService{repo: repo, p: p, c: c, orchestrator: orchestrator}
}

func (oS *orderService) CreateOrder(order *order.Order) error {
	creationFactory := ordercreation.CreateStandardOrderProcessFactory(order, oS.repo, oS.p, oS.c)
	saga := orchestrator.Orchestrator{}
	err := saga.Execute(&command.OrderCreationCommand{Receiver: creationFactory.CreateOrderCreation()})
	if err != nil {
		return err
	}
	err = saga.Execute(&command.ConsumerVerificationCommand{Receiver: creationFactory.CreateConsumerVerification()})
	if err != nil {
		return err
	}
	err = saga.Execute(&command.TicketCreationCommand{Receiver: creationFactory.CreateTicketCreation()})
	if err != nil {
		return err
	}
	err = saga.Execute(&command.CardAuthenticationCommand{Receiver: creationFactory.CreateCardAuthentication()})
	if err != nil {
		return err
	}
	err = saga.Execute(&command.OrderVerificationCommand{Receiver: creationFactory.CreateOrderVerification()})
	if err != nil {
		return err
	}
	err = saga.Execute(&command.TicketApprovalCommand{Receiver: creationFactory.CreateTicketApproval()})
	if err != nil {
		return err
	}
	err = saga.Execute(&command.OrderApprovalCommand{Receiver: creationFactory.CreateOrderApproval()})
	if err != nil {
		return err
	}
	return nil
}

func (oS *orderService) FindOrderByID(id uuid.UUID) (*order.Order, error) {
	order, err := oS.repo.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "FindOrderByID(id uuid.UUID)")
	}
	return order, nil
}
