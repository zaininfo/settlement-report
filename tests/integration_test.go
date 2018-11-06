package tests

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"

	"settlement-report/src/application"

	"github.com/stretchr/testify/assert"
)

func TestEndToEnd_NoInstruction(t *testing.T) {
	app := application.NewApplication(&application.Config{
		DataFilename: "data1_test.txt",
	})

	output := captureOutput(app.Run)
	expectedOutput := "\n\n\n"

	assert.Equal(t, expectedOutput, output)
}

func TestEndToEnd_OneInstruction(t *testing.T) {
	app := application.NewApplication(&application.Config{
		DataFilename: "data2_test.txt",
	})

	output := captureOutput(app.Run)
	expectedOutput := "\nSettlement date: 04 Jan 2016, Outgoing amount: 10025, Incoming amount: 0\n\nRank: 1, Entity name: foo, Outgoing amount: 10025\n\n"

	assert.Equal(t, expectedOutput, output)
}

func TestEndToEnd_ManyInstructions(t *testing.T) {
	app := application.NewApplication(&application.Config{
		DataFilename: "data3_test.txt",
	})

	output := captureOutput(app.Run)
	expectedOutput := "\nSettlement date: 01 Mar 2018, Outgoing amount: 0, Incoming amount: 37248.75\nSettlement date: 02 Mar 2018, Outgoing amount: 56980, Incoming amount: 0\nSettlement date: 04 Mar 2018, Outgoing amount: 33679.4, Incoming amount: 30391.199999999997\nSettlement date: 05 Mar 2018, Outgoing amount: 27936, Incoming amount: 10320\n\nRank: 1, Entity name: entity7, Outgoing amount: 56980\nRank: 2, Entity name: entity2, Outgoing amount: 27936\nRank: 3, Entity name: entity6, Outgoing amount: 23654.4\nRank: 4, Entity name: entity1, Outgoing amount: 10025\n\nRank: 1, Entity name: entity3, Incoming amount: 37248.75\nRank: 2, Entity name: entity5, Incoming amount: 30391.199999999997\nRank: 3, Entity name: entity4, Incoming amount: 10320\n"

	assert.Equal(t, expectedOutput, output)
}

// Adapted from: https://stackoverflow.com/a/10476304
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout = w

	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	f()

	w.Close()
	os.Stdout = old
	out := <-outC

	return out
}
