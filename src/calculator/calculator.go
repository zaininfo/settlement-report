package calculator

import (
	"time"

	"settlement-report/src/loader"
)

// Calculator defines the interface for calculating a set of values
type Calculator interface {
	Calculate([]interface{}) map[string]interface{}
}

// Settlement defines the structure of a settlement
type Settlement struct {
	Date           time.Time
	OutgoingAmount float64
	IncomingAmount float64
}

type settlementsCalculator struct{}

// NewSettlementsCalculator creates and returns a calculator for settlements
func NewSettlementsCalculator() Calculator {
	return &settlementsCalculator{}
}

// Calculate calculates the daily settlements
func (c *settlementsCalculator) Calculate(instructions []interface{}) map[string]interface{} {
	settlements := make(map[string]interface{})

	for _, rawInstruction := range instructions {
		if rawInstruction == nil {
			continue
		}

		instruction := rawInstruction.(*loader.Instruction)
		workingSettlementDate := workingSettlementDate(instruction.SettlementDate, instruction.Currency)
		workingSettlementDateKey := workingSettlementDate.String()

		if _, ok := settlements[workingSettlementDateKey]; !ok {
			settlements[workingSettlementDateKey] = &Settlement{
				Date: workingSettlementDate,
			}
		}

		settlement := settlements[workingSettlementDateKey].(*Settlement)

		switch instruction.InstructionType {
		case loader.Buy:
			settlement.OutgoingAmount += AmountInUsd(instruction)
		case loader.Sell:
			settlement.IncomingAmount += AmountInUsd(instruction)
		}

		settlements[workingSettlementDateKey] = settlement
	}

	return settlements
}

func workingSettlementDate(settlementDate time.Time, currency string) time.Time {
	settlementWeekday := settlementDate.Weekday()

	switch currency {
	case "AED", "SAR":
		if settlementWeekday == time.Friday {
			return settlementDate.AddDate(0, 0, 2)
		}
		if settlementWeekday == time.Saturday {
			return settlementDate.AddDate(0, 0, 1)
		}
	default:
		if settlementWeekday == time.Saturday {
			return settlementDate.AddDate(0, 0, 2)
		}
		if settlementWeekday == time.Sunday {
			return settlementDate.AddDate(0, 0, 1)
		}
	}

	return settlementDate
}

// AmountInUsd calculates the amount of an instruction in USD
func AmountInUsd(instruction *loader.Instruction) float64 {
	return instruction.PricePerUnit * instruction.Units * instruction.AgreedFx
}
