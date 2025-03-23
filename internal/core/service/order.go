package service

import (
	order "delivery-food/order/internal/core/domain"
	"delivery-food/order/internal/core/port"
	"delivery-food/order/internal/core/port/dto"
	"delivery-food/order/internal/core/port/workflow"
	"delivery-food/order/internal/core/service/command"
	"delivery-food/order/internal/core/service/orchestrator"
	"delivery-food/order/internal/core/service/receiver"

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
	receiver := receiver.NewOrderCreation(order, oS.repo, oS.p, oS.c)
	saga := orchestrator.Orchestrator{}
	err := saga.Execute(&command.OrderCreationCommand{Receiver: &receiver})
	if err != nil {
		return err
	}
	err = saga.Execute(&command.ConsumerVerificationCommand{Receiver: &receiver})
	if err != nil {
		return err
	}
	err = saga.Execute(&command.TicketCreationCommand{Receiver: &receiver})
	if err != nil {
		return err
	}
	err = saga.Execute(&command.CardAuthenticationCommand{Receiver: &receiver})
	if err != nil {
		return err
	}
	err = saga.Execute(&command.OrderVerificationCommand{Receiver: &receiver})
	if err != nil {
		return err
	}
	err = saga.Execute(&command.TicketApprovalCommand{Receiver: &receiver})
	if err != nil {
		return err
	}
	err = saga.Execute(&command.OrderApprovalCommand{Receiver: &receiver})
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

func (oS *orderService) createOrderWorkflowDefinition(o *order.Order) *workflow.WorkflowDefinition {
	var workflowDefinition workflow.WorkflowDefinition
	var compensateTransaction workflow.Activity
	compensateTransaction.AddStep(&workflow.Step{
		Command: func() error { return oS.p.VerifyConsumer(o) },
	})
	compensateTransaction.AddStep(&workflow.Step{
		Command:        func() error { return oS.p.CreateTicket(o) },
		CompensateFunc: func() error { return oS.p.CompensateTicket(o) },
	})
	compensateTransaction.AddStep(&workflow.Step{
		Command: func() error { return oS.p.AuthenticateCard(o) },
	})
	workflowDefinition.Steps = append(workflowDefinition.Steps, &compensateTransaction)
	var confirmOrderCreationActivity workflow.Activity
	confirmOrderCreationActivity.AddStep(&workflow.Step{
		Command: func() error {
			return oS.c.ConfirmOrderCreation(&dto.ConfirmCreateOrder{
				OrderID: o.ID,
				ChannelNamesReply: map[string]bool{
					"order-service.kitchen.create.dev.v1":       false,
					"order-service.customer.verify.dev.v1":      false,
					"order-service.payment.authenticate.dev.v1": false,
				},
			})
		},
	})
	workflowDefinition.Steps = append(workflowDefinition.Steps, &confirmOrderCreationActivity)
	var approveOrderCreationActivity workflow.Activity
	approveOrderCreationActivity.AddStep(&workflow.Step{
		Command: func() error { return oS.p.ApproveTicketCreation(o) },
	})
	approveOrderCreationActivity.AddStep(&workflow.Step{
		Command: func() error { return oS.p.ApproveOrderCreation(o) },
	})
	workflowDefinition.Steps = append(workflowDefinition.Steps, &approveOrderCreationActivity)
	return &workflowDefinition
}

//  order có data rồi
// define command
// order service đóng vai trò là invoker
// receiver là ai
// command là ai

type OrderCreateSaga struct {
	Order order.Order
}

// defines funcs for create order

type Command interface {
	Execute()
	Compensate()
}
type OrderCreate struct {
	Order order.Order
	repo  port.OrderRepository
	p     port.OrderProducer
	c     port.OrderConsumer
}

func (o *OrderCreate) Execute() {
}

func (o *OrderCreate) Compensate() {

}
