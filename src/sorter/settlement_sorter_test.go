package sorter

import (
	"testing"
	"time"

	"settlement-report/src/calculator"

	"github.com/stretchr/testify/assert"
)

func TestSettlementsSortedSlice_InsertAndString(t *testing.T) {
	testCases := []struct {
		settlement             *calculator.Settlement
		settlementsSortedSlice string
	}{
		{
			settlement: &calculator.Settlement{
				Date:           time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC),
				OutgoingAmount: 12000,
				IncomingAmount: 26000,
			},
			settlementsSortedSlice: "\nSettlement date: 04 Jan 2016, Outgoing amount: 12000, Incoming amount: 26000\n",
		},
		{
			settlement: &calculator.Settlement{
				Date:           time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC),
				OutgoingAmount: 14000,
				IncomingAmount: 30000,
			},
			settlementsSortedSlice: "\nSettlement date: 02 Jan 2016, Outgoing amount: 14000, Incoming amount: 30000\nSettlement date: 04 Jan 2016, Outgoing amount: 12000, Incoming amount: 26000\n",
		},
		{
			settlement: &calculator.Settlement{
				Date:           time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
				OutgoingAmount: 31000,
				IncomingAmount: 22000,
			},
			settlementsSortedSlice: "\nSettlement date: 01 Jan 2016, Outgoing amount: 31000, Incoming amount: 22000\nSettlement date: 02 Jan 2016, Outgoing amount: 14000, Incoming amount: 30000\nSettlement date: 04 Jan 2016, Outgoing amount: 12000, Incoming amount: 26000\n",
		},
		{
			settlement: &calculator.Settlement{
				Date:           time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC),
				OutgoingAmount: 20000,
				IncomingAmount: 17000,
			},
			settlementsSortedSlice: "\nSettlement date: 01 Jan 2016, Outgoing amount: 31000, Incoming amount: 22000\nSettlement date: 02 Jan 2016, Outgoing amount: 14000, Incoming amount: 30000\nSettlement date: 03 Jan 2016, Outgoing amount: 20000, Incoming amount: 17000\nSettlement date: 04 Jan 2016, Outgoing amount: 12000, Incoming amount: 26000\n",
		},
	}

	settlementsSortedSlice := NewSettlementsSortedSlice()

	for _, testCase := range testCases {
		settlementsSortedSlice.Insert(testCase.settlement)
		assert.Equal(t, testCase.settlementsSortedSlice, settlementsSortedSlice.String())
	}
}
