package loader

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInstructionLoader_Load(t *testing.T) {
	instructions := []*Instruction{
		{
			Entity:          "foo",
			InstructionType: Buy,
			AgreedFx:        0.50,
			Currency:        "SGP",
			InstructionDate: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
			SettlementDate:  time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC),
			Units:           200,
			PricePerUnit:    100.25,
		},
		{
			Entity:          "bar",
			InstructionType: Sell,
			AgreedFx:        0.22,
			Currency:        "AED",
			InstructionDate: time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC),
			SettlementDate:  time.Date(2016, time.January, 7, 0, 0, 0, 0, time.UTC),
			Units:           450,
			PricePerUnit:    150.5,
		},
		{
			Entity:          "entity1",
			InstructionType: Buy,
			AgreedFx:        0.70,
			Currency:        "SAR",
			InstructionDate: time.Date(2016, time.January, 12, 0, 0, 0, 0, time.UTC),
			SettlementDate:  time.Date(2016, time.January, 14, 0, 0, 0, 0, time.UTC),
			Units:           300,
			PricePerUnit:    80,
		},
	}

	instructionLoader := NewInstructionLoader(InstructionLoaderConfig{
		DataFilename: "data_test.txt",
	})

	data, err := instructionLoader.Load()
	if err != nil {
		t.Fatal("Data loading failed: ", err)
	}

	for index, item := range data {
		assert.Equal(t, instructions[index], item.(*Instruction))
	}
}
