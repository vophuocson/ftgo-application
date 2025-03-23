package orchestrator

import (
	"delivery-food/order/internal/core/port"
	"delivery-food/order/internal/core/service/command"
	"errors"
)

type CommandNode struct {
	Command command.Command
	Next    *CommandNode
}

type CommandStack struct {
	Current *CommandNode
}

func (s *CommandStack) Push(c command.Command) {
	s.Current = &CommandNode{Next: s.Current, Command: c}
}

func (s *CommandStack) Pop() (command.Command, error) {
	if current := s.Current; current != nil {
		s.Current = s.Current.Next
		return current.Command, nil
	}
	return nil, errors.New("stack is empty")
}

type Orchestrator struct {
	worker port.Worker
	Stack  *CommandStack
}

func (o *Orchestrator) Execute(c command.Command) error {
	err := o.worker.Execute(func() error {
		err := c.Execute()
		if err != nil {
			o.Undo()
			return err
		}
		o.Stack.Push(c)
		return nil
	})
	return err
}

func (o *Orchestrator) Undo() {
	cmd, err := o.Stack.Pop()
	if err != nil {
		return
	}
	cmd.Compensate()
	o.Undo()
}
