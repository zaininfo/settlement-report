package loader

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"
)

const DateFormat = "02 Jan 2006"

// DataLoader defines the interface for loading data
type DataLoader interface {
	Load() ([]interface{}, error)
}

// Instruction defines the structure of an instruction
type Instruction struct {
	Entity          string
	InstructionType instructionType
	AgreedFx        float64
	Currency        string
	InstructionDate time.Time
	SettlementDate  time.Time
	Units           float64
	PricePerUnit    float64
}

type instructionLoader struct {
	dataFilename string
}

// InstructionLoaderConfig defines the structure of configurations for instruction loader
type InstructionLoaderConfig struct {
	DataFilename string
}

// NewInstructionLoader creates and returns a loader for instruction data
func NewInstructionLoader(c InstructionLoaderConfig) DataLoader {
	return &instructionLoader{
		dataFilename: c.DataFilename,
	}
}

// Load reads and returns all instructions
func (l *instructionLoader) Load() ([]interface{}, error) {
	file, err := os.Open(l.dataFilename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	var instructions []interface{}
	first := true

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if first {
			first = false
			continue
		}

		var instruction *Instruction
		if instruction, err = parseColumns(line); err != nil {
			return nil, err
		}

		instructions = append(instructions, instruction)
	}

	return instructions, nil
}

func parseColumns(columns []string) (instruction *Instruction, err error) {
	var instructionType instructionType
	var agreedFx, units, pricePerUnit float64
	var instructionDate, settlementDate time.Time

	if instructionType, err = parseInstructionType(columns[1]); err != nil {
		return
	}

	if agreedFx, err = strconv.ParseFloat(columns[2], 64); err != nil {
		return
	}

	if units, err = strconv.ParseFloat(columns[6], 64); err != nil {
		return
	}

	if pricePerUnit, err = strconv.ParseFloat(columns[7], 64); err != nil {
		return
	}

	if instructionDate, err = time.Parse(DateFormat, columns[4]); err != nil {
		return
	}

	if settlementDate, err = time.Parse(DateFormat, columns[5]); err != nil {
		return
	}

	return &Instruction{
		Entity:          columns[0],
		InstructionType: instructionType,
		AgreedFx:        agreedFx,
		Currency:        columns[3],
		InstructionDate: instructionDate,
		SettlementDate:  settlementDate,
		Units:           units,
		PricePerUnit:    pricePerUnit,
	}, nil
}
