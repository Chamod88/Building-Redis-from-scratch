package main

import (
	"os"
	"testing"
)

func TestAofWriteAndRead(t *testing.T) {
	tempFile := "test_database.aof"
	defer os.Remove(tempFile)

	// Clean up any existing file
	os.Remove(tempFile)

	aof, err := NewAof(tempFile)
	if err != nil {
		t.Fatalf("failed to create NewAof: %v", err)
	}
	defer aof.Close()

	// Write some commands
	val1 := Value{
		typ: "array",
		array: []Value{
			{typ: "bulk", bulk: "SET"},
			{typ: "bulk", bulk: "key1"},
			{typ: "bulk", bulk: "val1"},
		},
	}

	val2 := Value{
		typ: "array",
		array: []Value{
			{typ: "bulk", bulk: "HSET"},
			{typ: "bulk", bulk: "hash1"},
			{typ: "bulk", bulk: "field1"},
			{typ: "bulk", bulk: "val2"},
		},
	}

	err = aof.Write(val1)
	if err != nil {
		t.Fatalf("failed to write val1: %v", err)
	}

	err = aof.Write(val2)
	if err != nil {
		t.Fatalf("failed to write val2: %v", err)
	}

	// Close to release resources before simulating reload
	aof.Close()

	// Open again to simulate server restart and reloading state
	aof2, err := NewAof(tempFile)
	if err != nil {
		t.Fatalf("failed to reopen NewAof: %v", err)
	}
	defer aof2.Close()

	var readVals []Value
	err = aof2.Read(func(value Value) {
		readVals = append(readVals, value)
	})

	if err != nil {
		t.Fatalf("failed to read from aof: %v", err)
	}

	if len(readVals) != 2 {
		t.Fatalf("expected 2 read values, got %d", len(readVals))
	}

	// Verify val1 contents
	if readVals[0].typ != "array" || len(readVals[0].array) != 3 {
		t.Errorf("val1 structure mismatch: %+v", readVals[0])
	} else {
		if readVals[0].array[0].bulk != "SET" || readVals[0].array[1].bulk != "key1" || readVals[0].array[2].bulk != "val1" {
			t.Errorf("val1 value mismatch: %+v", readVals[0])
		}
	}

	// Verify val2 contents
	if readVals[1].typ != "array" || len(readVals[1].array) != 4 {
		t.Errorf("val2 structure mismatch: %+v", readVals[1])
	} else {
		if readVals[1].array[0].bulk != "HSET" || readVals[1].array[1].bulk != "hash1" || readVals[1].array[2].bulk != "field1" || readVals[1].array[3].bulk != "val2" {
			t.Errorf("val2 value mismatch: %+v", readVals[1])
		}
	}
}
