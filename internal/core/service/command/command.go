package command

type Command interface {
	Execute() error
	Compensate() error
}
