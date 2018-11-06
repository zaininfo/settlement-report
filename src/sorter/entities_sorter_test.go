package sorter

import (
	"testing"

	"settlement-report/src/loader"

	"github.com/stretchr/testify/assert"
)

func TestEntitiesSortedSlice_InsertAndString(t *testing.T) {
	testCases := []struct {
		instruction         *loader.Instruction
		entitiesSortedSlice string
	}{
		{
			instruction: &loader.Instruction{
				Entity:          "foo",
				InstructionType: loader.Buy,
				AgreedFx:        0.50,
				Units:           200,
				PricePerUnit:    100.25,
			},
			entitiesSortedSlice: "\nRank: 1, Entity name: foo, Outgoing amount: 10025\n\n",
		},
		{
			instruction: &loader.Instruction{
				Entity:          "bar",
				InstructionType: loader.Sell,
				AgreedFx:        0.22,
				Units:           450,
				PricePerUnit:    150.5,
			},
			entitiesSortedSlice: "\nRank: 1, Entity name: foo, Outgoing amount: 10025\n\nRank: 1, Entity name: bar, Incoming amount: 14899.5\n",
		},
		{
			instruction: &loader.Instruction{
				Entity:          "entity1",
				InstructionType: loader.Buy,
				AgreedFx:        0.70,
				Units:           300,
				PricePerUnit:    80,
			},
			entitiesSortedSlice: "\nRank: 1, Entity name: entity1, Outgoing amount: 16800\nRank: 2, Entity name: foo, Outgoing amount: 10025\n\nRank: 1, Entity name: bar, Incoming amount: 14899.5\n",
		},
		{
			instruction: &loader.Instruction{
				Entity:          "entity2",
				InstructionType: loader.Sell,
				AgreedFx:        0.50,
				Units:           200,
				PricePerUnit:    200.25,
			},
			entitiesSortedSlice: "\nRank: 1, Entity name: entity1, Outgoing amount: 16800\nRank: 2, Entity name: foo, Outgoing amount: 10025\n\nRank: 1, Entity name: entity2, Incoming amount: 20025\nRank: 2, Entity name: bar, Incoming amount: 14899.5\n",
		},
	}

	entitiesSortedSlice := NewEntitiesSortedSlice()

	for _, testCase := range testCases {
		entitiesSortedSlice.Insert(testCase.instruction)
		assert.Equal(t, testCase.entitiesSortedSlice, entitiesSortedSlice.String())
	}
}
