package calculator

import (
	"testing"
	"time"

	"settlement-report/src/loader"

	"github.com/stretchr/testify/assert"
)

func TestSettlementsCalculator_Calculate(t *testing.T) {
	testCases := []struct {
		instructions []interface{}
		settlements  map[string]interface{}
	}{
		{ // validate USD amount calculation
			instructions: []interface{}{
				&loader.Instruction{
					InstructionType: loader.Buy,
					AgreedFx:        0.50,
					Currency:        "SGP",
					SettlementDate:  time.Date(2018, time.March, 1, 0, 0, 0, 0, time.UTC),
					Units:           200,
					PricePerUnit:    100.25,
				},
				&loader.Instruction{
					InstructionType: loader.Sell,
					AgreedFx:        0.22,
					Currency:        "AED",
					SettlementDate:  time.Date(2018, time.March, 1, 0, 0, 0, 0, time.UTC),
					Units:           450,
					PricePerUnit:    150.5,
				},
				&loader.Instruction{
					InstructionType: loader.Sell,
					AgreedFx:        0.70,
					Currency:        "SAR",
					SettlementDate:  time.Date(2018, time.March, 1, 0, 0, 0, 0, time.UTC),
					Units:           300,
					PricePerUnit:    80,
				},
				&loader.Instruction{
					InstructionType: loader.Buy,
					AgreedFx:        1,
					Currency:        "USD",
					SettlementDate:  time.Date(2018, time.March, 1, 0, 0, 0, 0, time.UTC),
					Units:           200,
					PricePerUnit:    200.25,
				},
			},
			settlements: map[string]interface{}{
				"2018-03-01 00:00:00 +0000 UTC": &Settlement{
					Date:           time.Date(2018, time.March, 1, 0, 0, 0, 0, time.UTC),
					OutgoingAmount: 50075,
					IncomingAmount: 31699.5,
				},
			},
		},
		{ // validate settlements on working days
			instructions: []interface{}{
				&loader.Instruction{
					InstructionType: loader.Buy,
					AgreedFx:        0.50,
					Currency:        "SAR",
					SettlementDate:  time.Date(2018, time.March, 3, 0, 0, 0, 0, time.UTC),
					Units:           200,
					PricePerUnit:    100.25,
				},
				&loader.Instruction{
					InstructionType: loader.Buy,
					AgreedFx:        0.60,
					Currency:        "SGP",
					SettlementDate:  time.Date(2018, time.March, 3, 0, 0, 0, 0, time.UTC),
					Units:           300,
					PricePerUnit:    155.2,
				},
				&loader.Instruction{
					InstructionType: loader.Sell,
					AgreedFx:        0.55,
					Currency:        "SAR",
					SettlementDate:  time.Date(2018, time.March, 1, 0, 0, 0, 0, time.UTC),
					Units:           450,
					PricePerUnit:    150.5,
				},
				&loader.Instruction{
					InstructionType: loader.Sell,
					AgreedFx:        0.43,
					Currency:        "PKR",
					SettlementDate:  time.Date(2018, time.March, 4, 0, 0, 0, 0, time.UTC),
					Units:           300,
					PricePerUnit:    80,
				},
				&loader.Instruction{
					InstructionType: loader.Sell,
					AgreedFx:        0.70,
					Currency:        "AED",
					SettlementDate:  time.Date(2018, time.March, 4, 0, 0, 0, 0, time.UTC),
					Units:           270,
					PricePerUnit:    160.8,
				},
				&loader.Instruction{
					InstructionType: loader.Buy,
					AgreedFx:        0.66,
					Currency:        "AED",
					SettlementDate:  time.Date(2018, time.March, 2, 0, 0, 0, 0, time.UTC),
					Units:           320,
					PricePerUnit:    112,
				},
				&loader.Instruction{
					InstructionType: loader.Buy,
					AgreedFx:        0.80,
					Currency:        "EUR",
					SettlementDate:  time.Date(2018, time.March, 2, 0, 0, 0, 0, time.UTC),
					Units:           350,
					PricePerUnit:    203.5,
				},
			},
			settlements: map[string]interface{}{
				"2018-03-01 00:00:00 +0000 UTC": &Settlement{
					Date:           time.Date(2018, time.March, 1, 0, 0, 0, 0, time.UTC),
					OutgoingAmount: 0,
					IncomingAmount: 37248.75,
				},
				"2018-03-02 00:00:00 +0000 UTC": &Settlement{
					Date:           time.Date(2018, time.March, 2, 0, 0, 0, 0, time.UTC),
					OutgoingAmount: 56980,
					IncomingAmount: 0,
				},
				"2018-03-04 00:00:00 +0000 UTC": &Settlement{
					Date:           time.Date(2018, time.March, 4, 0, 0, 0, 0, time.UTC),
					OutgoingAmount: 33679.4,
					IncomingAmount: 30391.199999999997,
				},
				"2018-03-05 00:00:00 +0000 UTC": &Settlement{
					Date:           time.Date(2018, time.March, 5, 0, 0, 0, 0, time.UTC),
					OutgoingAmount: 27936,
					IncomingAmount: 10320,
				},
			},
		},
	}

	settlementsCalculator := NewSettlementsCalculator()

	for _, testCase := range testCases {
		settlements := settlementsCalculator.Calculate(testCase.instructions)
		assert.Equal(t, testCase.settlements, settlements)
	}
}
