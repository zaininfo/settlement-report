package loader

import (
	"errors"
)

type instructionType int

const (
	Unknown instructionType = iota
	Buy
	Sell
)

func (i instructionType) String() string {
	names := [...]string{
		"",
		"Outgoing",
		"Incoming",
	}

	return names[i]
}

func parseInstructionType(instructionType string) (instructionType, error) {
	switch instructionType {
	case "B":
		return Buy, nil
	case "S":
		return Sell, nil
	default:
		return Unknown, errors.New("unknown instruction type")
	}
}
