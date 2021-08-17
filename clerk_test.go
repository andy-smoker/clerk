package main

import (
	"testing"
	"time"
)

func TestPrinter(t *testing.T) {
	p := NewPrinter("trace", "test", "")
	t.Log(p.WriteLog(0, time.Now(), "test log2"))
}
