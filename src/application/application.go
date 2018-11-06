package application

import (
	"fmt"
	"log"

	"settlement-report/src/calculator"
	"settlement-report/src/loader"
	"settlement-report/src/sorter"
)

// Config contains all the configurations of application
type Config struct {
	DataFilename string
}

// Application contains values required for running application
type Application struct {
	dataFilename           string
	loader                 loader.DataLoader
	calculator             calculator.Calculator
	settlementsSortedSlice sorter.SortedSlice
	entitiesSortedSlice    sorter.SortedSlice
}

// NewApplication creates and returns application
func NewApplication(c *Config) *Application {
	customerLoader := loader.NewInstructionLoader(loader.InstructionLoaderConfig{
		DataFilename: c.DataFilename,
	})

	return &Application{
		dataFilename:           c.DataFilename,
		loader:                 customerLoader,
		calculator:             calculator.NewSettlementsCalculator(),
		settlementsSortedSlice: sorter.NewSettlementsSortedSlice(),
		entitiesSortedSlice:    sorter.NewEntitiesSortedSlice(),
	}
}

// Run starts the processing
func (a *Application) Run() {
	instructions, err := a.loader.Load()
	if err != nil {
		log.Fatal(err)
	}

	settlements := a.calculator.Calculate(instructions)

	for _, settlement := range settlements {
		a.settlementsSortedSlice.Insert(settlement)
	}

	for _, instruction := range instructions {
		a.entitiesSortedSlice.Insert(instruction)
	}

	fmt.Print(a.settlementsSortedSlice.String())
	fmt.Print(a.entitiesSortedSlice.String())
}
