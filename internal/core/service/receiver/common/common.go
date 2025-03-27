package common

type AbsTractCreation interface {
	Create() error
	Compensate() error
}

type AbsTractVerification interface {
	Verify() error
}

type AbstractApproval interface {
	Approve() error
}
