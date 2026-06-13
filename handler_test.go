package main

import (
	"net"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	// Test without arguments
	res := ping([]Value{})
	expected := Value{typ: "string", str: "PONG"}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %+v, got %+v", expected, res)
	}

	// Test with arguments
	res = ping([]Value{{typ: "bulk", bulk: "hello"}})
	expected = Value{typ: "string", str: "hello"}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %+v, got %+v", expected, res)
	}
}

func TestSetAndGet(t *testing.T) {
	// Clear SETs map
	SETs = make(map[string]string)

	// Test GET non-existent
	res := get([]Value{{typ: "bulk", bulk: "foo"}})
	expectedNull := Value{typ: "null"}
	if !reflect.DeepEqual(res, expectedNull) {
		t.Errorf("expected %+v, got %+v", expectedNull, res)
	}

	// Test SET
	res = set([]Value{
		{typ: "bulk", bulk: "foo"},
		{typ: "bulk", bulk: "bar"},
	})
	expectedOK := Value{typ: "string", str: "OK"}
	if !reflect.DeepEqual(res, expectedOK) {
		t.Errorf("expected %+v, got %+v", expectedOK, res)
	}

	// Test GET existent
	res = get([]Value{{typ: "bulk", bulk: "foo"}})
	expectedBulk := Value{typ: "bulk", bulk: "bar"}
	if !reflect.DeepEqual(res, expectedBulk) {
		t.Errorf("expected %+v, got %+v", expectedBulk, res)
	}

	// Test SET invalid args
	res = set([]Value{{typ: "bulk", bulk: "foo"}})
	if res.typ != "error" {
		t.Errorf("expected error, got %+v", res)
	}

	// Test GET invalid args
	res = get([]Value{})
	if res.typ != "error" {
		t.Errorf("expected error, got %+v", res)
	}
}

func TestHSetAndHGet(t *testing.T) {
	// Clear HSETs map
	HSETs = make(map[string]map[string]string)

	// Test HGET non-existent
	res := hget([]Value{
		{typ: "bulk", bulk: "myhash"},
		{typ: "bulk", bulk: "field1"},
	})
	expectedNull := Value{typ: "null"}
	if !reflect.DeepEqual(res, expectedNull) {
		t.Errorf("expected %+v, got %+v", expectedNull, res)
	}

	// Test HSET
	res = hset([]Value{
		{typ: "bulk", bulk: "myhash"},
		{typ: "bulk", bulk: "field1"},
		{typ: "bulk", bulk: "val1"},
	})
	expectedOK := Value{typ: "string", str: "OK"}
	if !reflect.DeepEqual(res, expectedOK) {
		t.Errorf("expected %+v, got %+v", expectedOK, res)
	}

	// Test HGET existent
	res = hget([]Value{
		{typ: "bulk", bulk: "myhash"},
		{typ: "bulk", bulk: "field1"},
	})
	expectedBulk := Value{typ: "bulk", bulk: "val1"}
	if !reflect.DeepEqual(res, expectedBulk) {
		t.Errorf("expected %+v, got %+v", expectedBulk, res)
	}

	// Test HGETALL
	res = hgetall([]Value{
		{typ: "bulk", bulk: "myhash"},
	})
	if res.typ != "array" {
		t.Fatalf("expected array, got %+v", res)
	}
	// Verify it contains field1 and val1
	if len(res.array) != 2 {
		t.Fatalf("expected array length 2, got %d", len(res.array))
	}
	if res.array[0].bulk != "field1" || res.array[1].bulk != "val1" {
		t.Errorf("unexpected HGETALL response elements: %+v", res.array)
	}

	// Test HGETALL non-existent
	res = hgetall([]Value{
		{typ: "bulk", bulk: "nonexistent"},
	})
	if res.typ != "array" || len(res.array) != 0 {
		t.Errorf("expected empty array, got %+v", res)
	}
}

func TestIntegration(t *testing.T) {
	os.Remove("database.aof")
	defer os.Remove("database.aof")

	// Start server in goroutine
	go main()

	// Wait a moment for server to start
	time.Sleep(100 * time.Millisecond)

	// Connect to server
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		t.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Send PING
	_, err = conn.Write([]byte("*1\r\n$4\r\nPING\r\n"))
	if err != nil {
		t.Fatalf("failed to write to server: %v", err)
	}

	// Read response
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		t.Fatalf("failed to read from server: %v", err)
	}

	expected := "+PONG\r\n"
	if string(buf[:n]) != expected {
		t.Errorf("expected %q, got %q", expected, string(buf[:n]))
	}

	// Test SET key value
	_, err = conn.Write([]byte("*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n"))
	if err != nil {
		t.Fatalf("failed to write to server: %v", err)
	}
	n, err = conn.Read(buf)
	if err != nil {
		t.Fatalf("failed to read from server: %v", err)
	}
	if string(buf[:n]) != "+OK\r\n" {
		t.Errorf("expected \"+OK\\r\\n\", got %q", string(buf[:n]))
	}

	// Test GET key
	_, err = conn.Write([]byte("*2\r\n$3\r\nGET\r\n$3\r\nkey\r\n"))
	if err != nil {
		t.Fatalf("failed to write to server: %v", err)
	}
	n, err = conn.Read(buf)
	if err != nil {
		t.Fatalf("failed to read from server: %v", err)
	}
	if string(buf[:n]) != "$5\r\nvalue\r\n" {
		t.Errorf("expected \"$5\\r\\nvalue\\r\\n\", got %q", string(buf[:n]))
	}
}
