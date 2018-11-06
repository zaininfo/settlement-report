package sorter

import (
	"fmt"
	"sort"

	"settlement-report/src/calculator"
	"settlement-report/src/loader"
)

type settlementsSortedSlice struct {
	sortedSlice []*calculator.Settlement
}

// NewSettlementsSortedSlice creates and returns a data structure for storing sorted settlements
func NewSettlementsSortedSlice() SortedSlice {
	return &settlementsSortedSlice{
		sortedSlice: []*calculator.Settlement{},
	}
}

// Insert adds a new settlement to the storage according to its date
func (s *settlementsSortedSlice) Insert(settlement interface{}) {
	settlementStructure := settlement.(*calculator.Settlement)
	length := len(s.sortedSlice)

	fitsAt := sort.Search(length, func(index int) bool {
		return s.sortedSlice[index].Date.After(settlementStructure.Date)
	})

	newSortedSlice := make([]*calculator.Settlement, fitsAt+1)

	copy(newSortedSlice, s.sortedSlice[:fitsAt])
	newSortedSlice[fitsAt] = settlementStructure
	newSortedSlice = append(newSortedSlice, s.sortedSlice[fitsAt:]...)

	s.sortedSlice = newSortedSlice
}

// String returns the data in the storage encoded as string
func (s *settlementsSortedSlice) String() string {
	stringifiedSettlements := "\n"

	for _, settlement := range s.sortedSlice {
		stringifiedSettlements += fmt.Sprintf("Settlement date: %s, Outgoing amount: %g, Incoming amount: %g\n",
			settlement.Date.Format(loader.DateFormat),
			settlement.OutgoingAmount,
			settlement.IncomingAmount,
		)
	}

	return stringifiedSettlements
}
