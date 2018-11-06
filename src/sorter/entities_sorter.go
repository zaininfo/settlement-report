package sorter

import (
	"fmt"
	"sort"

	"settlement-report/src/calculator"
	"settlement-report/src/loader"
)

// Entity defines the structure of an entity
type Entity struct {
	name   string
	amount float64
}

type entitiesSortedSlice struct {
	outgoingSortedSlice []*Entity
	incomingSortedSlice []*Entity
}

// NewEntitiesSortedSlice creates and returns a data structure for storing sorted entities
func NewEntitiesSortedSlice() SortedSlice {
	return &entitiesSortedSlice{
		outgoingSortedSlice: []*Entity{},
		incomingSortedSlice: []*Entity{},
	}
}

// Insert adds a new entity to the storage according to its instruction's amount
func (s *entitiesSortedSlice) Insert(instruction interface{}) {
	instructionStructure := instruction.(*loader.Instruction)
	entity := &Entity{
		name:   instructionStructure.Entity,
		amount: calculator.AmountInUsd(instructionStructure),
	}

	switch instructionStructure.InstructionType {
	case loader.Buy:
		s.outgoingSortedSlice = insertEntity(entity, s.outgoingSortedSlice)
	case loader.Sell:
		s.incomingSortedSlice = insertEntity(entity, s.incomingSortedSlice)
	}

}

// String returns the data in the storage encoded as string
func (s *entitiesSortedSlice) String() string {
	return stringify(s.outgoingSortedSlice, loader.Buy.String()) + stringify(s.incomingSortedSlice, loader.Sell.String())
}

func insertEntity(entity *Entity, sortedSlice []*Entity) []*Entity {
	length := len(sortedSlice)

	fitsAt := sort.Search(length, func(index int) bool {
		return sortedSlice[index].amount <= entity.amount
	})

	newSortedSlice := make([]*Entity, fitsAt+1)

	copy(newSortedSlice, sortedSlice[:fitsAt])
	newSortedSlice[fitsAt] = entity
	newSortedSlice = append(newSortedSlice, sortedSlice[fitsAt:]...)

	return newSortedSlice
}

func stringify(sortedSlice []*Entity, amountType string) string {
	stringifiedEntities := "\n"

	for index, entity := range sortedSlice {
		stringifiedEntities += fmt.Sprintf("Rank: %d, Entity name: %s, %s amount: %g\n",
			index+1,
			entity.name,
			amountType,
			entity.amount,
		)
	}

	return stringifiedEntities
}
