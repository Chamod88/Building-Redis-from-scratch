package main

import (
	"bytes"
	"testing"
)

func TestMarshalString(t *testing.T) {
	v := Value{typ: "string", str: "OK"}
	res := v.Marshal()
	expected := "+OK\r\n"
	if string(res) != expected {
		t.Errorf("expected %q, got %q", expected, string(res))
	}
}

func TestMarshalBulk(t *testing.T) {
	v := Value{typ: "bulk", bulk: "hello"}
	res := v.Marshal()
	expected := "$5\r\nhello\r\n"
	if string(res) != expected {
		t.Errorf("expected %q, got %q", expected, string(res))
	}
}

func TestMarshalArray(t *testing.T) {
	v := Value{
		typ: "array",
		array: []Value{
			{typ: "bulk", bulk: "hello"},
			{typ: "bulk", bulk: "world"},
		},
	}
	res := v.Marshal()
	expected := "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n"
	if string(res) != expected {
		t.Errorf("expected %q, got %q", expected, string(res))
	}
}

func TestMarshalError(t *testing.T) {
	v := Value{typ: "error", str: "Error message"}
	res := v.Marshal()
	expected := "-Error message\r\n"
	if string(res) != expected {
		t.Errorf("expected %q, got %q", expected, string(res))
	}
}

func TestMarshalNull(t *testing.T) {
	v := Value{typ: "null"}
	res := v.Marshal()
	expected := "$-1\r\n"
	if string(res) != expected {
		t.Errorf("expected %q, got %q", expected, string(res))
	}
}

func TestWriterWrite(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)
	err := writer.Write(Value{typ: "string", str: "OK"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "+OK\r\n"
	if buf.String() != expected {
		t.Errorf("expected %q, got %q", expected, buf.String())
	}
}
